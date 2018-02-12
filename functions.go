package main 

import "fmt"
import "net/http"
//import "log"
//import "github.com/gorilla/mux"
//import "encoding/json"
import "gopkg.in/mgo.v2"
//import "gopkg.in/mgo.v2/bson"



//Conexión a Mongo
func getSession() *mgo.Session {   //devuelve un tipo mgo.Session, de la libreria mgo.Session
	session, err := mgo.Dial("mongodb://localhost")  
	//objeto mgo y método Dial para conectar a la bbdd mongodb en local

	if err != nil {
		panic(err)
	}

	return session //devuelve la sesión de mongodb
}

var collection = getSession().DB("APIrestDB").C("hostings")
//añado a variable collection la invocación a la función que 
//inicia la sesión con la bbdd, indicando la bbdd y la coleccion

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola mundo desde mi servidor web")
}