package utils

type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type WebOrderH2HResponse struct {
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type WebArrayResponse struct {
	Code     int                `json:"code"`
	Status   string             `json:"status"`
	Data     interface{}        `json:"data"`
	Paginate PaginationResponse `json:"paginate"`
}

type PaginationResponse struct {
	Page  int16  `json:"page"`
	Size  int16  `json:"limit"`
	Total uint32 `json:"total"`
}
