package application

import (
	"context"
	"testing"

	"barnlog/backend/internal/ports"
)

func TestCreateAnimalWriter_PhotoNotFound(t *testing.T) {
	t.Parallel()

	w := NewCreateAnimalWriter(&fakeAnimalWriteStore{
		photoExists: false,
	})

	_, err := w.Create(context.Background(), CreateAnimalInput{
		Name:    "Nanny",
		Species: "goat",
		PhotoID: "p1",
		Meta: RequestMeta{
			Source:    "test",
			RequestID: "req-1",
		},
	})
	if err == nil {
		t.Fatalf("expected error")
	}
	be, ok := AsBusinessError(err)
	if !ok {
		t.Fatalf("expected business error, got %T", err)
	}
	if be.Code != CodePhotoNotFound {
		t.Fatalf("expected code %q, got %q", CodePhotoNotFound, be.Code)
	}
}

func TestCreateAnimalWriter_ValidationErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   CreateAnimalInput
		code BusinessCode
	}{
		{
			name: "name required",
			in: CreateAnimalInput{
				Name:    "",
				Species: "goat",
				Meta: RequestMeta{
					Source:    "test",
					RequestID: "req-1",
				},
			},
			code: CodeNameRequired,
		},
		{
			name: "species invalid",
			in: CreateAnimalInput{
				Name:    "Nanny",
				Species: "horse",
				Meta: RequestMeta{
					Source:    "test",
					RequestID: "req-1",
				},
			},
			code: CodeSpeciesInvalid,
		},
		{
			name: "birthdate invalid",
			in: CreateAnimalInput{
				Name:      "Nanny",
				Species:   "goat",
				Birthdate: "2021-31-12",
				Meta: RequestMeta{
					Source:    "test",
					RequestID: "req-1",
				},
			},
			code: CodeBirthdateInvalid,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			w := NewCreateAnimalWriter(&fakeAnimalWriteStore{
				photoExists: true,
			})
			_, err := w.Create(context.Background(), tc.in)
			if err == nil {
				t.Fatalf("expected error")
			}
			be, ok := AsBusinessError(err)
			if !ok {
				t.Fatalf("expected business error, got %T", err)
			}
			if be.Code != tc.code {
				t.Fatalf("expected code %q, got %q", tc.code, be.Code)
			}
		})
	}
}

func TestCreateAnimalWriter_Conflict(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		code BusinessCode
	}{
		{name: "generic conflict", err: ports.ErrConflict, code: CodeConflict},
		{name: "payload mismatch", err: ports.ErrIdempotencyPayloadMismatch, code: CodeIdempotencyPayloadMismatch},
		{name: "event type mismatch", err: ports.ErrIdempotencyEventTypeMismatch, code: CodeIdempotencyEventTypeMismatch},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			w := NewCreateAnimalWriter(&fakeAnimalWriteStore{
				photoExists: true,
				createErr:   tc.err,
			})

			_, err := w.Create(context.Background(), CreateAnimalInput{
				Name:    "Nanny",
				Species: "goat",
				PhotoID: "p1",
				Meta: RequestMeta{
					Source:    "test",
					RequestID: "req-1",
				},
			})
			if err == nil {
				t.Fatalf("expected error")
			}
			be, ok := AsBusinessError(err)
			if !ok {
				t.Fatalf("expected business error, got %T", err)
			}
			if be.Code != tc.code {
				t.Fatalf("expected code %q, got %q", tc.code, be.Code)
			}
		})
	}
}

