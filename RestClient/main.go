package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	SyndicatedProduct()
}

func SyndicatedProduct() {

	fmt.Println("--------Syndicated Products GET---------")

	url := "https://ecom-stage.in.samsung.com/v4/configurator/syndicated-product"

	var jsonString = []byte(`{"skus": ["SM-G950FZKD"],"store_id": 0,"offset": 0,"count": 0}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response body: ", resp.Body)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

}
