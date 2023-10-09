package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetBalanceSpacetel() (string, error) {

	jsonAccess := new(AccessSite)
	jsonAccess.Login = SpacetelLogin
	jsonAccess.Password = SpacetelPass

	jsonConversion, _ := json.Marshal(jsonAccess)

	jsonData := bytes.NewReader(jsonConversion)

	req, _ := http.NewRequest("POST", "https://online.spacetel.ru:9092/v1/client/login", jsonData)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	defer WrapError(err)
	defer resp.Body.Close()

	bodyText, _ := io.ReadAll(resp.Body)

	anserFromSite := &ResponseJson{}

	_ = json.Unmarshal(bodyText, anserFromSite)

	balance := anserFromSite.Balance
	convertBalanceToString := fmt.Sprintf("%8.2f", balance)
	balanceInfo := "Баланс звонков(spacetel): " + convertBalanceToString

	return balanceInfo, nil

}

func GetBalanceQtelecom() (string, error) {
	client := &http.Client{}
	authData := strings.NewReader(AuthorizationQtelecom)
	req, err := http.NewRequest("POST", "https://go.qtelecom.ru/public/http/", authData)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	anserFromSite := &Response{}

	_ = xml.Unmarshal(bodyText, anserFromSite)

	balance := anserFromSite.Balance.AGTBALANCE

	convertBalanceToFloat, _ := strconv.ParseFloat(balance, 64)
	convertBalanceToString := fmt.Sprintf("%8.2f", convertBalanceToFloat)
	balanceInfo := "Баланс смс(qtelecom): " + convertBalanceToString

	return balanceInfo, nil
}
