package doscg

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/kr/pretty"
	"github.com/labstack/echo"
	"googlemaps.github.io/maps"
)

func FindXYZ(c echo.Context) error {
	numSequence := []string{"X", "Y", "5", "9", "15", "23", "Z"}
	var d2 []string
	var d1 int

	for i := range numSequence {
		if len(numSequence)-1 == i {
			break
		}
		num1, err1 := strconv.Atoi(numSequence[i])
		num2, err2 := strconv.Atoi(numSequence[i+1])
		if err1 == nil && err2 == nil {
			d2 = append(d2, strconv.Itoa(num2-num1))
		} else {
			d2 = append(d2, "")
		}
	}

	for i := range d2 {
		if len(d2)-1 == i {
			break
		}
		num1, err1 := strconv.Atoi(d2[i])
		num2, err2 := strconv.Atoi(d2[i+1])
		if err1 == nil && err2 == nil {
			d1 = num2 - num1
		}
	}

	d, err := strconv.Atoi(d2[2])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error")
	}
	a, err := strconv.Atoi(numSequence[2])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error")
	}

	d -= d1
	Y := a - d
	d -= d1
	X := Y - d

	d, err = strconv.Atoi(d2[4])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error")
	}
	a, err = strconv.Atoi(numSequence[5])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error")
	}
	Z := a + (d + d1)

	return c.JSON(http.StatusOK, map[string]int{
		"X": X,
		"Y": Y,
		"Z": Z,
	})
}

func FindBC(c echo.Context) error {
	A := 21

	return c.JSON(http.StatusOK, map[string]int{
		"B": 23 - A,
		"C": -21 - A,
	})
}

func FindDirection(c echo.Context) error {
	cli, err := maps.NewClient(maps.WithAPIKey("AIzaSyB-mGuzizns7nTZAJ8C3p5cVqHzWlGkWGw"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      "Sydney",
		Destination: "Perth",
	}
	route, _, err := cli.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	pretty.Println(route)
	return c.JSON(200, "ok")
}

type LineMessage struct {
	Destination string `json:"destination"`
	Events      []struct {
		ReplyToken string    `json:"replyToken"`
		Type       string    `json:"type"`
		Mode       string    `json:"mode"`
		TimeStamp  time.Time `json:"timestamp"`
		Source     struct {
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"source"`
		Message struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"events"`
}

func ReceiveLineMessage(c echo.Context) error {
	var input LineMessage
	if err := c.Bind(input); err != nil {
		return err
	}

	return c.JSON(200, input)
}
