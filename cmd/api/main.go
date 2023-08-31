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

	uSvc := service.NewUserSvc(db)
	aSvc := service.NewAuthSvc()

	h := ihttp.NewHandler(uSvc, aSvc)

	ihttp.GetRoutes(r, h)

	r.Run("0.0.0.0:8080")
}