func TestCreateAnimalWriter_Replayed(t *testing.T) {
	t.Parallel()

	w := NewCreateAnimalWriter(&fakeAnimalWriteStore{
		photoExists: true,
		createOut: ports.CreateAnimalRecordOutput{
			AnimalID: "a1",
			EventID:  "e1",
			Replayed: true,
		},
	})

	out, err := w.Create(context.Background(), CreateAnimalInput{
		Name:    "Nanny",
		Species: "goat",
		Meta: RequestMeta{
			Source:    "test",
			RequestID: "req-1",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Replayed {
		t.Fatalf("expected replayed output")
	}
}

func TestCreateAnimalWriter_ReplayedBeforePhotoExistsCheck(t *testing.T) {
	t.Parallel()

	w := NewCreateAnimalWriter(&fakeAnimalWriteStore{
		photoExists: false,
		replayFound: true,
		replayOut: ports.CreateAnimalRecordOutput{
			AnimalID: "a1",
			EventID:  "e1",
			Replayed: true,
		},
	})

	out, err := w.Create(context.Background(), CreateAnimalInput{
		Name:    "Nanny",
		Species: "goat",
		PhotoID: "photo_1",
		Meta: RequestMeta{
			Source:    "test",
			RequestID: "req-1",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Replayed {
		t.Fatalf("expected replayed output")
	}
}

func TestCreateAnimalWriter_NormalizesInput(t *testing.T) {
	t.Parallel()

	store := &fakeAnimalWriteStore{
		photoExists: true,
	}
	w := NewCreateAnimalWriter(store)

	out, err := w.Create(context.Background(), CreateAnimalInput{
		Name:      " Nanny ",
		Species:   " goat ",
		Tag:       " G-7 ",
		Birthdate: " 2021-03-04 ",
		PhotoID:   " photo_1 ",
		Meta: RequestMeta{
			Source:    "test",
			RequestID: "req-1",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if store.createIn.Name != "Nanny" {
		t.Fatalf("expected trimmed Name, got %q", store.createIn.Name)
	}
	if store.createIn.Species != "goat" {
		t.Fatalf("expected trimmed Species, got %q", store.createIn.Species)
	}
	if store.createIn.Tag != "G-7" {
		t.Fatalf("expected trimmed Tag, got %q", store.createIn.Tag)
	}
	if store.createIn.Birthdate != "2021-03-04" {
		t.Fatalf("expected trimmed Birthdate, got %q", store.createIn.Birthdate)
	}
	if store.createIn.PhotoID != "photo_1" {
		t.Fatalf("expected trimmed PhotoID, got %q", store.createIn.PhotoID)
	}

	if out.Name != "Nanny" {
		t.Fatalf("expected normalized output Name, got %q", out.Name)
	}
	if out.Species != "goat" {
		t.Fatalf("expected normalized output Species, got %q", out.Species)
	}
	if out.Tag != "G-7" {
		t.Fatalf("expected normalized output Tag, got %q", out.Tag)
	}
	if out.Birthdate != "2021-03-04" {
		t.Fatalf("expected normalized output Birthdate, got %q", out.Birthdate)
	}
	if out.PhotoID != "photo_1" {
		t.Fatalf("expected normalized output PhotoID, got %q", out.PhotoID)
	}
}

type fakeAnimalWriteStore struct {
	photoExists bool
	photoErr    error
	createErr   error
	createIn    ports.CreateAnimalRecordInput
	createOut   ports.CreateAnimalRecordOutput
	replayErr   error
	replayOut   ports.CreateAnimalRecordOutput
	replayFound bool
	replayIn    ports.CreateAnimalRecordInput
}

func (f *fakeAnimalWriteStore) CreateAnimalRecord(_ context.Context, in ports.CreateAnimalRecordInput) (ports.CreateAnimalRecordOutput, error) {
	f.createIn = in
	if f.createErr != nil {
		return ports.CreateAnimalRecordOutput{}, f.createErr
	}
	if f.createOut != (ports.CreateAnimalRecordOutput{}) {
		return f.createOut, nil
	}
	return ports.CreateAnimalRecordOutput{
		AnimalID: "a1",
		EventID:  "e1",
	}, nil
}

func (f *fakeAnimalWriteStore) FindCreateAnimalReplay(_ context.Context, in ports.CreateAnimalRecordInput) (ports.CreateAnimalRecordOutput, bool, error) {
	f.replayIn = in
	if f.replayErr != nil {
		return ports.CreateAnimalRecordOutput{}, false, f.replayErr
	}
	if f.replayFound {
		return f.replayOut, true, nil
	}
	return ports.CreateAnimalRecordOutput{}, false, nil
}

func (f *fakeAnimalWriteStore) PhotoExists(context.Context, string) (bool, error) {
	if f.photoErr != nil {
		return false, f.photoErr
	}
	return f.photoExists, nil
}

var _ ports.AnimalWriteStore = (*fakeAnimalWriteStore)(nil)
