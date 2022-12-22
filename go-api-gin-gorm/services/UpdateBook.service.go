package services

func (s *service) Update(ID int, bookRequest UpdateBookRequest) (Book, error) {

	book, err := s.repository.FindByID(ID)

	book.Title = bookRequest.Title
	book.Subtitle = bookRequest.Subtitle
	book.Description = bookRequest.Description
	book.Price = int(bookRequest.Price)
	book.Rating = int(bookRequest.Rating)
	book.Discount = int(bookRequest.Discount)

	newBook, err := s.repository.Update(book)

	return newBook, err
}
