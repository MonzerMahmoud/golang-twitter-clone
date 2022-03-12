package helpers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func HandleErr(err error) {
	if err != nil {
		// TODO: check if using panic is a good idea
		fmt.Println("There is error here")
		panic(err.Error())
	}
}

func HashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

func ConnectDB() *gorm.DB {
	fmt.Println("Connecting to DB")
	//db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=Twitter-Clone sslmode=disable")
	db, err := gorm.Open("postgres", "host=ec2-54-158-26-89.compute-1.amazonaws.com port=5432 user=mchmkgthfrusfo dbname=das0j74cai8cer sslmode=disable password=d8c321a53c3d4e617295b89bb6ad45a5a95ab3c8850f122f8ac71bc8e6bef1ef")
	HandleErr(err)
	return db
}