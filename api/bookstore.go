package api

type BookRequest struct {
	Name     string `json:name`
	Quantity int64  `json:quantity`
}

type BookCreateResponse struct {
	*BookRequest
	ID string `json:id`
}

type UpdateBookRequest struct {
	Name     string `json:name`
	Quantity int64  `json:quantity`
}
