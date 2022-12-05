package main

import (
	"fmt"
	"imguessr/pkg/db"
	ihttp "imguessr/pkg/http"
	"imguessr/pkg/service"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	db, err := db.NewMongoStore()
	if err != nil {
		fmt.Printf("Can't load db, err=%v", err)
		return
	}

	s := service.NewUserSvc(db)
	h := ihttp.NewHandler(s)

	ihttp.GetRoutes(r, h)

	r.Run("0.0.0.0:8080")
}
