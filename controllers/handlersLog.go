package controllers

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/google/uuid"

	"net/http"

	"github.com/gin-gonic/gin"

	db "Compendium/db"
	globals "Compendium/globals"
	mid "Compendium/middleware"
	utils "Compendium/utils"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserByUsernameToAdd(str string) bool {
	Username := str
	user, err := db.GetUserByUsername(Username)
	CheckErr(err)
	// if the name is blank we can assume nothing is found
	if user.Username == "" {
		return false
	} else {
		return true
	}
}

func CheckPasswordByUsername(str string, pwd string) bool {
	Username := str
	user, err := db.GetUserByUsername(Username)
	CheckErr(err)
	if user.Password == "" {
		return false
	} else {
		return mid.CheckPasswordHash(pwd, user.Password)
	}
}

func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//verif de la session, si il n'est pas vide c'est qu'une session est encore en cours
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{"content": "Please logout first"})
			return
		}

		//récupération des données du form
		var json globals.User
		json.Username = c.PostForm("username")
		json.Password = c.PostForm("password")
		if utils.EmptyUserPass(json.Username, json.Password) {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
			return
		}

		//on regarde si ses données existent déjà en bd
		if GetUserByUsernameToAdd(json.Username) {
			if CheckPasswordByUsername(json.Username, json.Password) {
				session.Set(globals.Userkey, json.Username)
				if err := session.Save(); err != nil {
					c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
					return
				}
				c.Redirect(http.StatusMovedPermanently, "/")
			} else {
				c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect username or password"})
				return
			}
		} else { //sinon on crée un nouveau user
			json.Id = uuid.New().String()
			json.Password = mid.Hpwd(json.Password)
			success, _ := db.AddUser(json)
			if success {
				session.Set(globals.Userkey, json.Username) //on crée sa session
				if err := session.Save(); err != nil {
					c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
					return
				}
				c.Redirect(http.StatusMovedPermanently, "/")
			} else {
				c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to register new user"})
				return
			}
		}
	}
}

func Log4() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		log.Println("logging out user:", user)
		if user == nil {
			log.Println("Invalid session token")
			return
		}
		session.Delete(globals.Userkey)
		if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
			return
		}

		c.Redirect(http.StatusFound, "/login")
	}
}
