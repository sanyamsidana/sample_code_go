package main

import (
	"fmt"
	"net/http"

	"sample_project/database"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"welcome")
}

func handleRequest () {
	http.HandleFunc("/",homePage)
	http.ListenAndServe(":10000",nil)
}

func main(){
	db, err := database.GetDbInstance()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Yeah db connected %T",db)

		// var notifications []database.Notification
		var notification database.Notification
		db.First(&notification)
		fmt.Println(notification)
		handleRequest()
	}
}