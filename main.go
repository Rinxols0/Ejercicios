package main 

//import "fmt"
import "net/http"
import "log"


func main() {
	router:= NewRouter()

	server:= http.ListenAndServe(":8080", router) //Levanta y hace que escuche el servidor 
	log.Fatal(server)
}