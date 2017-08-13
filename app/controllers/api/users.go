package api

import (
	"at-field-control/app/models"
	"fmt"
	"github.com/revel/revel"
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
	user := &models.User{}

	fmt.Print(user)

	if err := c.Params.BindJSON(&user); err != nil {
		return c.HandleBadRequestError(err.Error())
	}

	if err := validator.Validate(&user); err != nil {
		return c.HandleBadRequestError(err.Error())
	}

	old := models.User{}
	models.DB.First(&old, user.ID)

	if old.ID == 0 {
		user.ID = 0
		models.DB.Create(user)
	} else {
		models.DB.Save(user)
	}

	r := Response{user}

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
