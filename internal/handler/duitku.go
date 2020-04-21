package handler

import (
	"bytes"
	"duitku-integration-api/internal/helper"
	"duitku-integration-api/internal/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// DuitKuCallBack handler
// Docs: https://docs.duitku.com/docs-api.html#callback
func DuitKuCallBack(c *gin.Context) {
	payload := model.DuitKuCallBack{}

	err := c.Bind(&payload)
	if err != nil {
		log.Printf("could not parse json from request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	// TODO: Add your business logic
}

// DuitKuRedirect handler
// Docs: https://docs.duitku.com/docs-api.html#callback
func DuitKuRedirect(c *gin.Context) {
	payload := model.DuitKuRedirectURL{}

	err := c.Bind(&payload)
	if err != nil {
		log.Printf("could not parse json from request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	// TODO: Add your business logic
}

// DuitKuRequestTransaction handler
func DuitKuRequestTransaction(c *gin.Context) {
	payload := model.DuitKuRequestTransaction{}

	err := c.Bind(&payload)
	if err != nil {
		log.Printf("could not parse json from request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	merchantCode := viper.GetString("duitku.merchant_code")
	merchantKey := viper.GetString("duitku.merchant_key")
	payload.ReturnURL = viper.GetString("duitku.redirect_url")
	payload.CallbackURL = viper.GetString("duitku.callback_url")
	payload.MerchantCode = merchantCode
	payload.Signature = helper.MakeMD5String(merchantCode + payload.MerchantOrderID + strconv.Itoa(int(payload.PaymentAmount)) + merchantKey)

	byteReqBody, err := json.Marshal(&payload)
	if err != nil {
		log.Printf("could not parse json from request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	url := fmt.Sprintf("%s/webapi/api/merchant/v2/inquiry", viper.GetString("duitku.base_url"))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteReqBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("could not parse json from request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	resbodyByte, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(resbodyByte))

	resBody := model.DuitKuResponseTransaction{}
	err = json.Unmarshal(resbodyByte, &resBody)
	if err != nil {
		log.Printf("could not parse json from request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	// TODO: Add your business logic
}

// DuitKuCheckTransaction handler
func DuitKuCheckTransaction(c *gin.Context) {
	payload := model.DuitKuRequestCheckTransaction{}

	err := c.Bind(&payload)
	if err != nil {
		log.Printf("could not parse json from request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	merchantCode := viper.GetString("duitku.merchant_code")
	merchantKey := viper.GetString("duitku.merchant_key")
	payload.MerchantCode = merchantCode
	payload.Signature = helper.MakeMD5String(merchantCode + payload.MerchantOrderID + merchantKey)

	byteReqBody, err := json.Marshal(&payload)
	if err != nil {
		log.Printf("could not parse json from request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	url := fmt.Sprintf("%s/webapi/api/merchant/transactionStatus", viper.GetString("duitku.base_url"))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteReqBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("could not parse json from request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	resbodyByte, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(resbodyByte))

	resBody := model.DuitKuResponseCheckTransaction{}
	err = json.Unmarshal(resbodyByte, &resBody)
	if err != nil {
		log.Printf("could not parse json from request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	// TODO: Add your business logic
}
