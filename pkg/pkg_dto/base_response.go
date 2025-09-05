package pkg_dto

type IdResponse struct {
	ID string `json:"id"`
}

type BaseResponse struct {
	Data      interface{} `json:"data"`
	TraceID   string      `json:"trace_id"`
	Succeeded bool        `json:"succeeded"`
	Errors    []string    `json:"errors"`
}

type ArrayResponse struct {
	BaseResponse
	Page       string `json:"page"`
	Size       string `json:"size"`
	TotalItems string `json:"total_items"`
}
