package controllers

import (
	"strconv"

	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	db "Compendium/db"
	globals "Compendium/globals"
)

func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": "",
			"user":    user,
		})
	}
}

type Display struct {
	User      interface{}
	Campaigns []globals.Campaign
}

func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			campaigns, _ := db.GetCampaignDB()
			if campaigns == nil {
				c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Echec de la récupération des données", "user": user})
			} else {
				display := Display{}
				display.User = user
				display.Campaigns = campaigns
				c.HTML(http.StatusFound, "index.html", display)
			}
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": "Please login first",
			"user":    user,
		})
	}
}

type CpDisplay struct {
	User        interface{}
	Title       string
	Campaign_id int
	Img         string
	World       []globals.World
	Lore        []globals.Lore
	Characters  []globals.Characters
}

func EnterCampaign() gin.HandlerFunc {
	return func(c *gin.Context) {
		campaignId, _ := strconv.Atoi(c.Param("id"))
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			title, image, err := db.GetCampaignTitleImage(campaignId)
			CheckErr(err)
			world, err := db.GetWorldDB(campaignId)
			CheckErr(err)
			characters, err := db.GetCharactersDB(campaignId)
			CheckErr(err)
			lore, err := db.GetLoreDB(campaignId)
			CheckErr(err)
			if title == "" {
				c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Echec de la récupération des données", "user": user})
			} else {
				CpDisplay := CpDisplay{}
				CpDisplay.Title = title
				CpDisplay.Img = image
				CpDisplay.Campaign_id = campaignId
				CpDisplay.User = user
				CpDisplay.World = world
				CpDisplay.Characters = characters
				CpDisplay.Lore = lore
				c.HTML(http.StatusFound, "campaign.html", CpDisplay)
			}
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"content": "Please login first",
				"user":    user,
			})
		}
	}
}

type Update struct {
	Id int
	Origin int
	Campaign   []globals.Campaign
	World       []globals.World
	Lore        []globals.Lore
	Characters  []globals.Characters
}

func Edit() gin.HandlerFunc { //avant page de modif
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			id, _ := strconv.Atoi(c.Param("id"))
			campaign, err := db.GetCampaignIdDB(id)
			CheckErr(err)
			Update := Update{}
			Update.Id = id
			Update.Origin = 1
			Update.Campaign = campaign
			c.HTML(http.StatusFound, "update.html", Update)
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"content": "Please login first",
				"user":    user,
			})
		}
	}
}

func EditP() gin.HandlerFunc { //avant page de modif
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			id, _ := strconv.Atoi(c.Param("id"))
			characters, err := db.GetCharactersDB(id)
			CheckErr(err)
			Update := Update{}
			Update.Id = id
			Update.Origin = 4
			Update.Characters = characters
			c.HTML(http.StatusFound, "update.html", Update)
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"content": "Please login first",
				"user":    user,
			})
		}
	}
}

func EditW() gin.HandlerFunc { //avant page de modif
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			id, _ := strconv.Atoi(c.Param("id"))
			world, err := db.GetWorldDB(id)
			CheckErr(err)
			Update := Update{}
			Update.Id = id
			Update.Origin = 2
			Update.World = world
			c.HTML(http.StatusFound, "update.html", Update)
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"content": "Please login first",
				"user":    user,
			})
		}
	}
}

func EditL() gin.HandlerFunc { //avant page de modif
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			id, _ := strconv.Atoi(c.Param("id"))
			lore, err := db.GetLoreDB(id)
			CheckErr(err)
			Update := Update{}
			Update.Id = id
			Update.Origin = 3
			Update.Lore = lore
			c.HTML(http.StatusFound, "update.html", Update)
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"content": "Please login first",
				"user":    user,
			})
		}
	}
}
