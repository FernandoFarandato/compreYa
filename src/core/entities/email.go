package entities

type EmailType string

const (
	LoginNotification  EmailType = "Login Notification"
	SignUpNotification EmailType = "Login Notification"
	RecoverPassword    EmailType = "Login Notification"
)

type EmailInformation struct {
	FromName         string
	UserName         string
	Subject          string
	RecipientAddress string
	TextContent      *string
	HtmlBody         *string
}

func (e *EmailType) String() string {
	return e.String()
}

func NewEmailInformation(userName, subject, recipientAddress string, htmlBody *string) *EmailInformation {
	return &EmailInformation{
		FromName:         "compreYa",
		UserName:         userName,
		Subject:          subject,
		RecipientAddress: recipientAddress,
		HtmlBody:         htmlBody,
	}
}
