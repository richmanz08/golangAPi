package common

type StatusType string

const (
	Inactive StatusType = "INACTIVE"
	Active   StatusType = "ACTIVE"
)

type Pagination struct {
	PageSize int `json:"pageSize"`
	Current  int `json:"current"`
	Total    int `json:"total"`
}

type Response struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
	StatusCode int         `json:"statusCode"`
}