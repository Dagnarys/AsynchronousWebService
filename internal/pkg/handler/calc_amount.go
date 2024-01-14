package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)
import "github.com/gin-gonic/gin"

type CalcRequest struct {
	Amount    int   `json:"premium_amount"`
	AccessKey int64 `json:"access_key"`
}

type Request struct {
	InsuranceID int64 `json:"insurance_id"`
	Drivers     int   `json:"drivers"`
}

func (h *Handler) issueCalc(c *gin.Context) {
	var input Request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("handler.issueCalc:", input)

	c.Status(http.StatusOK)

	go func() {
		time.Sleep(4 * time.Second)
		sendCalcRequest(input)
	}()
}

func sendCalcRequest(request Request) {

	var premium_amount = 1000 + rand.Intn(10000)

	answer := CalcRequest{
		AccessKey: 123,
		Amount:    premium_amount,
	}

	client := &http.Client{}
	fmt.Print("amount", premium_amount)
	jsonAnswer, _ := json.Marshal(answer)
	bodyReader := bytes.NewReader(jsonAnswer)

	requestURL := fmt.Sprintf("http://127.0.0.1:8000/api/insurances/%d/calc_amount/", request.InsuranceID)

	req, _ := http.NewRequest(http.MethodPost, requestURL, bodyReader)

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("PUT Request Status:3", response.Status)
}
