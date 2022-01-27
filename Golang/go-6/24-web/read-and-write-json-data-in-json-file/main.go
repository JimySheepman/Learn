package main

import (
	"net/http"

	handlers "./handlers"
)

func main() {
	http.HandleFunc("/addnewuser/", handlers.AddNewUserFunc)
	http.HandleFunc("/notsucceded", handlers.NotSucceded)

	http.HandleFunc("/deleted", handlers.DeletedFunc)
	http.HandleFunc("/deleteuser/deleted", handlers.DeleteUserFunc)
	http.HandleFunc("/deleteuser/", handlers.DeleteUserServe)
	http.HandleFunc("/deleteuser/notsuccededdelete", handlers.NotSuccededDelete)

	http.HandleFunc("/", handlers.IndexFunc)

	http.HandleFunc("/showuser/show", handlers.ShowUserFunc)
	http.HandleFunc("/showuser/", handlers.ShowUser)
	http.HandleFunc("/showuser/notsuccededshow/", handlers.NotSuccededShow)

	http.ListenAndServe(":8080", nil)
}
