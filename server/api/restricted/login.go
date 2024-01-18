package restricted

import (
	"io"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pelletier/go-toml/v2"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type users struct {
	api_users []struct {
		username string
		password string
	}
}

var SIGNING_KEY_SECRET = GenerateSecret()

func LoadUsers() map[string]string {
	file, err := os.Open("users.toml")
	if err != nil {
		log.Panicf("unable to open users.toml: %v\n", err)
	}
	defer file.Close()

	var api_users users

	in, err := io.ReadAll(file)
	if err != nil {
		log.Panicf("unable to read users.toml: %v\n", err)
	}

	err = toml.Unmarshal(in, &api_users)
	if err != nil {
		log.Fatalf("unable to unmarshal users.toml: %v\n", err)
	}

	user_list := map[string]string{}

	for _, v := range api_users.api_users {
		user_list[v.username] = v.password
	}

	//user_list["admin"] = "password"

	return user_list
}
