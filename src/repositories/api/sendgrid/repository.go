package sendgrid

import (
	"bytes"
	"compreYa/src/core/emails"
	"compreYa/src/core/entities"
	"compreYa/src/core/errors"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"html/template"
	"log"
	"os"
)

const (
	loginNotificationFile  = "src/core/emails/login_notification.html"
	signUpNotificationFile = "src/core/emails/signup_notification.html"
	passwordRecoveryFile   = "src/core/emails/recover_password.html"
)

type Repository struct {
	Client  *sendgrid.Client
	BaseURL string
}

func (r *Repository) SendLoginEmail(userData *entities.User) *errors.ApiError {
	htmlData := emails.NewLoginNotificationData(userData.Email, userData.UserName)
	htmlBody, err := r.getHTMLBody(htmlData, loginNotificationFile)
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

func (r *Repository) SendPasswordRecovery(userData *entities.User, token string) *errors.ApiError {
	link := fmt.Sprintf("%s/compreYa/auth/recover_password/change?recovery_token=%s", r.BaseURL, token)
	htmlData := emails.NewRecoverPasswordData(link)
	htmlBody, err := r.getHTMLBody(htmlData, passwordRecoveryFile)
	if err != nil {
		return nil
	}

	emailData := entities.NewEmailInformation(userData.UserName, entities.RecoverPassword.String(), userData.Email, htmlBody)
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

func (r *Repository) getHTMLBody(emailData interface{}, htmlFileAddress string) (*string, *errors.ApiError) {
	var htmlTemplate *template.Template

	htmlTemplate, err := htmlTemplate.ParseFiles(htmlFileAddress)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, errors.CreatingEmailError)
	}

	buffer := new(bytes.Buffer)
	err = htmlTemplate.Execute(buffer, emailData)
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
