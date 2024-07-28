package files

import (
	"log"

	_ "github.com/lib/pq"
	gonanoid "github.com/matoous/go-nanoid/v2"

	"karma_files_go/database"
)

func GetFiles() ([]database.Files, error) {
	db, err := database.DBConn()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	files := []database.Files{}

	rows, err := db.Queryx("SELECT id, user_id, filename, description FROM files")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for rows.Next() {
		var file database.Files
		err := rows.StructScan(&file)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

func CreateFile(user_id string, filename string, description string) string {
	db, err := database.DBConn()
	if err != nil {
		log.Fatalln(err)
	}

	// Generate a uid
	uid, _ := gonanoid.Generate("qwertyuiopasdfghjklzxcvbnm1234567890_-", 10)

	r, err := db.Exec(
		`INSERT INTO files (id, user_id, filename, description) VALUES ($1, $2, $3, $4)`,
		uid,
		user_id,
		filename,
		description,
	)

	if err != nil || r == nil {
		log.Fatalln(err)
		return ""
	}
	return uid
}
