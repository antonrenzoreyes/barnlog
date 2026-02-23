package sqlite

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"barnlog/backend/internal/infrastructure/sqlite/sqlc"
	"barnlog/backend/internal/ports"

	modernsqlite "modernc.org/sqlite"
)

const (
	createAnimalAggregateType = "animal"
	createAnimalEventType     = "animal.created"
	createAnimalCreatedBy     = "system"
)

type animalWriteStore struct {
	queries  *sqlc.Queries
	photoDir string
	now      func() time.Time
}

// NewAnimalWriteStore builds the SQLite implementation of ports.AnimalWriteStore.
func NewAnimalWriteStore(db *sql.DB, photoDir string) ports.AnimalWriteStore {
	return animalWriteStore{
		queries:  sqlc.New(db),
		photoDir: photoDir,
		now:      time.Now,
	}
}

func (s animalWriteStore) CreateAnimalRecord(ctx context.Context, in ports.CreateAnimalRecordInput) (ports.CreateAnimalRecordOutput, error) {
	animalID, err := newID()
	if err != nil {
		return ports.CreateAnimalRecordOutput{}, fmt.Errorf("generate animal id: %w", err)
	}
	eventID, err := newID()
	if err != nil {
		return ports.CreateAnimalRecordOutput{}, fmt.Errorf("generate event id: %w", err)
	}

	payloadJSON, err := createAnimalPayloadJSON(in)
	if err != nil {
		return ports.CreateAnimalRecordOutput{}, err
	}

	metadataJSON, err := json.Marshal(map[string]string{
		"source":     in.Source,
		"request_id": in.RequestID,
	})
	if err != nil {
		return ports.CreateAnimalRecordOutput{}, fmt.Errorf("marshal metadata: %w", err)
	}

	occurredAt := s.now().UTC().Format(time.RFC3339)
	if err := s.queries.CreateEvent(ctx, sqlc.CreateEventParams{
		ID:            eventID,
		AggregateType: createAnimalAggregateType,
		AggregateID:   animalID,
		EventType:     createAnimalEventType,
		CreatedBy:     createAnimalCreatedBy,
		Source:        in.Source,
		RequestID:     in.RequestID,
		EventVersion:  1,
		PayloadJson:   string(payloadJSON),
		MetadataJson: sql.NullString{
			String: string(metadataJSON),
			Valid:  true,
		},
		OccurredAt: occurredAt,
	}); err != nil {
		if isUniqueConstraint(err) {
			out, found, replayErr := s.FindCreateAnimalReplay(ctx, in)
			if replayErr != nil {
				return ports.CreateAnimalRecordOutput{}, replayErr
			}
			if found {
				return out, nil
			}
			return ports.CreateAnimalRecordOutput{}, fmt.Errorf("%w", ports.ErrConflict)
		}
		return ports.CreateAnimalRecordOutput{}, fmt.Errorf("create event: %w", err)
	}

	return ports.CreateAnimalRecordOutput{
		AnimalID: animalID,
		EventID:  eventID,
		Replayed: false,
	}, nil
}

func (s animalWriteStore) FindCreateAnimalReplay(ctx context.Context, in ports.CreateAnimalRecordInput) (ports.CreateAnimalRecordOutput, bool, error) {
	payloadJSON, err := createAnimalPayloadJSON(in)
	if err != nil {
		return ports.CreateAnimalRecordOutput{}, false, err
	}

	existing, err := s.queries.GetEventBySourceRequestID(ctx, sqlc.GetEventBySourceRequestIDParams{
		Source:    in.Source,
		RequestID: in.RequestID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ports.CreateAnimalRecordOutput{}, false, nil
		}
		return ports.CreateAnimalRecordOutput{}, false, fmt.Errorf("load existing event by idempotency key: %w", err)
	}

	if existing.AggregateType != createAnimalAggregateType || existing.EventType != createAnimalEventType {
		return ports.CreateAnimalRecordOutput{}, false, fmt.Errorf(
			"%w: %s/%s",
			ports.ErrIdempotencyEventTypeMismatch,
			existing.AggregateType,
			existing.EventType,
		)
	}

	if existing.PayloadJson != string(payloadJSON) {
		return ports.CreateAnimalRecordOutput{}, false, fmt.Errorf("%w", ports.ErrIdempotencyPayloadMismatch)
	}

	return ports.CreateAnimalRecordOutput{
		AnimalID: existing.AggregateID,
		EventID:  existing.ID,
		Replayed: true,
	}, true, nil
}

func createAnimalPayloadJSON(in ports.CreateAnimalRecordInput) ([]byte, error) {
	payloadJSON, err := json.Marshal(map[string]string{
		"name":      in.Name,
		"species":   in.Species,
		"tag":       in.Tag,
		"birthdate": in.Birthdate,
		"photo_id":  in.PhotoID,
	})
	if err != nil {
		return nil, fmt.Errorf("marshal payload: %w", err)
	}
	return payloadJSON, nil
}

func (s animalWriteStore) PhotoExists(_ context.Context, photoID string) (exists bool, err error) {
	root, err := os.OpenRoot(s.photoDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, fmt.Errorf("open photo root: %w", err)
	}
	defer func() {
		if closeErr := root.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close photo root: %w", closeErr)
		}
	}()

	file, err := root.Open(photoID)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, fmt.Errorf("open photo: %w", err)
	}
	_ = file.Close()
	return true, nil
}

func isUniqueConstraint(err error) bool {
	if sqliteErr, ok := errors.AsType[*modernsqlite.Error](err); ok {
		// SQLite constraint codes: 19=constraint, 2067=unique constraint.
		return sqliteErr.Code() == 19 || sqliteErr.Code() == 2067
	}
	return false
}

func newID() (string, error) {
	var bytes [16]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		return "", fmt.Errorf("generate random bytes: %w", err)
	}
	return hex.EncodeToString(bytes[:]), nil
}
