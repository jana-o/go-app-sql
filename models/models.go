package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Photo struct {
	ID  int64  `json:"photoId"`
	Src string `json:"src"`
}

type PhotoCollection struct {
	Photos []Photo `json:"items"`
}

func GetPhotos(db *sql.DB) PhotoCollection {
	sql := "SELECT * FROM photos"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := PhotoCollection{}

	for rows.Next() {
		photo := Photo{}

		err2 := rows.Scan(&photo.ID, &photo.Src)
		if err2 != nil {
			panic(err2)
		}

		result.Photos = append(result.Photos, photo)
	}

	return result
}

//UpdatePhoto updates existing photo in DB
func UpdatePhoto(db *sql.DB, index int, src string) (int64, error) {
	sql := "UPDATE photos SET (src) = (?) WHERE id = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(src, index)
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

//UploadPhoto to DB
func UploadPhoto(db *sql.DB, fsrc string) Photo {

	sql := "INSERT INTO photos (src) VALUES(?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(fsrc)
	if err2 != nil {
		panic(err2)
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	photo := Photo{
		ID:  insertedID,
		Src: fsrc,
	}

	return photo

	// return result.RowsAffected()
}
