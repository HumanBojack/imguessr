package main

import (
	"fmt"
	"imguessr/pkg/db"
	ihttp "imguessr/pkg/http"
	"imguessr/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World!")
}

func main() {
	r := gin.Default()

	db, err := db.NewMongoStore()
	if err != nil {
		fmt.Printf("Can't load db, err=%v", err)
		return
	}

	s := service.NewUserSvc(db)
	h := ihttp.NewUserHandler(s)

	ihttp.GetRoutes(r, h)

	r.Run("0.0.0.0:8080")
}
