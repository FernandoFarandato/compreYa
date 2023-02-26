package store

type HideStoreData struct {
	Status  string `json:"status" binding:"required"`
	URLName string `json:"url_name" binding:"required"`
}
