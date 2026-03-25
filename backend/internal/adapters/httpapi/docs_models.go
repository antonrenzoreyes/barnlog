package httpapi

type statusResponse struct {
	Status string `json:"status"`
}

type readyResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type uploadFileResponse struct {
	FileID      string `json:"file_id"`
	FileName    string `json:"file_name"`
	ContentType string `json:"content_type"`
	SizeBytes   int64  `json:"size_bytes"`
}
