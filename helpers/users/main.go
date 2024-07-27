package users

import (
	"karma_files_go/database"
	"log"

	_ "github.com/lib/pq"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GetUsers() ([]database.Users, error) {
	db, err := database.DBConn()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	users := []database.Users{}

	rows, err := db.Queryx("SELECT id, userid, password, date, api_token, sites FROM users")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for rows.Next() {
		var user database.Users
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(oid string, userid string, password string, date string, api_token string, sites string) {
	db, err := database.DBConn()
	if err != nil {
		log.Fatalln(err)
	}

	//Generate a uid
	id, _ := gonanoid.Generate("qwertyuiopasdfghjklzxcvbnm1234567890_-", 10)
	uid := oid + "---" + id

	r, err := db.Exec(`INSERT INTO users (id, userid, password, date, api_token, sites) VALUES ($1, $2, $3, $4, $5, $6)`, uid, userid, password, date, api_token, sites)

	if err != nil || r == nil {
		log.Fatalln(err)
	}
}
