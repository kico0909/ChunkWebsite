package view

import (
	"net/http"
	"middleWare"
	"utils/ChunkLib/template"
)

func Index(w http.ResponseWriter, r *http.Request){

	template.RenderHtml(w, "index", nil)

}

func Login(w http.ResponseWriter, r *http.Request){
	middleWare.Cas.Login(w, r)
}

func Logout(w http.ResponseWriter, r *http.Request){
	middleWare.Cas.Logout(w, r)
}

