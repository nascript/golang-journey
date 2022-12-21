package services

func (s *service) Delete(ID int) (Book, error) {

	book, err := s.repository.FindByID(ID)

	bookDeleted, err := s.repository.Delete(book)

	return bookDeleted, err
}
