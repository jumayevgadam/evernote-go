package abstract

// For responsing json in beautiful mode.
type SuccessResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
