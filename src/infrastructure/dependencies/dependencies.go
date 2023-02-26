package dependencies

import (
	"compreYa/src/config/database"
	usecasesAuth "compreYa/src/core/usecases/auth"
	usecasesStore "compreYa/src/core/usecases/store"
	"compreYa/src/infrastructure/entrypoints"
	entrypointsAuth "compreYa/src/infrastructure/entrypoints/api/auth"
	entrypointsStore "compreYa/src/infrastructure/entrypoints/api/store"
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
	DeleteStore            entrypoints.Handler
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

	storeRepository := &repositories.Store{
		DB: db,
	}

	sendGridRepository := &sendgridRepository.Repository{
		Client:  sendgrid.NewSendClient(os.Getenv("EMAIL_SENDGRID_API_KEY")),
		BaseURL: os.Getenv("BASE_URL"),
	}

	// use cases
	signUpUseCase := &usecasesAuth.SignUpImpl{
		AuthRepository: authRepository,
	}

	logInUseCase := &usecasesAuth.LogInImpl{
		AuthRepository:    authRepository,
		EmailNotification: sendGridRepository,
	}

	recoverPasswordRequestUseCase := &usecasesAuth.PasswordRecoveryImpl{
		AuthRepository: authRepository,
		Email:          sendGridRepository,
	}

	createStoreUseCase := &usecasesStore.CreateStoreImpl{
		StoreRepository: storeRepository,
	}

	deleteStoreUseCase := &usecasesStore.DeleteStoreImpl{
		StoreRepository: storeRepository,
	}

	// validate email sub usecase?

	// handlers
	handlers := HandlerContainer{}

	handlers.SignUp = &entrypointsAuth.SignUp{
		SignUp: signUpUseCase,
	}

	handlers.Login = &entrypointsAuth.LogIn{
		LogIn: logInUseCase,
	}

	handlers.RecoverPasswordRequest = &entrypointsAuth.RecoverPasswordRequest{
		RecoverPassword: recoverPasswordRequestUseCase,
	}

	handlers.RecoverPasswordChange = &entrypointsAuth.RecoverPasswordChange{
		RecoverPassword: recoverPasswordRequestUseCase,
	}

	handlers.AuthValidation = &middleware.Validation{}

	handlers.CreateStore = &entrypointsStore.CreateStore{
		CreateStore: createStoreUseCase,
	}

	handlers.DeleteStore = &entrypointsStore.DeleteStore{
		DeleteStore: deleteStoreUseCase,
	}

	return &handlers
}
