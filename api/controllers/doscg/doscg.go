package doscg

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
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

type LineMessage struct {
	Destination string `json:"destination"`
	Events      []struct {
		ReplyToken string `json:"replyToken"`
		Type       string `json:"type"`
		Mode       string `json:"mode"`
		TimeStamp  int64  `json:"timestamp"`
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

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type ReplyMessageRequest struct {
	ReplyToken string    `json:"replyToken"`
	Messages   []Message `json:"messages"`
}

type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func ReceiveLineMessage(c echo.Context) error {
	var input LineMessage
	if err := c.Bind(&input); err != nil {
		return err
	}

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Add("client_id", "1654025042")
	data.Add("client_secret", "f7e89d5ec699573e0fa5bfd68291ac72")

	req, err := http.NewRequest("POST", "https://api.line.me/v2/oauth/accessToken", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var token Token
	if err := json.Unmarshal(body, &token); err != nil {
		return err
	}

	for _, event := range input.Events {
		msgTime := time.Unix(0, event.TimeStamp*int64(time.Millisecond))
		replyReq := ReplyMessageRequest{
			ReplyToken: event.ReplyToken,
			Messages: []Message{
				{
					Type: "text",
					Text: "hello",
				},
			},
		}

		if time.Now().After(msgTime.Add(10 * time.Second)) {
			log.Println("sdas")
			return errors.New("line bot can't answer in 10 sec")
		}

		b, err := json.Marshal(replyReq)
		if err != nil {
			return err
		}

		req, err := http.NewRequest("POST", "https://api.line.me/v2/bot/message/reply", bytes.NewBuffer(b))
		if err != nil {
			return err
		}
		req.Header.Add("Authorization", "Bearer "+token.AccessToken)
		req.Header.Add("Content-Type", "application/json")

		res, err = client.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "received message",
	})
}
