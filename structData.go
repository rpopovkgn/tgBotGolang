package main

import (
	"encoding/xml"
	"fmt"
)

type Response struct {
	XMLName  xml.Name `xml:"output"`
	Text     string   `xml:",chardata"`
	RECEIVER struct {
		Text       string `xml:",chardata"`
		AGTID      string `xml:"AGT_ID,attr"`
		DATEREPORT string `xml:"DATE_REPORT,attr"`
	} `xml:"RECEIVER"`
	Balance struct {
		Text       string `xml:",chardata"`
		AGTBALANCE string `xml:"AGT_BALANCE"`
		OVERDRAFT  string `xml:"OVERDRAFT"`
	} `xml:"balance"`
}

type ResponseJson struct {
	Balance float64 `json:"balance"`
}

type AccessSite struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func WrapError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
