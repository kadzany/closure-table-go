package dto

type ApiResponseSuccess struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ApiResponseFail struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ApiResponseError struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrorCode string      `json:"error_code"`
	Errors    interface{} `json:"errors"`
}
