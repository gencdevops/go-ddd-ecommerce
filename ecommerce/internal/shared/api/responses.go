package api

type ErrorResponse struct {
    Error string `json:"error"`
}

type SuccessResponse struct {
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
} 