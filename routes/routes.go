package routes

import (
	"Compendium/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/login", controllers.LoginGetHandler())   //affichage de la page login
	g.POST("/login", controllers.LoginPostHandler()) //bouton submit pour connexion de la page login

}

func PrivateRoutes(g *gin.RouterGroup) {

	g.GET("/", controllers.IndexGetHandler())      //page principale
	g.GET("/logout4", controllers.Log4())          //déco
	g.POST("/campaign", controllers.AddCampaign()) //créer des campagnes
	g.GET("/delete/:id", controllers.DeleteCampaign())
	g.GET("/update/:id", controllers.Edit())
	g.POST("/saveC", controllers.UpdateCampaign())
	g.GET("/enter/:id", controllers.EnterCampaign())

	//in campaign
	g.POST("/character", controllers.AddCharacter())
	g.GET("/deleteP/:id", controllers.DeleteCharacter())
	g.GET("/updateP/:id", controllers.EditP())
	g.POST("/saveP", controllers.UpdateCharacter())

	g.POST("/world", controllers.AddWorld())
	g.GET("/deleteW/:id", controllers.DeleteWorld())
	g.GET("/updateW/:id", controllers.EditW())
	g.POST("/saveW", controllers.UpdateWorld())

	g.POST("/lore", controllers.AddLore())
	g.GET("/deleteL/:id", controllers.DeleteLore())
	g.GET("/updateL/:id", controllers.EditL())
	g.POST("/saveL", controllers.UpdateLore())
}
