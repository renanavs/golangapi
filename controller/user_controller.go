package controller

import (
	"log"
	"net/http"
	"projectapi/database"
	"projectapi/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

const MSG_DATABASE_CONNECTION = "Failed to receive database connection"
const MSG_JSON_BIND = "Cannot bind to json"
const MSG_MISSING_ID = "Id must be defined"
const MSG_ID_INTEGER = "Id must be numeric"
const MSG_CREATE_USER = "Cannot create user"
const MSG_LIST_USER = "Cannot create user"

func (controller UserController) List(c *gin.Context) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": MSG_DATABASE_CONNECTION,
		})
		return
	}
	var users []model.User
	err = db.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": MSG_LIST_USER,
		})
		log.Println(err.Error())
		return
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func (controller UserController) Create(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": MSG_JSON_BIND,
		})
		return
	}
	db, err := database.GetDatabase()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": MSG_DATABASE_CONNECTION,
		})
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": MSG_CREATE_USER,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "Created",
		"user":   user,
	})
}

func (controller UserController) Read(c *gin.Context) {
	user_id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": MSG_MISSING_ID,
		})
		return
	}
	id, err := strconv.Atoi(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": MSG_ID_INTEGER,
		})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":   id,
		"status": "read",
	})
}

func (controller UserController) Update(c *gin.Context) {
	user_id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": MSG_MISSING_ID,
		})
		return
	}
	id, _ := strconv.Atoi(user_id)
	c.JSON(http.StatusOK, gin.H{
		"user":   id,
		"status": "update",
	})
}

func (controller UserController) Delete(c *gin.Context) {
	user_id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": MSG_MISSING_ID,
		})
		return
	}
	id, _ := strconv.Atoi(user_id)
	c.JSON(http.StatusOK, gin.H{
		"user":   id,
		"status": "delete",
	})
}
