package servers

import (
	"database/sql"

	"job-application/controllers"
	"job-application/helpers"
	"job-application/helpers/logger"
	"job-application/repositories"
	"job-application/services"
)

type HandlerOps struct {
	UserController *controllers.UserController
	JobController  *controllers.JobController
}

func SetupController(db *sql.DB) *HandlerOps {
	logrusLogger := logger.NewLogger()
	logger.SetLogger(logrusLogger)

	bcrypt := helpers.NewBcryptStruct()
	jwt := helpers.NewJWTProviderHS256()

	transactionsRepository := repositories.NewTransactionRepositoryImpelementation(db)
	userRepository := repositories.NewUserRepositoryImplementation(db)
	jobRepository := repositories.NewJobRepositoryImplementation(db)
	jobApplicantRepository := repositories.NewJobApplicantRepositoryImplementation(db)

	userService := services.NewUserServiceImplementation(userRepository, transactionsRepository, bcrypt, jwt)
	jobService := services.NewJobServiceImplementation(userRepository, jobRepository, jobApplicantRepository, transactionsRepository)

	userController := controllers.NewUserController(userService)
	jobController := controllers.NewJobController(jobService)

	return &HandlerOps{
		UserController: userController,
		JobController:  jobController,
	}
}
