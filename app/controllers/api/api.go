package api

import (
	"github.com/revel/revel"
	"net/http"
)

// 埋め込みによって revel.Controller をラップした ApiV1Controller を定義する
type Api struct {
	*revel.Controller
}

// 正常な際に返す Json 用の構造体(今回は1種類で統一する)
type Response struct {
	Results interface{} `json:"results"`
}

func (c *Api) SetMessage(s string) {
	c.Args["Message"] = s
}

func (c *Api) GetMessage() string {
	message, ok := c.Args["Message"].(string)
	if !ok {
		message = ""
	}
	return message
}

// エラーの際に返す Json 用の構造体
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Bad Request Error を返すやつ
func (c *Api) HandleBadRequestError(s string) revel.Result {
	c.Response.Status = http.StatusBadRequest
	if s == "" {
		s = "Bad Request"
	}
	c.SetMessage(s)
	return c.RenderJSON(nil)
}

// Not Found Error を返すやつ
func (c *Api) HandleNotFoundError(s string) revel.Result {
	c.Response.Status = http.StatusNotFound
	if s == "" {
		s = "Not Found"
	}
	c.SetMessage(s)
	return c.RenderJSON(nil)
}

// Internal Server Error を返すやつ
func (c *Api) HandleInternalServerError(s string) revel.Result {
	c.Response.Status = http.StatusInternalServerError
	if s == "" {
		s = "Internal Server Error"
	}
	c.SetMessage(s)
	return c.RenderJSON(nil)
}

// Precondition Failed を返すやつ(存在しないリソースを更新した場合等)
func (c *Api) HandlePreconditionFailed(s string) revel.Result {
	c.Response.Status = http.StatusPreconditionFailed
	if s == "" {
		s = "Precondition Failed"
	}
	c.SetMessage(s)
	return c.RenderJSON(nil)
}

type ApiResultJSON struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

func (c *Api) RenderJSON(o interface{}) revel.Result {
	if c.Response.Status == 0 {
		c.Response.Status = http.StatusOK
	}

	api_result_json := ApiResultJSON{
		Data:    o,
		Code:    c.Response.Status,
		Message: c.GetMessage(),
	}
	return c.Controller.RenderJSON(api_result_json)
}
