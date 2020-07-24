package domain

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	//"log"
	"stark/database"
	"stark/model"
)

type UserDomain interface {
	GetUsers() ([]*model.User, error)
	CreateUser(input model.User) (*model.User, error)
}

type domain struct {
	db database.User
}

func NewUserDomain(dbs database.User) UserDomain { return &domain{db: dbs} }

func (d *domain) GetUsers() ([]*model.User, error) {
	return d.db.GetAllUsers()
}

var (
	s3session *s3.S3
)

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})))
}

func (d *domain) CreateUser(input model.User) (*model.User, error) {

	//var file string
	//d.db.UploadFile(file)

	user := &model.User{
		Firstname:   input.Firstname,
		Middlename:  input.Middlename,
		Lastname:    input.Lastname,
		Email:       input.Email,
		//CV:			 UploadFile(file, filename),
		WorkHistory: input.WorkHistory,
	}
	return d.db.CreateUser(user)
}

//func UploadFile(file, filename string) string {
//	//var filename string
//
//	data, err := ioutil.ReadFile(file)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	bucket, err := gridfs.NewBucket(database.Collection.Database())
//	if err != nil {
//		log.Fatal(err)
//		os.Exit(1)
//	}
//
//
//	uploadStream, err := bucket.OpenUploadStream(
//		filename,
//	)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	defer uploadStream.Close()
//
//	fileSize, err := uploadStream.Write(data)
//	if err != nil {
//		log.Fatal(err)
//		os.Exit(1)
//	}
//	log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)
//	return file
//}

