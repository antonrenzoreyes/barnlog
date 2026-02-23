package httpapi

type statusResponse struct {
	Status string `json:"status" example:"ok"`
}

type readyResponse struct {
	Status    string `json:"status" example:"ready"`
	Timestamp string `json:"timestamp" format:"date-time" example:"2026-02-22T20:32:13Z"`
}

type errorResponse struct {
	Error string `json:"error" example:"invalid_json"`
}

type uploadPhotoResponse struct {
	PhotoID     string `json:"photo_id" example:"photo_123"`
	ContentType string `json:"content_type" example:"image/jpeg"`
	SizeBytes   int64  `json:"size_bytes" example:"248123"`
}
