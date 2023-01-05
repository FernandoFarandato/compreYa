package store

type CreateStoreData struct {
	Name    string `json:"name" binding:"required"`
	URLName string `json:"url_name" binding:"required"`
}
