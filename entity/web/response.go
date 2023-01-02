package web

type WebResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ErrorResponse struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	Description any    `json:"description"`
}
