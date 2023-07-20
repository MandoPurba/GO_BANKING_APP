package errors

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
