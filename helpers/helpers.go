package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-twitter-clone/interfaces"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	redis "github.com/go-redis/redis/v8"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var cache = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

var ctx = context.Background()

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

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)

				resp := interfaces.ErrResponse{Message: "Internal server error"}
				json.NewEncoder(w).Encode(resp)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func ValidateToken(id string, jwtToken string) bool {
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})
	HandleErr(err)

	var userId, _ = strconv.ParseFloat(id, 8)
	if token.Valid && tokenData["user_id"] == userId {
		return true
	} else {
		return false
	}
}

func ValidateTokenExp(jwtToken string) bool {
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})
	HandleErr(err)

	if token.Valid {
		return true
	} else {
		return false
	}
}

func GetUserIdFromToken(jwtToken string) string {
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	
	_, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})
	HandleErr(err)

	return strconv.Itoa(int(tokenData["user_id"].(float64)))
}

func SetTweetCache(tweet interfaces.Tweet) {
	cacheErr := cache.Set(ctx, strconv.Itoa(int(tweet.ID)), tweet, time.Second*5).Err()
	HandleErr(cacheErr)
}

func GetTweetCache(id string) interfaces.Tweet {
	var tweet interfaces.Tweet
	cacheErr := cache.Get(ctx, id).Scan(&tweet)
	HandleErr(cacheErr)

	return tweet
}

func ValidateCache(id string) bool {
	var tweet interfaces.Tweet
	cacheErr := cache.Get(ctx, id).Scan(&tweet)
	HandleErr(cacheErr)

	if tweet.ID == 0 {
		return false
	} else {
		return true
	}
}