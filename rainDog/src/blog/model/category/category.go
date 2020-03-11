package category

import (
	"github.com/gin-gonic/gin"
)

type PostField struct {
	Name string `json:"name"`
}

type PutField struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type GetField struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Num string `json:"num"`
}

type RestfulFunc func(*gin.Context)

func RestfulHandle(f RestfulFunc) {

}
