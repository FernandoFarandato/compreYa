package entities

type Store struct {
	Name    string `json:"name"`
	URLName string `json:"url_name"`
	OwnerID *int64 `json:"owner_id"`
}

func NewStore(name, urlName string, ownerID *int64) *Store {
	return &Store{
		Name:    name,
		URLName: urlName,
		OwnerID: ownerID,
	}
}
