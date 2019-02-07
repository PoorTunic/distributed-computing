package dao

// GetDao method
func GetDao() (DataDao, error) {
	return NewDataDAOMock(), nil
}
