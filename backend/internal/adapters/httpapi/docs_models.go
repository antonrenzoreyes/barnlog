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

type uploadFileResponse struct {
	FileID      string `json:"file_id" example:"file_123" binding:"required"`
	FileName    string `json:"file_name" example:"pepper.png" binding:"required"`
	ContentType string `json:"content_type" example:"image/png" binding:"required"`
	SizeBytes   int64  `json:"size_bytes" example:"248123" binding:"required"`
}
