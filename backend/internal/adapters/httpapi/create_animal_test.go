package httpapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"barnlog/backend/internal/application"

	"github.com/go-chi/chi/v5"
)

func TestCreateAnimal(t *testing.T) {
	t.Parallel()

	t.Run("created", testCreateAnimalCreated)
	t.Run("replayed", testCreateAnimalReplayed)
	t.Run("errors", testCreateAnimalErrors)
}

func testCreateAnimalCreated(t *testing.T) {
	t.Parallel()

	uc := &fakeAnimalWriter{
		out: application.CreateAnimalOutput{
			AnimalID:  "animal_123",
			EventID:   "event_123",
			Name:      "Nanny",
			Species:   "goat",
			Tag:       "G-7",
			Birthdate: "2021-03-04",
			PhotoID:   "photo_1",
		},
	}
	router := animalTestRouter(uc)

	rec := performCreateAnimal(t, router, `{
		"name":" Nanny ",
		"species":"goat",
		"tag":" G-7 ",
		"birthdate":"2021-03-04",
		"photo_id":" photo_1 "
	}`, withCreateAnimalHeaders(
		"X-Request-Id", "req-123",
		"X-Barnlog-Source", "web.app",
	))

	assertJSONStatus(t, rec, http.StatusCreated)

	var payload map[string]any
	decodeJSON(t, rec, &payload)

	if payload["animal_id"] != "animal_123" {
		t.Fatalf("expected animal_id=animal_123, got %#v", payload["animal_id"])
	}
	if payload["event_id"] != "event_123" {
		t.Fatalf("expected event_id=event_123, got %#v", payload["event_id"])
	}
	if payload["name"] != "Nanny" {
		t.Fatalf("expected name=Nanny, got %#v", payload["name"])
	}
	if payload["species"] != "goat" {
		t.Fatalf("expected species=goat, got %#v", payload["species"])
	}
	if payload["tag"] != "G-7" {
		t.Fatalf("expected tag=G-7, got %#v", payload["tag"])
	}
	if payload["birthdate"] != "2021-03-04" {
		t.Fatalf("expected birthdate=2021-03-04, got %#v", payload["birthdate"])
	}
	if payload["photo_id"] != "photo_1" {
		t.Fatalf("expected photo_id=photo_1, got %#v", payload["photo_id"])
	}

	if uc.in.Name != " Nanny " {
		t.Fatalf("expected original name, got %q", uc.in.Name)
	}
	if uc.in.Tag != " G-7 " {
		t.Fatalf("expected original tag, got %q", uc.in.Tag)
	}
	if uc.in.PhotoID != " photo_1 " {
		t.Fatalf("expected original photo_id, got %q", uc.in.PhotoID)
	}
	if uc.in.Meta.RequestID != "req-123" {
		t.Fatalf("expected request id req-123, got %q", uc.in.Meta.RequestID)
	}
	if uc.in.Meta.Source != "web.app" {
		t.Fatalf("expected source web.app, got %q", uc.in.Meta.Source)
	}
}

func testCreateAnimalErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		body       string
		headers    map[string]string
		writerErr  error
		wantStatus int
		wantCode   string
	}{
		{
			name:       "invalid json",
			body:       `{"name":"Nanny"`,
			wantStatus: http.StatusBadRequest,
			wantCode:   "invalid_json",
		},
		{
			name: "unsupported media type",
			body: `{"name":"Nanny","species":"goat"}`,
			headers: withCreateAnimalHeaders(
				"Content-Type", "text/plain",
			),
			wantStatus: http.StatusUnsupportedMediaType,
			wantCode:   "unsupported_media_type",
		},
		{
			name:       "missing content type",
			body:       `{"name":"Nanny","species":"goat"}`,
			headers:    withCreateAnimalHeaders("Content-Type", ""),
			wantStatus: http.StatusUnsupportedMediaType,
			wantCode:   "unsupported_media_type",
		},
		{
			name:       "unknown field",
			body:       `{"name":"Nanny","species":"goat","extra":"x"}`,
			wantStatus: http.StatusBadRequest,
			wantCode:   "invalid_json",
		},
		{
			name:       "trailing json",
			body:       `{"name":"Nanny","species":"goat"}{"name":"Extra"}`,
			wantStatus: http.StatusBadRequest,
			wantCode:   "invalid_json",
		},
		{
			name:       "name required",
			body:       `{"name":" ","species":"goat"}`,
			writerErr:  businessErr(application.CodeNameRequired, "name is required"),
			wantStatus: http.StatusBadRequest,
			wantCode:   "name_required",
		},
		{
			name:       "species invalid",
			body:       `{"name":"Nanny","species":"horse"}`,
			writerErr:  businessErr(application.CodeSpeciesInvalid, "species is invalid"),
			wantStatus: http.StatusBadRequest,
			wantCode:   "species_invalid",
		},
		{
			name:       "birthdate invalid",
			body:       `{"name":"Nanny","species":"goat","birthdate":"2021-31-12"}`,
			writerErr:  businessErr(application.CodeBirthdateInvalid, "birthdate is invalid"),
			wantStatus: http.StatusBadRequest,
			wantCode:   "birthdate_invalid",
		},
		{
			name:       "use case invalid input",
			body:       `{"name":"Nanny","species":"goat"}`,
			writerErr:  businessErr(application.CodeInvalidInput, "bad data"),
			wantStatus: http.StatusBadRequest,
			wantCode:   "invalid_input",
		},
		{
			name:       "use case conflict",
			body:       `{"name":"Nanny","species":"goat"}`,
			writerErr:  businessErr(application.CodeConflict, "duplicate"),
			wantStatus: http.StatusConflict,
			wantCode:   "conflict",
		},
		{
			name:       "use case idempotency payload mismatch",
			body:       `{"name":"Nanny","species":"goat"}`,
			writerErr:  businessErr(application.CodeIdempotencyPayloadMismatch, "payload mismatch"),
			wantStatus: http.StatusConflict,
			wantCode:   "idempotency_payload_mismatch",
		},
		{
			name:       "use case idempotency event type mismatch",
			body:       `{"name":"Nanny","species":"goat"}`,
			writerErr:  businessErr(application.CodeIdempotencyEventTypeMismatch, "event type mismatch"),
			wantStatus: http.StatusConflict,
			wantCode:   "idempotency_event_type_mismatch",
		},
		{
			name:       "use case internal error",
			body:       `{"name":"Nanny","species":"goat"}`,
			writerErr:  errors.New("boom"),
			wantStatus: http.StatusInternalServerError,
			wantCode:   "internal_error",
		},
		{
			name:       "photo not found",
			body:       `{"name":"Nanny","species":"goat","photo_id":"photo_404"}`,
			writerErr:  businessErr(application.CodePhotoNotFound, "photo not found"),
			wantStatus: http.StatusBadRequest,
			wantCode:   "photo_not_found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			router := animalTestRouter(&fakeAnimalWriter{err: tc.writerErr})

			rec := performCreateAnimal(t, router, tc.body, tc.headers)
			assertJSONStatus(t, rec, tc.wantStatus)
			assertErrorCode(t, rec, tc.wantCode)
		})
	}

}

