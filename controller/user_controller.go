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

const (
	MSG_DATABASE_CONNECTION string = "Failed to receive database connection"
	MSG_JSON_BIND           string = "Cannot bind to json"
	MSG_MISSING_ID          string = "Id must be defined"
	MSG_ID_INTEGER          string = "Id must be numeric"
	MSG_CREATE_USER         string = "Cannot create user"
	MSG_LIST_USER           string = "Cannot create user"
	MSG_FIND_BY_ID          string = "Cannot find user"
	MSG_UPDATE_ERROR        string = "Cannot update user"
	MSG_DELETE_ERROR        string = "Cannot delete user"
	MSG_DELETED             string = "User Deleted"
)

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
	db, err := database.GetDatabase()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": MSG_DATABASE_CONNECTION,
		})
		return
	}
	var user model.User
	err = db.First(&user, id).Error
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": MSG_FIND_BY_ID,
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller UserController) Update(c *gin.Context) {
	user_id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": MSG_MISSING_ID,
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
	var user model.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": MSG_JSON_BIND,
		})
		return
	}
	result := db.Where("id = ?", user_id).Updates(model.User{
		Name:  user.Name,
		Login: user.Login,
		Email: user.Email,
	})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Rows affected: 0",
		})
		return
	}
	err = result.Error
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": MSG_UPDATE_ERROR,
		})
		return
	}
	err = db.First(&user, user_id).Error
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": MSG_FIND_BY_ID,
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller UserController) Delete(c *gin.Context) {
	user_id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": MSG_MISSING_ID,
		})
		return
	}
	id, err := strconv.Atoi(user_id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": MSG_ID_INTEGER,
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
	result := db.Delete(&model.User{}, id)
	err = result.Error
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": MSG_FIND_BY_ID,
		})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Rows affected: 0",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": MSG_DELETED,
	})
}
