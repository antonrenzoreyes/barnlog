package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"barnlog/backend/internal/infrastructure/sqlite/sqlc"
	"barnlog/backend/internal/ports"
)

func TestAnimalWriteStore_CreateAnimalRecord_IdempotentReplay(t *testing.T) {
	store, db := newTestAnimalWriteStore(t)
	t.Cleanup(func() { _ = db.Close() })

	in := ports.CreateAnimalRecordInput{
		Name:      "Nanny",
		Species:   "goat",
		Tag:       "G-7",
		Birthdate: "2021-03-04",
		PhotoID:   "photo_1",
		Source:    "test.api",
		RequestID: "req-1",
	}

	first, err := store.CreateAnimalRecord(context.Background(), in)
	if err != nil {
		t.Fatalf("first create: %v", err)
	}
	if first.Replayed {
		t.Fatalf("expected first write not replayed")
	}

	second, err := store.CreateAnimalRecord(context.Background(), in)
	if err != nil {
		t.Fatalf("second create (replay): %v", err)
	}
	if !second.Replayed {
		t.Fatalf("expected replayed second write")
	}
	if second.AnimalID != first.AnimalID {
		t.Fatalf("expected replay animal_id=%q, got %q", first.AnimalID, second.AnimalID)
	}
	if second.EventID != first.EventID {
		t.Fatalf("expected replay event_id=%q, got %q", first.EventID, second.EventID)
	}
}

func TestAnimalWriteStore_CreateAnimalRecord_PayloadMismatch(t *testing.T) {
	store, db := newTestAnimalWriteStore(t)
	t.Cleanup(func() { _ = db.Close() })

	_, err := store.CreateAnimalRecord(context.Background(), ports.CreateAnimalRecordInput{
		Name:      "Nanny",
		Species:   "goat",
		Tag:       "G-7",
		Birthdate: "2021-03-04",
		PhotoID:   "photo_1",
		Source:    "test.api",
		RequestID: "req-1",
	})
	if err != nil {
		t.Fatalf("seed create: %v", err)
	}

	_, err = store.CreateAnimalRecord(context.Background(), ports.CreateAnimalRecordInput{
		Name:      "Nanny",
		Species:   "goat",
		Tag:       "G-8",
		Birthdate: "2021-03-04",
		PhotoID:   "photo_1",
		Source:    "test.api",
		RequestID: "req-1",
	})
	if !errors.Is(err, ports.ErrIdempotencyPayloadMismatch) {
		t.Fatalf("expected ErrIdempotencyPayloadMismatch, got %v", err)
	}
}

func TestAnimalWriteStore_CreateAnimalRecord_EventTypeMismatch(t *testing.T) {
	store, db := newTestAnimalWriteStore(t)
	t.Cleanup(func() { _ = db.Close() })

	queries := sqlc.New(db)
	err := queries.CreateEvent(context.Background(), sqlc.CreateEventParams{
		ID:            "event_existing",
		AggregateType: "photo",
		AggregateID:   "photo_1",
		EventType:     "photo.uploaded",
		CreatedBy:     "system",
		Source:        "test.api",
		RequestID:     "req-1",
		EventVersion:  1,
		PayloadJson:   `{"photo_id":"photo_1"}`,
		MetadataJson: sql.NullString{
			String: `{"source":"test.api","request_id":"req-1"}`,
			Valid:  true,
		},
		OccurredAt: time.Now().UTC().Format(time.RFC3339),
	})
	if err != nil {
		t.Fatalf("insert existing event: %v", err)
	}

	_, err = store.CreateAnimalRecord(context.Background(), ports.CreateAnimalRecordInput{
		Name:      "Nanny",
		Species:   "goat",
		Tag:       "G-7",
		Birthdate: "2021-03-04",
		PhotoID:   "photo_1",
		Source:    "test.api",
		RequestID: "req-1",
	})
	if !errors.Is(err, ports.ErrIdempotencyEventTypeMismatch) {
		t.Fatalf("expected ErrIdempotencyEventTypeMismatch, got %v", err)
	}
}
