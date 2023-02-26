package entities

type Store struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	URLName string `db:"url_name"`
	OwnerID int64  `db:"owner_id"`
}

func NewStore(name, urlName string, ownerID *int64) *Store {
	return &Store{
		Name:    name,
		URLName: urlName,
		OwnerID: *ownerID,
	}
}
