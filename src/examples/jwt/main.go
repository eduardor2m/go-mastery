package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"

	httpSwagger "github.com/swaggo/http-swagger"

	database "github.com/eduardor2m/go-mastery/src/examples/jwt/data"
	_ "github.com/eduardor2m/go-mastery/src/examples/jwt/docs"
)

// @title Example API JWT
// @version 1.0.0
// @contact.name Eduardo Melo
// @description This is a example API JWT
// @host localhost:8080
// @BasePath /api/v1
func main() {
	c := chi.NewRouter()

	// base path api/v1

	c.Route("/api/v1", func(c chi.Router) {
		c.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		))

		c.Post("/login", login)
		c.Post("/register", register)
	})

	http.ListenAndServe(":8080", c)
}

// model

// type user struct {
// 	email    string
// 	password string
// }

// controller

type UserRequestDTO struct {
	Email    string `json:"email" example:"johndoe@gmail.com"`
	Password string `json:"password" example:"12345678"`
}

// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body UserRequestDTO true "User"
// @Success 200 {string} string "ok"
// @Router /login [post]
// @Security ApiKeyAuth
// @Security BasicAuth
func login(w http.ResponseWriter, r *http.Request) {
	// get user

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("not ok"))
		return
	}

	userReceived := UserRequestDTO{}

	err = json.Unmarshal(body, &userReceived)

	if err != nil {
		w.Write([]byte("not ok"))
		return
	}

	conn, err := database.SQLite{}.Connect()

	if err != nil {
		w.Write([]byte("not ok"))
		return
	}

	ctx := context.Background()

	// verify user exists

	user, err := conn.QueryContext(ctx, "SELECT * FROM users WHERE email = ?", userReceived.Email)

	if err != nil {
		w.Write([]byte("not ok"))
		return
	}

	if !user.Next() {
		w.Write([]byte("not ok"))
		return
	}

	// verify password

	var email string
	var password string

	err = user.Scan(&email, &password)

	if err != nil {
		w.Write([]byte("not ok"))
		return
	}

	if password != userReceived.Password {
		w.Write([]byte("not ok"))
		return
	}

	// create token

	token, err := createToken(email)

	if err != nil {
		w.Write([]byte("not ok"))
		return
	}

	w.Write([]byte(token))

}

// @Summary Register
// @Description Register
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body UserRequestDTO true "User"
// @Success 200 {string} string "ok"
// @Router /register [post]
// @Security ApiKeyAuth
func register(w http.ResponseWriter, r *http.Request) {
	conn, err := database.SQLite{}.Connect()

	if err != nil {
		w.Write([]byte("not ok"))
		return
	}

	ctx := context.Background()

	// get user

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("not ok"))
		return
	}

	userReceived := UserRequestDTO{}

	err = json.Unmarshal(body, &userReceived)

	if err != nil {
		w.Write([]byte("not ok"))
		return
	}

	// verify user exists

	user, err := conn.QueryContext(ctx, "SELECT * FROM users WHERE email = ?", userReceived.Email)

	if err != nil {
		w.Write([]byte("not ok"))
		return
	}

	if user.Next() {
		w.Write([]byte("not ok"))
		return
	}

	// save user
	_, err = conn.ExecContext(ctx, "INSERT INTO users (email, password) VALUES (?, ?)", userReceived.Email, userReceived.Password)

	if err != nil {
		w.Write([]byte("not ok"))
		return
	}

	w.Write([]byte("ok"))
}

var mySigningKey = []byte("my-secret-key")

func createToken(email string) (string, error) {
	// Criar um token
	token := jwt.New(jwt.SigningMethodHS256)

	// Definir as claims (informações) do token
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expira em 24 horas

	// Assinar o token com a chave secreta
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	// Parse o token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Verifique se o método de assinatura é válido
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inválido")
	}

	// Verifique se o token está expirado
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("Token inválido")
	}
}
