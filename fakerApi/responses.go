package fakerApi

import (
	"NikitaTest/newClient"
	"bytes"
	"io"
	"net/http"
	"net/url"
	"time"
)

type RespUser struct {
	Status string     `json:"status"`
	Code   int        `json:"code"`
	Total  int        `json:"total"`
	Data   []DataResp `json:"data"`
}

type DataResp struct {
	Id         int    `json:"id"`
	Uuid       string `json:"uuid"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Ip         string `json:"ip"`
	MacAddress string `json:"macAddress"`
	Website    string `json:"website"`
	Image      string `json:"image"`
}

func GetUserFakerApi(quantity string, gender string) []byte {
	baseUrl := "https://fakerapi.it/api/v1/users" //?_quantity=100&_gender=male
	baseParse, err := url.Parse(baseUrl)
	if err != nil {
		panic(err)
	}
	rawQuery := url.Values{}
	rawQuery.Add("_quantity", quantity)
	rawQuery.Add("_gender", gender)
	baseParse.RawQuery = rawQuery.Encode()
	finishUrl := baseParse.String()
	client, err := newClient.NewClient(time.Second * 2)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("GET", finishUrl, bytes.NewBuffer(nil))
	if err != nil {
		panic(err)
	}
	resp, err := client.C.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return respBytes
	//itog := dataResp{}
	//json.Unmarshal(respBytes, &itog)
}
