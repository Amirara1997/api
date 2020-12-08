package handler

import (
	"os"
	"html/template"

)

func temp()  {
	templates , _ := template.ParseFiles("home.html")
	templates.Execute(os.Stdout,"Minidorks")
}