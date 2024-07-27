package files

import (
	"karma_files_go/database"
	"log"

	_ "github.com/lib/pq"
	gonanoid "github.com/matoous/go-nanoid/v2"
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

func CreateFile(oid string, user_id string, filename string, description string) {
	db, err := database.DBConn()
	if err != nil {
		log.Fatalln(err)
	}

	//Generate a uid
	id, _ := gonanoid.Generate("qwertyuiopasdfghjklzxcvbnm1234567890_-", 10)
	uid := oid + "---" + id

	r, err := db.Exec(`INSERT INTO files (id, user_id, filename, description) VALUES ($1, $2, $3, $4)`, uid, user_id, filename, description)

	if err != nil || r == nil {
		log.Fatalln(err)
	}
}
