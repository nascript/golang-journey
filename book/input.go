package book

type BookInput struct {
	Title    string      `json:"title" binding:"required"`
	Price    interface{} `json:"price" binding:"required,number"`
	Subtitle string      `json:"sub_title"`
}
