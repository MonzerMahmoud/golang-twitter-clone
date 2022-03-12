package helpers

import (
	"fmt"
	"golang-twitter-clone/interfaces"
	"os"
	"regexp"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
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
	//db, err := gorm.Open("postgres", "host=ec2-54-158-26-89.compute-1.amazonaws.com port=5432 user=mchmkgthfrusfo dbname=das0j74cai8cer sslmode=disable password=d8c321a53c3d4e617295b89bb6ad45a5a95ab3c8850f122f8ac71bc8e6bef1ef")
	DBURL := os.Getenv("DATABASE_URL")
	if DBURL == "" {
		DBURL = "host=localhost port=5432 user=postgres dbname=Twitter-Clone sslmode=disable"
	}
	db, err := gorm.Open("postgres", DBURL)
	HandleErr(err)
	return db
}

func Validation(values []interfaces.Validation) bool{
    fullName := regexp.MustCompile(`^([A-Za-z0-9]{5,})+$`)
    email := regexp.MustCompile(`^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$`)

    for i := 0; i < len(values); i++ {
        switch values[i].Valid {
            case "fullName":
                if !fullName.MatchString(values[i].Value) {
                    return false
                }
            case "email":
                if !email.MatchString(values[i].Value) {
                    return false
                }
            case "password":
                if len(values[i].Value) < 5 {
                    return false
                }
        }
    }
    return true
}