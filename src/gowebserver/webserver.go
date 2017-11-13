package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func helloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)

	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:   ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello chonggege\n") // write to w to client
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func httpRequstTestFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "today is a good day")
}

func main() {
	http.HandleFunc("/", helloName) // //设置访问的路由
	http.HandleFunc("/login", login)
	http.HandleFunc("/get", httpRequstTestFunc)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("listernAndServe", err)
	}
}
