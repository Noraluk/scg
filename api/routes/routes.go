package routes

import (
	"os"
	"scg/api/controllers/doscg"

	"github.com/labstack/echo"
)

func Run() {
	e := echo.New()

	e.GET("/find-xyz", doscg.FindXYZ)
	e.GET("/find-bc", doscg.FindBC)
	e.POST("/callback", doscg.ReceiveLineMessage)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
