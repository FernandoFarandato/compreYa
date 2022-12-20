package sendgrid

import (
	"bytes"
	"compreYa/src/core/entities"
	"compreYa/src/core/errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"html/template"
	"log"
	"os"
)

type Repository struct {
	Client *sendgrid.Client
}

func (r *Repository) SendLoginEmail(userData *entities.User) *errors.ApiError {
	htmlBody, err := r.getHTMLBody(userData)
	if err != nil {
		return nil
	}

	emailData := entities.NewEmailInformation(userData.UserName, entities.LoginNotification.String(), userData.Email, htmlBody)
	email := r.getEmail(emailData)

	err = r.sendEmail(email)
	if err != nil {
		return errors.NewInternalServerError(nil, errors.SendingEmailError)
	}

	return nil
}

func (r *Repository) getEmail(emailData *entities.EmailInformation) *mail.SGMailV3 {
	from := mail.NewEmail(emailData.FromName, os.Getenv("EMAIL_SENDER"))
	to := mail.NewEmail(emailData.UserName, emailData.RecipientAddress)

	email := mail.NewSingleEmail(from, emailData.Subject, to, "", *emailData.HtmlBody)
	return email
}

func (r *Repository) getHTMLBody(userData *entities.User) (*string, *errors.ApiError) {
	var htmlTemplate *template.Template

	htmlTemplate, err := htmlTemplate.ParseFiles("src/core/constants/emails/login_notification.html")
	if err != nil {
		return nil, errors.NewInternalServerError(nil, errors.CreatingEmailError)
	}

	buffer := new(bytes.Buffer)
	err = htmlTemplate.Execute(buffer, userData)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, errors.CreatingEmailError)
	}

	stringHtml := buffer.String()
	return &stringHtml, nil
}

func (r *Repository) sendEmail(email *mail.SGMailV3) *errors.ApiError {
	response, err := r.Client.Send(email)
	if err != nil {
		log.Println(err)
		return errors.NewInternalServerError(nil, errors.SendingEmailError)
	}
	print(response)
	return nil
}
