package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func getuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintln(w, "you are get user %s", uid)
}

func modifyuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintln(w, "you are modify user %s", uid)
}

func deleteuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintln(w, "you are delete user %s", uid)
}

func adduser(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")
	fmt.Fprintln(w, "you are add user %s", uid)
}

func main() {
	mux := routes.New()
	mux.Get("/user/:uid", getuser)
	//	mux.Post("/user/", adduser)
	//	mux.Del("/user/:uid", deleteuser)
	mux.Post("/user/:uid", modifyuser)
	http.Handle("/", mux)

	http.ListenAndServe(":8088", nil)
}
