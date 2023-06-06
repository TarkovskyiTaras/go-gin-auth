package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	_ "github.com/lib/pq"
	"go-gin-auth/pkg/database"
	"go-gin-auth/pkg/hash"
	"go-gin-auth/repository"
	"go-gin-auth/service"
	"go-gin-auth/transport"
	"net/http"
	"time"
)

func generateSalt(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	salt := base64.StdEncoding.EncodeToString(randomBytes)
	return salt, nil
}

func generateHMACSecret(length int) ([]byte, error) {
	secret := make([]byte, length)
	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func main() {
	db, err := database.NewPostgresConnection(database.ConnectionInfo{Host: "localhost", Port: 5432, UserName: "crud-6", DBName: "crud-6-db", SSLMode: "disable", Password: "12345"})
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	usersRepo := repository.NewUsers(db)
	tokensRepo := repository.NewTokens(db)

	//salt, _ := generateSalt(16)
	salt := "LAl4U69N0UnqzzgpRjRRnQ=="
	hasher := hash.NewSHA1Hasher(salt)
	HMACSecret, _ := generateHMACSecret(32)

	usersService := service.NewUsers(usersRepo, tokensRepo, hasher, HMACSecret, 15*time.Minute, 720*time.Hour)

	handler := transport.NewHandler(usersService)

	server := &http.Server{Addr: "localhost:8080", Handler: handler.InitRoutes()}
	server.ListenAndServe()
}
