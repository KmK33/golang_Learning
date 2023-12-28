package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w,"method is not supported ",http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"hello!")

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w,"err:%d",err)
		return
	}
	fmt.Fprintf(w,"POST req succesful")
	name := r.FormValue("name")
	fmt.Fprintf(w,"Name:\n",name)
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Server starting at 3000\n")
	if err := http.ListenAndServe(":3000",nil);
	err != nil {
		log.Fatal(err)
	}


}