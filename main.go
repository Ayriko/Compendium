package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	//"html/template"
	//"strings"

	db "Compendium/db"
	globals "Compendium/globals"
	middleware "Compendium/middleware"
	routes "Compendium/routes"
)

func main() {
	//création de la db / initialisation
	err := db.ConnectDatabase()
	CheckErr(err)
	db.CreateTable(db.DB)
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	store := cookie.NewStore(globals.Secret)
	store.Options(sessions.Options{MaxAge:  60*60*24})
	router.Use(sessions.Sessions("session", store))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired) //permet de vérifier si un utilisateur est connecté avant de servir les routes privées
	routes.PrivateRoutes(private)

	router.Run(":8080") //8080 port par défaut
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
