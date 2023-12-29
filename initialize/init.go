package initialize

import (
	"go-jti/config"
	"go-jti/controllers/phone_number_controller"
	"go-jti/controllers/user_controller"
	"go-jti/repositories/phone_number_repository"
	"go-jti/services/phone_number_service"
	"log"
	"time"

	"github.com/joho/godotenv"
)

//Phone Number
var phoneNumberRepository phone_number_repository.PhoneNumberRepository
var phoneNumberService phone_number_service.PhoneNumberService
var PhoneNumberController phone_number_controller.PhoneNumberController

//User
var UserController user_controller.UserController

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Println(err.Error())
	}
	time.Local = location
	config.GoogleConfig()

	initRepositories()
	initServices()
	initControllers()
}

func initRepositories() {
	db := connectDatabase()

	phoneNumberRepository = phone_number_repository.NewPhoneNumberRepository(db)
}

func initServices() {
	phoneNumberService = phone_number_service.NewPhoneNumberService(phoneNumberRepository)
}

func initControllers() {
	PhoneNumberController = phone_number_controller.NewPhoneNumberController(phoneNumberService)
	UserController = user_controller.NewUserController()
}
