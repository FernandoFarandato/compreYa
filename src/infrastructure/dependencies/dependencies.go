package dependencies

import (
	"compreYa/src/config/database"
	usecases "compreYa/src/core/usecases/auth"
	"compreYa/src/infrastructure/entrypoints"
	"compreYa/src/infrastructure/entrypoints/api"
	"compreYa/src/infrastructure/entrypoints/api/auth"
	"compreYa/src/middleware"
	sendgridRepository "compreYa/src/repositories/api/sendgrid"
	repositories "compreYa/src/repositories/db"
	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"os"
)

type HandlerContainer struct {
	SignUp                 entrypoints.Handler
	Login                  entrypoints.Handler
	AuthValidation         entrypoints.Handler
	RecoverPasswordRequest entrypoints.Handler
	RecoverPasswordChange  entrypoints.Handler
	CreateStore            entrypoints.Handler
}

func Start() *HandlerContainer {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection
	db := database.Connect()

	// repositories
	authRepository := &repositories.Auth{
		DB: db,
	}

	sendGridRepository := &sendgridRepository.Repository{
		Client:  sendgrid.NewSendClient(os.Getenv("EMAIL_SENDGRID_API_KEY")),
		BaseURL: os.Getenv("BASE_URL"),
	}

	// use cases
	signUpUseCase := &usecases.SignUpImpl{
		AuthRepository: authRepository,
	}

	logInUseCase := &usecases.LogInImpl{
		AuthRepository:    authRepository,
		EmailNotification: sendGridRepository,
	}

	recoverPasswordRequestUseCase := &usecases.PasswordRecoveryImpl{
		AuthRepository: authRepository,
		Email:          sendGridRepository,
	}

	// validate email sub usecase?

	// handlers
	handlers := HandlerContainer{}

	handlers.SignUp = &auth.SignUp{
		SignUp: signUpUseCase,
	}

	handlers.Login = &auth.LogIn{
		LogIn: logInUseCase,
	}

	handlers.RecoverPasswordRequest = &auth.RecoverPasswordRequest{
		RecoverPassword: recoverPasswordRequestUseCase,
	}

	handlers.RecoverPasswordChange = &auth.RecoverPasswordChange{
		RecoverPassword: recoverPasswordRequestUseCase,
	}

	handlers.AuthValidation = &middleware.Validation{}

	handlers.CreateStore = &api.CreateStore{}

	return &handlers
}
