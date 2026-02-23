package ports

import (
	"context"
	"errors"
)

// ErrConflict signals an idempotency or uniqueness conflict in write storage.
var ErrConflict = errors.New("conflict")

// ErrIdempotencyPayloadMismatch signals same idempotency key but different payload.
var ErrIdempotencyPayloadMismatch = errors.New("idempotency_payload_mismatch")

// ErrIdempotencyEventTypeMismatch signals same idempotency key used for another event type.
var ErrIdempotencyEventTypeMismatch = errors.New("idempotency_event_type_mismatch")

// CreateAnimalRecordInput is the storage-level payload for writing animal-created events.
type CreateAnimalRecordInput struct {
	Name      string
	Species   string
	Tag       string
	Birthdate string
	PhotoID   string
	Source    string
	RequestID string
}

// CreateAnimalRecordOutput contains IDs produced by persisted animal creation.
type CreateAnimalRecordOutput struct {
	AnimalID string
	EventID  string
	Replayed bool
}

// AnimalWriteStore defines persistence operations needed by create-animal use cases.
type AnimalWriteStore interface {
	FindCreateAnimalReplay(ctx context.Context, in CreateAnimalRecordInput) (CreateAnimalRecordOutput, bool, error)
	CreateAnimalRecord(ctx context.Context, in CreateAnimalRecordInput) (CreateAnimalRecordOutput, error)
	PhotoExists(ctx context.Context, photoID string) (bool, error)
}
