package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javinc/graham/domain"
	"github.com/javinc/graham/model"
)

// Foo handler
func Foo(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		switch c.Param("id") {
		case "":
			find(c)
		default:
			get(c)
		}
	case http.MethodPost:
		create(c)
	case http.MethodPatch:
		update(c)
	case http.MethodDelete:
		remove(c)
	}
}

func find(c *gin.Context) {
	o, err := domain.FindFoo(c)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, o)
}

func get(c *gin.Context) {
	id := c.Param("id")
	o, err := domain.GetFoo(c, id)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, o)
}

func create(c *gin.Context) {
	p := new(model.Foo)
	c.BindJSON(p)
	o, err := domain.CreateFoo(c, p)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, o)
}

func update(c *gin.Context) {
	p := new(model.Foo)
	c.BindJSON(p)
	p.ID = c.Param("id")
	o, err := domain.UpdateFoo(c, p)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, o)
}

func remove(c *gin.Context) {
	id := c.Param("id")
	o, err := domain.RemoveFoo(c, id)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, o)
}
