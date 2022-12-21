package services

func (s *service) Create(bookRequest BookRequest) (Book, error) {

	book := Book{
		Title:       bookRequest.Title,
		Subtitle:    bookRequest.Subtitle,
		Description: bookRequest.Description,
		Price:       int(bookRequest.Price),
		Rating:      int(bookRequest.Rating),
		Discount:    int(bookRequest.Discount),
	}
	newBook, err := s.repository.Create(book)

	return newBook, err
}
