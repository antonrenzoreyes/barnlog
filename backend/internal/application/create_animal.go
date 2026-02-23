package application

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"barnlog/backend/internal/ports"
)

// BusinessCode identifies stable business-level error categories.
type BusinessCode string

const (
	// CodeInvalidInput indicates missing or invalid command metadata.
	CodeInvalidInput BusinessCode = "invalid_input"
	// CodePhotoNotFound indicates referenced photo content does not exist.
	CodePhotoNotFound BusinessCode = "photo_not_found"
	// CodeConflict indicates an idempotency conflict for the same request.
	CodeConflict BusinessCode = "conflict"
	// CodeIdempotencyPayloadMismatch indicates same request key with different payload.
	CodeIdempotencyPayloadMismatch BusinessCode = "idempotency_payload_mismatch"
	// CodeIdempotencyEventTypeMismatch indicates same request key used for another event type.
	CodeIdempotencyEventTypeMismatch BusinessCode = "idempotency_event_type_mismatch"
)

// BusinessError wraps a business code and optional underlying cause.
type BusinessError struct {
	Code BusinessCode
	Err  error
}

// Error returns the underlying error message or business code.
func (e BusinessError) Error() string {
	if e.Err == nil {
		return string(e.Code)
	}
	return e.Err.Error()
}

// Unwrap returns the underlying cause.
func (e BusinessError) Unwrap() error {
	return e.Err
}

// AsBusinessError extracts a BusinessError from wrapped errors.
func AsBusinessError(err error) (BusinessError, bool) {
	var be BusinessError
	if !errors.As(err, &be) {
		return BusinessError{}, false
	}
	return be, true
}

// RequestMeta carries request-scoped metadata used for idempotent writes.
type RequestMeta struct {
	Source    string
	RequestID string
}

// CreateAnimalInput is the application command for animal creation.
type CreateAnimalInput struct {
	Name      string
	Species   string
	Tag       string
	Birthdate string
	PhotoID   string
	Meta      RequestMeta
}

// CreateAnimalOutput is the application result for animal creation.
type CreateAnimalOutput struct {
	AnimalID  string
	EventID   string
	Replayed  bool
	Name      string
	Species   string
	Tag       string
	Birthdate string
	PhotoID   string
}

// AnimalWriter executes animal write commands in the application layer.
type AnimalWriter interface {
	Create(ctx context.Context, in CreateAnimalInput) (CreateAnimalOutput, error)
}

type createAnimalWriter struct {
	store ports.AnimalWriteStore
}

// NewCreateAnimalWriter builds the create-animal application service.
func NewCreateAnimalWriter(store ports.AnimalWriteStore) AnimalWriter {
	return createAnimalWriter{store: store}
}

func (w createAnimalWriter) Create(ctx context.Context, in CreateAnimalInput) (CreateAnimalOutput, error) {
	in = normalizeCreateAnimalInput(in)

	if err := validateCreateAnimalInput(in); err != nil {
		return CreateAnimalOutput{}, err
	}

	if in.Meta.Source == "" || in.Meta.RequestID == "" {
		return CreateAnimalOutput{}, BusinessError{
			Code: CodeInvalidInput,
			Err:  errors.New("source and request_id are required"),
		}
	}

	storeIn := ports.CreateAnimalRecordInput{
		Name:      in.Name,
		Species:   in.Species,
		Tag:       in.Tag,
		Birthdate: in.Birthdate,
		PhotoID:   in.PhotoID,
		Source:    in.Meta.Source,
		RequestID: in.Meta.RequestID,
	}

	replay, found, err := w.store.FindCreateAnimalReplay(ctx, storeIn)
	if err != nil {
		if code, ok := createAnimalConflictCode(err); ok {
			return CreateAnimalOutput{}, BusinessError{Code: code, Err: err}
		}
		return CreateAnimalOutput{}, fmt.Errorf("find create-animal replay: %w", err)
	}
	if found {
		return CreateAnimalOutput{
			AnimalID:  replay.AnimalID,
			EventID:   replay.EventID,
			Replayed:  true,
			Name:      in.Name,
			Species:   in.Species,
			Tag:       in.Tag,
			Birthdate: in.Birthdate,
			PhotoID:   in.PhotoID,
		}, nil
	}

	if in.PhotoID != "" {
		exists, err := w.store.PhotoExists(ctx, in.PhotoID)
		if err != nil {
			return CreateAnimalOutput{}, fmt.Errorf("photo exists: %w", err)
		}
		if !exists {
			return CreateAnimalOutput{}, BusinessError{
				Code: CodePhotoNotFound,
				Err:  errors.New("photo not found"),
			}
		}
	}

	out, err := w.store.CreateAnimalRecord(ctx, storeIn)
	if err != nil {
		if code, ok := createAnimalConflictCode(err); ok {
			return CreateAnimalOutput{}, BusinessError{
				Code: code,
				Err:  err,
			}
		}
		return CreateAnimalOutput{}, fmt.Errorf("create animal record: %w", err)
	}

	return CreateAnimalOutput{
		AnimalID:  out.AnimalID,
		EventID:   out.EventID,
		Replayed:  out.Replayed,
		Name:      in.Name,
		Species:   in.Species,
		Tag:       in.Tag,
		Birthdate: in.Birthdate,
		PhotoID:   in.PhotoID,
	}, nil
}

func createAnimalConflictCode(err error) (BusinessCode, bool) {
	if errors.Is(err, ports.ErrIdempotencyPayloadMismatch) {
		return CodeIdempotencyPayloadMismatch, true
	}
	if errors.Is(err, ports.ErrIdempotencyEventTypeMismatch) {
		return CodeIdempotencyEventTypeMismatch, true
	}
	if errors.Is(err, ports.ErrConflict) {
		return CodeConflict, true
	}
	return "", false
}

func normalizeCreateAnimalInput(in CreateAnimalInput) CreateAnimalInput {
	in.Name = strings.TrimSpace(in.Name)
	in.Species = strings.TrimSpace(in.Species)
	in.Tag = strings.TrimSpace(in.Tag)
	in.Birthdate = strings.TrimSpace(in.Birthdate)
	in.PhotoID = strings.TrimSpace(in.PhotoID)
	return in
}
