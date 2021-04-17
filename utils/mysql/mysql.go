package mysql

import (
	"database/sql"
	// _ mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type Option struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func SetupDatabase(readerConfig, writerConfig Option) (reader, writer *sql.DB, err error) {
	//err = godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	reader, err = createConnection(readerConfig)
	if err != nil {
		return nil, nil, err
	}

	writer, err = createConnection(writerConfig)
	if err != nil {
		return nil, nil, err
	}

	return reader, writer, nil
}

func createConnection(config Option) (db *sql.DB, err error) {
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}

	if config.Port == "" {
		config.Port = "3306"
	}

	auth := config.User + ":" + config.Password
	uri := "tcp(" + config.Host + ":" + config.Port + ")"
	dsn := auth + "@" + uri + "/" + config.Database

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(5)

	err = db.Ping()

	return
}
