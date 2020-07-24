package database

import (
	"context"
	//"fmt"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo/gridfs"
	//"io/ioutil"
	"log"
	//"os"
	"stark/model"
)

type User interface {
	GetUserByField(field, value string) (*model.User, error)
	GetUserById(id string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	//UploadFile(file string)
}

type db struct {
}

func NewUserDB() User { return &db{} }

//func (d *db) UploadFile(file string) string {
//	var filename string
//
//	data, err := ioutil.ReadFile(file)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	bucket, err := gridfs.NewBucket(Collection.Database())
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

func (d *db) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	cursor, err := Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatalf("could not get all users %v", err)
	}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var user *model.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}

func (d *db) GetUserByField(field, value string) (*model.User, error) {
	var user model.User
	err := Collection.FindOne(context.TODO(), bson.M{field: value}).Decode(&user)
	return &user, err
}

func (d *db) GetUserById(id string) (*model.User, error) {
	return d.GetUserByField("_id", id)
}

func (d *db) GetUserByEmail(email string) (*model.User, error) {
	return d.GetUserByField("email", email)
}

func (d *db) CreateUser(user *model.User) (*model.User, error) {
	_, err := Collection.InsertOne(context.TODO(), user)
	return user, err
}

