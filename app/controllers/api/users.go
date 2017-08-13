package api

import (
	"github.com/revel/revel"
)

type Users struct {
	Api
}

func (c Users) Index() revel.Result {
	r := Response{"index"}
	return c.RenderJSON(r)
}

func (c Users) Show(id int) revel.Result {
	r := Response{"show"}
	return c.RenderJSON(r)
}

func (c Users) Create() revel.Result {
	r := Response{"create"}
	return c.RenderJSON(r)
}

func (c Users) Delete(id int) revel.Result {
	r := Response{"delete"}
	return c.RenderJSON(r)
}
