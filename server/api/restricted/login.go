package restricted

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/pelletier/go-toml/v2"
)

type jwtCustomClaims struct {
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

func loadUsers() map[string]string {
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

	return user_list
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	users := loadUsers()
	pwd := users[username]

	// Throws unauthorized error
	if pwd == "" || password != pwd {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		username,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 60)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
