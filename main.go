package main

import (
	"code/plant-diary/handlers"
	"database/sql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type Photo struct {
	ID  int64  `json:"photoId"`
	Src string `json:"src"`
}

type PhotoCollection struct {
	Photos []Photo `json:"items"`
}

var db *sql.DB

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// initialize DB
	db := initDB("database/plantingdiary.sqlite")
	migrateDB(db)

	// HTTP routes
	e.File("/", "public/index.html")
	e.GET("/photos", handlers.GetPhotos(db))
	e.POST("/photos", handlers.UploadPhoto(db))
	e.DELETE("/photos/:photoId", handlers.DeletePhoto(db))
	e.Static("/uploads", "public/uploads")

	e.Logger.Fatal(e.Start(":8080"))
}

// DB
func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic("Error connecting to database")
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrateDB(db *sql.DB) {
	sql := `
        CREATE TABLE IF NOT EXISTS photos(
                id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                src VARCHAR NOT NULL
		);`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
