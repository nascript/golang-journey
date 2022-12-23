package entity

type AddBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Subtitle    string `json:"subtitle" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
	Rating      int64  `json:"rating" binding:"required"`
	Discount    int64  `json:"discount" binding:"required"`
}
