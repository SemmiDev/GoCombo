package main

import (
	"GoOrder/utils/mysql"
	configs "GoOrder/configurations"
	"context"
	"database/sql"
	"github.com/gorilla/mux"
	repoMysql "GoOrder/repositories/mysql"
	usecases "GoOrder/services"
	apis "GoOrder/controllers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server represents server
type Server struct {
	Reader      *sql.DB
	Writer      *sql.DB
	Port        string
	ServerReady chan bool
}

func main() {
	reader, writer := configureMySQL()
	serverReady := make(chan bool)
	server := Server{
		Reader:      reader,
		Writer:      writer,
		Port:        configs.Server.Port,
		ServerReady: serverReady,
	}
	server.Start()
}

func configureMySQL() (*sql.DB, *sql.DB) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	readerConfig := mysql.Option{
		Host:     os.Getenv("READER_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Database: os.Getenv("MYSQL_DATABASE_NAME"),
		User:     os.Getenv("READER_USER"),
		Password: os.Getenv("READER_PASSWORD"),
	}

	writerConfig := mysql.Option{
		Host:     os.Getenv("WRITER_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Database: os.Getenv("MYSQL_DATABASE_NAME"),
		User:     os.Getenv("WRITER_USER"),
		Password: os.Getenv("WRITER_PASSWORD"),
	}

	reader, writer, err := mysql.SetupDatabase(readerConfig, writerConfig)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect mysql", err)
	}

	log.Println("MySQL connection is successfully established!")

	return reader, writer
}

// Start will start server
func (s *Server) Start() {

	port := configs.Server.Port
	if port == "" {
		port = "8000"
	}

	r := new(mux.Router)

	registerRepo := repoMysql.NewRegistrasiRepo(s.Reader, s.Writer)
	kabupatenRepo := repoMysql.NewKabupatenKotaRepo(s.Reader, s.Writer)
	registerService := usecases.NewRegistrasiService(registerRepo)

	kabupatenService := usecases.NewKabupatenKotaService(kabupatenRepo)
	apis.NewRegistrasiController(r, registerService)
	apis.NewKabupatenKotaController(r, kabupatenService)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Printf("Starting server on port %s!", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Shutting Down Server...")
			log.Fatal(err.Error())
		}
	}()

	if s.ServerReady != nil {
		s.ServerReady <- true
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("failed to gracefully shutdown the server: %s", err)
	}
}