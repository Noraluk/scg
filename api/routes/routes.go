package routes

import (
	"scg/api/controllers/doscg"

	"github.com/labstack/echo"
)

func Run() {
	var port string = "8080"
	e := echo.New()

	e.GET("/find-xyz", doscg.FindXYZ)
	e.GET("/find-bc", doscg.FindBC)
	e.POST("/callback", doscg.ReceiveLineMessage)
	e.Logger.Fatal(e.Start(":" + port))
}
