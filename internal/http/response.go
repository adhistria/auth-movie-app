package http

// SuccessResponse ..
type SuccessResponse struct {
	Message string      `json:"string"`
	Data    interface{} `json:"data,omitempty"`
}
