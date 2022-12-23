package repositories

func (r *repository) Delete(book Book) (Book, error) {

	err := r.db.Delete(&book).Error

	return book, err
}
