package helpers

import "pustaka-api/entity"

// PRIVATE CONVERT
func ConvertToBookResponse(book entity.Book) entity.BookResponse {
	return entity.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Subtitle:    book.Subtitle,
		Description: book.Description,
		Price:       book.Price,
		Rating:      book.Rating,
		Discount:    book.Discount,
	}
}
