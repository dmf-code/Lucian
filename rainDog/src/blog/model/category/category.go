package category

import "github.com/gin-gonic/gin"

type Field struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Num string `json:"num"`
}

type RestfulFunc func(*gin.Context)

func RestfulHandle(f RestfulFunc) {

}
