package dependencies

import (
	"compreYa/src/config/database"
	usecases "compreYa/src/core/usecases/auth"
	"compreYa/src/infrastructure/entrypoints"
	"compreYa/src/infrastructure/entrypoints/api"
	"compreYa/src/infrastructure/entrypoints/api/auth"
	"compreYa/src/middleware"
	"compreYa/src/repositories/api/sendgrid"
	repositories "compreYa/src/repositories/db"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type HandlerContainer struct {
	SignUp         entrypoints.Handler
	Login          entrypoints.Handler
	AuthValidation entrypoints.Handler
	CreateStore    entrypoints.Handler
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

	sendGridRepository := &sendgrid.Repository{
		Client: sendgrid.NewSendClient(os.Getenv("EMAIL_SENDGRID_API_KEY")),
	}

	// use cases
	signUpUseCase := &usecases.SignUpImpl{
		AuthRepository: authRepository,
	}

	logInUseCase := &usecases.LogInImpl{
		AuthRepository:    authRepository,
		EmailNotification: sendGridRepository,
	}

	// handlers
	handlers := HandlerContainer{}

	handlers.SignUp = &auth.SignUp{
		SignUp: signUpUseCase,
	}

	handlers.Login = &auth.LogIn{
		LogIn: logInUseCase,
	}

	handlers.AuthValidation = &middleware.Validation{}

	handlers.CreateStore = &api.CreateStore{}

	return &handlers
}
