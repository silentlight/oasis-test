package utils

type ApiResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ApiError struct {
	Message string `json:"message"`
}
