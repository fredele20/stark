package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	inputError"stark/errors"
	"net/http"
	"stark/model"
	"stark/service"
)

type UserApi interface {
	Apply(w http.ResponseWriter, r *http.Request)
	Applicant(w http.ResponseWriter, r *http.Request)
	UploadFile(w http.ResponseWriter, r *http.Request)
}

type api struct {
	userService service.UserService
}

func NewUserApi(service service.UserService) UserApi { return &api{service} }

func (a *api) Apply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users model.User
	var works model.EmployerHistory

	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errors.New("error decoding data")
		return
	}

	err = a.userService.Validate(&users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(inputError.ServiceError{Message: err.Error()})
		return
	}

	err = a.userService.ValidateTwo(&works)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(inputError.ServiceError{Message: err.Error()})
		return
	}

	meetup, err := a.userService.Apply(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errors.New("error creating a meetup")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(meetup)
}

func (a *api) Applicant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := a.userService.GetApplications()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (a *api) UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

