package main 

import "fmt"
import "net/http"
import "log"
import "github.com/gorilla/mux"
import "encoding/json"
import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"



//Conexión a Mongo
func getSession() *mgo.Session {   //devuelve un tipo mgo.Session, de la libreria mgo.Session
	session, err := mgo.Dial("mongodb://localhost")  
	//objeto mgo y método Dial para conectar a la bbdd mongodb en local

	if err != nil {
		panic(err)
	}

	return session //devuelve la sesión de mongodb
}

//Funcion devuelve un objeto Hosting
func responseHosting(w http.ResponseWriter, status int, results Hosting) {
	w.Header().Set("Content-type", "application/json")
	//respuesta http
	w.WriteHeader(200)
	//Escribo estado, 200 ok
	json.NewEncoder(w).Encode(results)
	//devolvemos objeto
}
//Funcion devuelve una colleccion de objetos Hosting
func responseHostings(w http.ResponseWriter, status int, results []Hosting) {
	w.Header().Set("Content-type", "application/json")
	//respuesta http
	w.WriteHeader(200)
	//Escribo estado, 200 ok
	json.NewEncoder(w).Encode(results)
	//devolvemos objeto
}

var collection = getSession().DB("APIrestDB").C("hostings")
//añado a variable collection la invocación a la función que 
//inicia la sesión con la bbdd, indicando la bbdd y la coleccion

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola mundo desde mi servidor web")
}

func HostingAdd(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body) //recibe los datos que llegan por post

	var hosting_data Hosting
	err := decoder.Decode(&hosting_data)

	if (err != nil){
		panic(err)
	}

	defer r.Body.Close()  //Cerrar/limpiar la lectura

	err=collection.Insert(hosting_data) 
	//le decimos a la variable que tiene la sesión que inserte
	//los datos 

	if err !=nil{
		w.WriteHeader(500)
		return
	}


	responseHosting(w, 200, hosting_data)
}

func HostingList(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Listado de Hostings") //imprime en web

	var results []Hosting

	err:= collection.Find(nil).All(&results)

	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println( results)  //imprime en consola
	}

	responseHostings(w, 200, results)
}

func HostingUpdate(w http.ResponseWriter, r *http.Request){
	params:= mux.Vars(r)  //recogemos parametros por url 
	hosting_id := params["id"]

	if !bson.IsObjectIdHex(hosting_id){
		w.WriteHeader(404)
		return
	}

	oid:= bson.ObjectIdHex(hosting_id)

	
	decoder := json.NewDecoder(r.Body)
	//recogemos objeto json que nos llega del Body

	var hosting_data Hosting
	err:= decoder.Decode(&hosting_data)

	if err != nil {
		panic (err)
		w.WriteHeader(500)
		return		
	}
	defer r.Body.Close()

	document:= bson.M{"_id": oid}
	//Para comprobar q documento quiero actualizar
	change:= bson.M{"$set": hosting_data}
	//Guarda un modelo de bson, para hacer set d los datos que les pasamos
	err = collection.Update(document, change)
	//hace el update


	if err != nil {
		w.WriteHeader(404)
		return		
	}
	responseHosting(w, 200, hosting_data)
}