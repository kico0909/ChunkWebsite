package view

import (
	"net/http"
	"ChunkLib/template"
	"kernel/public"
)

func Index(w http.ResponseWriter, r *http.Request){
	template.RenderHtml(w, "index", nil)
}

func Login(w http.ResponseWriter, r *http.Request){
	public.Cas.Login(w, r)
}

func Logout(w http.ResponseWriter, r *http.Request){
	public.Cas.Logout(w, r, "/")
}

