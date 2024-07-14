package main

import(
	"karma_files_go/routes"
)

func main(){
	app := routes.SetupRoutes()
	app.Listen(":9090")
}