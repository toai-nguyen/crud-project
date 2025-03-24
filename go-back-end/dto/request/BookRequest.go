package request

type CreateBookRequest struct {
	Title    string `json:"title"`
	Price    uint   `json:"price"`
	Quantity uint   `json:"quantity"`
	Status   string `json:"status"`
	AuthorID uint   `json:"author_id"`
}
type UpdateBookRequest struct {
	Title    string `json:"title"`
	Price    uint   `json:"price"`
	Quantity uint   `json:"quantity"`
	Status   string `json:"status"`
}
