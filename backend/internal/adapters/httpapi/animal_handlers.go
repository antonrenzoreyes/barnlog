package httpapi

import (
	"log/slog"
	"net/http"

	"barnlog/backend/internal/application"
)

type animalHandlers struct {
	logger       *slog.Logger
	animalWriter application.AnimalWriter
}

func newAnimalHandlers(logger *slog.Logger, animalWriter application.AnimalWriter) animalHandlers {
	return animalHandlers{
		logger:       logger,
		animalWriter: animalWriter,
	}
}

type createAnimalRequest struct {
	Name      string `json:"name" example:"Nanny"`
	Species   string `json:"species" enums:"goat,pig,dog,cat" example:"goat"`
	Tag       string `json:"tag" example:"G-7"`
	Birthdate string `json:"birthdate" format:"date" example:"2021-03-04"`
	PhotoID   string `json:"photo_id" example:"photo_1"`
}

type createAnimalResponse struct {
	AnimalID  string `json:"animal_id" example:"animal_123"`
	EventID   string `json:"event_id" example:"event_123"`
	Name      string `json:"name" example:"Nanny"`
	Species   string `json:"species" example:"goat"`
	Tag       string `json:"tag,omitempty" example:"G-7"`
	Birthdate string `json:"birthdate,omitempty" format:"date" example:"2021-03-04"`
	PhotoID   string `json:"photo_id,omitempty" example:"photo_1"`
}

// createAnimal godoc
//
// @Summary Create animal
// @Description Creates an animal by appending an animal.created event.
// @Tags animals
// @Accept json
// @Produce json
// @Param X-Request-Id header string false "Idempotency request key (omit to disable idempotency)"
// @Param X-Barnlog-Source header string false "Request source"
// @Param request body createAnimalRequest true "Create animal payload"
// @Success 201 {object} createAnimalResponse
// @Success 200 {object} createAnimalResponse "Idempotent replay"
// @Failure 400 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Failure 413 {object} errorResponse
// @Failure 415 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /animals [post]
func (h animalHandlers) createAnimal(w http.ResponseWriter, r *http.Request) {
	var req createAnimalRequest
	if status, code, ok := decodeJSONRequest(w, r, &req); !ok {
		writeError(w, status, code)
		return
	}

	meta, ok := requestMeta(r.Context())
	if !ok {
		writeError(w, http.StatusInternalServerError, "internal_error")
		return
	}

	in := application.CreateAnimalInput{
		Name:      req.Name,
		Species:   req.Species,
		Tag:       req.Tag,
		Birthdate: req.Birthdate,
		PhotoID:   req.PhotoID,
		Meta: application.RequestMeta{
			Source:    meta.Source,
			RequestID: meta.RequestID,
		},
	}

	out, err := h.animalWriter.Create(r.Context(), in)
	if err != nil {
		if writeBusinessError(w, h.logger, err) {
			return
		}

		h.logger.Error("create animal failed", slog.Any("error", err))
		writeError(w, http.StatusInternalServerError, "internal_error")
		return
	}

	status := http.StatusCreated
	if out.Replayed {
		status = http.StatusOK
	}

	writeJSON(w, status, createAnimalResponse{
		AnimalID:  out.AnimalID,
		EventID:   out.EventID,
		Name:      out.Name,
		Species:   out.Species,
		Tag:       out.Tag,
		Birthdate: out.Birthdate,
		PhotoID:   out.PhotoID,
	})
}
