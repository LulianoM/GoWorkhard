package routes

import (
	"log"
	"net/http"

	"github.com/lulianom/go-workhard/tree/main/api-rest-go/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
