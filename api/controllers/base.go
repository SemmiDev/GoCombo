package controllers

import (
	"context"
	"flag"
	"fmt"
	"github.com/SemmiDev/go-combo/api/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"    //mysql database driver
	"gorm.io/driver/postgres" //postgres database driver
	"gorm.io/driver/sqlite"   //sqlite database driver
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) Initialize(Driver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Driver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		s.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Driver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Driver)
		}
	}
	if Driver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=Asia/Jakarta", DbHost, DbPort, DbUser, DbName, DbPassword)
		s.DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Driver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Driver)
		}
	}
	if Driver == "sqlite3" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		s.DB, err = gorm.Open(sqlite.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", Driver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", Driver)
		}
		s.DB.Exec("PRAGMA foreign_keys = ON")
	}

	db, _ := s.DB.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(time.Hour)

	err = s.DB.AutoMigrate(&models.User{}, &models.Village{})
	if err != nil {
		panic("failed to migrate")
	} //database migration
	s.Router = mux.NewRouter()
	s.initializeRoutes()
}

func (s *Server) Run(addr string) {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	srv := &http.Server{
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s.Router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
