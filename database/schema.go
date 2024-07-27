package database

type Files struct {
	Id          string `json:"id"`
	User_id     string `json:"user_id"`
	Filename    string `json:"filename"`
	Description string `json:"description"`
}

type Users struct {
	Id        string `json:"id"`
	Userid    string `json:"userid"`
	Password  string `json:"password"`
	Date      string `json:"date"`
	Api_token string `json:"api_token"`
	Sites     string `json:"sites"`
}