func testCreateAnimalReplayed(t *testing.T) {
	t.Parallel()

	uc := &fakeAnimalWriter{
		out: application.CreateAnimalOutput{
			AnimalID:  "animal_123",
			EventID:   "event_123",
			Replayed:  true,
			Name:      "Nanny",
			Species:   "goat",
			Tag:       "G-7",
			Birthdate: "2021-03-04",
			PhotoID:   "photo_1",
		},
	}
	router := animalTestRouter(uc)

	rec := performCreateAnimal(t, router, `{"name":"Nanny","species":"goat"}`, nil)
	assertJSONStatus(t, rec, http.StatusOK)

	var payload map[string]any
	decodeJSON(t, rec, &payload)
	if payload["animal_id"] != "animal_123" {
		t.Fatalf("expected animal_id=animal_123, got %#v", payload["animal_id"])
	}
	if payload["event_id"] != "event_123" {
		t.Fatalf("expected event_id=event_123, got %#v", payload["event_id"])
	}

	t.Run("replayed with photo_id still returns ok", func(t *testing.T) {
		t.Parallel()

		uc := &fakeAnimalWriter{
			out: application.CreateAnimalOutput{
				AnimalID:  "animal_123",
				EventID:   "event_123",
				Replayed:  true,
				Name:      "Nanny",
				Species:   "goat",
				Tag:       "G-7",
				Birthdate: "2021-03-04",
				PhotoID:   "photo_missing_now",
			},
		}
		router := animalTestRouter(uc)

		rec := performCreateAnimal(t, router, `{"name":"Nanny","species":"goat","photo_id":"photo_missing_now"}`, nil)
		assertJSONStatus(t, rec, http.StatusOK)

		var payload map[string]any
		decodeJSON(t, rec, &payload)
		if payload["photo_id"] != "photo_missing_now" {
			t.Fatalf("expected photo_id=photo_missing_now, got %#v", payload["photo_id"])
		}
	})
}

func businessErr(code application.BusinessCode, msg string) error {
	return application.BusinessError{
		Code: code,
		Err:  errors.New(msg),
	}
}

type fakeAnimalWriter struct {
	in  application.CreateAnimalInput
	out application.CreateAnimalOutput
	err error
}

func (f *fakeAnimalWriter) Create(_ context.Context, in application.CreateAnimalInput) (application.CreateAnimalOutput, error) {
	f.in = in
	return f.out, f.err
}

func animalTestRouter(writer application.AnimalWriter) http.Handler {
	r := chi.NewRouter()
	r.Use(withRequestMeta)
	animal := newAnimalHandlers(testLogger(), writer)
	r.Post("/animals", animal.createAnimal)
	return r
}

func performCreateAnimal(t *testing.T, router http.Handler, body string, headers map[string]string) *httptest.ResponseRecorder {
	t.Helper()

	req := httptest.NewRequest(http.MethodPost, "/animals", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", jsonContentType)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func withCreateAnimalHeaders(parts ...string) map[string]string {
	headers := make(map[string]string, len(parts)/2)
	for i := 0; i+1 < len(parts); i += 2 {
		headers[parts[i]] = parts[i+1]
	}
	return headers
}

func assertErrorCode(t *testing.T, rec *httptest.ResponseRecorder, expected string) {
	t.Helper()

	var payload map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}

	if payload["error"] != expected {
		t.Fatalf("expected error=%q, got %#v", expected, payload["error"])
	}
}
