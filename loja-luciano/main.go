package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LulianoM/go-workhard/tree/main/loja-luciano/routes"
)

func main() {
	routes.CarregaRotas()

	porta := ":8000"
	fmt.Println("Conecte-se  na porta" + porta)
	log.Fatal(http.ListenAndServe(porta, nil))
}
