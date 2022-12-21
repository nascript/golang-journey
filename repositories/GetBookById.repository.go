package repositories

func (r *repository) FindByID(ID int) (Book, error) {
	var book Book
	err := r.db.Find(&book, ID).Error

	return book, err
}
