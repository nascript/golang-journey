package entity

type UpdateBookRequest struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Rating      int64  `json:"rating"`
	Discount    int64  `json:"discount"`
}
