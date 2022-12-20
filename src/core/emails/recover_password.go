package emails

type RecoverPasswordData struct {
	ChangePasswordLink string
}

func NewRecoverPasswordData(recoverPasswordLink string) *RecoverPasswordData {
	return &RecoverPasswordData{
		ChangePasswordLink: recoverPasswordLink,
	}
}
