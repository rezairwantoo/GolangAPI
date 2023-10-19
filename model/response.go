package model

type ResponseDataCreate struct {
	Status bool `json:"status"`
}

type CreateResponse struct {
	Message string             `json:"message"`
	Data    ResponseDataCreate `json:"data"`
}

type DetailResponse struct {
	Message string   `json:"message"`
	Data    Products `json:"data"`
}

type Pagination struct {
	Total        int64
	Page         int64
	TotalAllData int64 `json:"total_all_data"`
}

type ListData struct {
	Data       []Products
	Pagination Pagination
}

type ListResponse struct {
	Message string
	Data    ListData
}

type UpdateResponse struct {
	Message string             `json:"message"`
	Data    ResponseDataCreate `json:"data"`
}

type DeleteResponse struct {
	Message string             `json:"message"`
	Data    ResponseDataCreate `json:"data"`
}
