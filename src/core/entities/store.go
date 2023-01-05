package entities

type Store struct {
	Name    string
	URLName string
	OwnerID *int64
}

func NewStore(name, urlName string, ownerID *int64) *Store {
	return &Store{
		Name:    name,
		URLName: urlName,
		OwnerID: ownerID,
	}
}
