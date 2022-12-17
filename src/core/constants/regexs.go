package constants

type Regex string

const (
	Email    Regex = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
	Password Regex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
)

func (r Regex) String() string {
	return string(r)
}
