package api

import (
	"at-field-control/app/models"
	"fmt"
	"github.com/revel/revel"
	_ "gopkg.in/validator.v2"
	"gopkg.in/validator.v2"
)

type Users struct {
	Api
}

func (c Users) Index() revel.Result {
	users := []models.User{}
	if err := models.DB.Find(&users).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	type Result struct {
		Users []models.User `json:"users"`
	}
	result := Result{
		Users: users,
	}

	return c.RenderJSON(result)
}

func (c Users) Show(id int) revel.Result {
	user := models.User{}
	if err := models.DB.First(&user, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	type Result struct {
		User models.User `json:"user"`
	}
	result := Result{
		User: user,
	}

	return c.RenderJSON(result)
}

func (c Users) CreateOrUpdate() revel.Result {
	request := struct {
		 User       models.User `json:"user"`
	}{}

	if err := c.Params.BindJSON(&request); err != nil {
		return c.HandleBadRequestError(err.Error())
	}

	if err := validator.Validate(&request.User); err != nil {
		return c.HandleBadRequestError(err.Error())
	}

	fmt.Println("こちら名前")
	fmt.Print(request.User.Name)
	fmt.Println("<-")

	old := models.User{}
	models.DB.First(&old, request.User.ID)

	if old.ID == 0 {
		request.User.ID = 0
		models.DB.Create(&request.User)
	} else {
		models.DB.Save(&request.User)
	}

	r := Response{&request.User}

	return c.RenderJSON(r)
}

func (c Users) Delete(id int) revel.Result {
	user := models.User{}

	if err := models.DB.First(&user, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	if err := models.DB.Delete(&user).Error; err != nil {
		return c.HandleInternalServerError("Record Delete Failure")
	}

	r := Response{"success"}
	return c.RenderJSON(r)
}
