package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	db "Compendium/db"
	_ "Compendium/globals"
)

func AddCampaign() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		image := c.PostForm("img")
		description := c.PostForm("desc")
		success, _ := db.AddCampaignDB(name, image, description)
		if success {
			println("Nouvelle campagne enregistré")
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Erreur dans la création d'une campagne"})
			return
		}
	}
}

func UpdateCampaign() gin.HandlerFunc { //après page de modif
	return func(c *gin.Context) {
		name := c.PostForm("name")
		image := c.PostForm("img")
		description := c.PostForm("desc")
		id := c.PostForm("id")
		id2, _ := strconv.Atoi(id)
		success, _ := db.UpCampaignDB(name, image, description, id2)
		if success {
			println("Campagne mise à jour")
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			println("aie")
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Erreur dans la mise à jour d'une campagne"})
			return
		}
	}
}

func DeleteCampaign() gin.HandlerFunc {
	return func(c *gin.Context) {
		campaignId, _ := strconv.Atoi(c.Param("id"))
		success, _ := db.DeleteCampaignDB(campaignId)
		if success {
			println("Campagne supprimée")
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Erreur dans la suppression"})
			return
		}
	}
}

func AddWorld() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		image := c.PostForm("img")
		description := c.PostForm("desc")
		id := c.PostForm("id")
		id2, _ := strconv.Atoi(id)
		success, _ := db.AddWorldDB(name, image, description, id2)
		if success {
			println("Nouveau élément world enregistré")
			c.Redirect(http.StatusMovedPermanently, "/enter/"+id)
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Erreur dans la création de world"})
			return
		}
	}
}

func UpdateWorld() gin.HandlerFunc { //après page de modif
	return func(c *gin.Context) {
		name := c.PostForm("name")
		image := c.PostForm("img")
		description := c.PostForm("desc")
		id := c.PostForm("id")
		id2, _ := strconv.Atoi(id)
		success, _ := db.UpWorldDB(name, image, description, id2)
		if success {
			println("World mise à jour")
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Erreur dans la mise à jour d'une campagne"})
			return
		}
	}
}

func DeleteWorld() gin.HandlerFunc {
	return func(c *gin.Context) {
		worldId, _ := strconv.Atoi(c.Param("id"))
		success, _ := db.DeleteWorldDB(worldId)
		if success {
			println("World supprimé")
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Erreur dans la suppression"})
			return
		}
	}
}

func AddLore() gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.PostForm("name")
		image := c.PostForm("img")
		description := c.PostForm("desc")
		id := c.PostForm("id")
		id2, _ := strconv.Atoi(id)
		success, _ := db.AddLoreDB(title, image, description, id2)
		if success {
			println("Nouveau élément lore enregistré")
			c.Redirect(http.StatusMovedPermanently, "/enter/"+id)
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Erreur dans la création de world"})
			return
		}
	}
}

func UpdateLore() gin.HandlerFunc { //après page de modif
	return func(c *gin.Context) {
		name := c.PostForm("name")
		image := c.PostForm("img")
		description := c.PostForm("desc")
		id := c.PostForm("id")
		id2, _ := strconv.Atoi(id)
		success, _ := db.UpLoreDB(name, image, description, id2)
		if success {
			println("Lore mise à jour")
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Erreur dans la mise à jour d'une campagne"})
			return
		}
	}
}

func DeleteLore() gin.HandlerFunc {
	return func(c *gin.Context) {
		loreId, _ := strconv.Atoi(c.Param("id"))
		success, _ := db.DeleteLoreDB(loreId)
		if success {
			println("Lore supprimé")
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Erreur dans la suppression"})
			return
		}
	}
}

func AddCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		image := c.PostForm("img")
		bio := c.PostForm("bio")
		id := c.PostForm("id")
		id2, _ := strconv.Atoi(id)
		println("ok")
		success, _ := db.AddCharacterDB(name, image, bio, id2)
		if success {
			println("Nouveau personnage enregistré")
			c.Redirect(http.StatusMovedPermanently, "/enter/"+id)
		} else {
			println("ah")
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Erreur dans la création de world"})
			return
		}
	}
}

func UpdateCharacter() gin.HandlerFunc { //après page de modif
	return func(c *gin.Context) {
		name := c.PostForm("name")
		image := c.PostForm("img")
		description := c.PostForm("desc")
		id := c.PostForm("id")
		id2, _ := strconv.Atoi(id)
		success, _ := db.UpCharDB(name, image, description, id2)
		if success {
			println("Personnage mise à jour")
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Erreur dans la mise à jour d'une campagne"})
			return
		}
	}
}

func DeleteCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		charId, _ := strconv.Atoi(c.Param("id"))
		success, _ := db.DeleteCharacterDB(charId)
		if success {
			println("Personnage supprimé")
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Erreur dans la suppression"})
			return
		}
	}
}
