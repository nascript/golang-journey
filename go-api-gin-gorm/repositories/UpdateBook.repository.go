package repositories


func (r *repository) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error

	return book, err
}
