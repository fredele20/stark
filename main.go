package main

import (
	"fmt"
	"net/http"
	"os"
	"stark/api"
	"stark/database"
	"stark/domain"
	"stark/router"
	"stark/service"
)

var (
	userDatabase database.User        = database.NewUserDB()
	userDomain   domain.UserDomain    = domain.NewUserDomain(userDatabase)
	userService  service.UserService = service.NewUserService(userDomain)
	userApi      api.UserApi          = api.NewMeetupApi(userService)
	httpRouter   router.Router        = router.NewMuxRouter()
)

func main() {
	port := os.Getenv("APP_ENV")
	defaultPort := "8080"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "up and running")
	})

	httpRouter.POST("/apply", userApi.Apply)
	httpRouter.GET("/applicant", userApi.Applicant)
	httpRouter.POST("/upload", userApi.UploadFile)

	if !(port == "") {
		httpRouter.SERVER(":"+port)
	} else {
		httpRouter.SERVER(":"+defaultPort)
	}


}

