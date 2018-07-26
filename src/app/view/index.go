package view

import (
	"net/http"
	"fmt"
	"middleWare"
)

func Index(w http.ResponseWriter, r *http.Request){


	fmt.Fprint(w, "<a href=\"/logout\">logout</a>")
}

func Login(w http.ResponseWriter, r *http.Request){
	middleWare.Cas.Login(w, r)
}

func Logout(w http.ResponseWriter, r *http.Request){
	middleWare.Cas.Logout(w, r)
}

