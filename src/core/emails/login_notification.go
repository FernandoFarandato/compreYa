package emails

type LoginNotificationData struct {
	Email    string
	UserName string
}

func NewLoginNotificationData(email, username string) *LoginNotificationData {
	return &LoginNotificationData{
		Email:    email,
		UserName: username,
	}
}
