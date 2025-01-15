package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CallBaconipsumClient(typeMeat string) (string, error) {
	apiURL := fmt.Sprintf("https://baconipsum.com/api/?type=%s&paras=99&format=text", typeMeat)

	// สร้าง HTTP request
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Call get api bacon failed: ", err)
		return "", err
	}
	defer resp.Body.Close()

	// อ่านข้อมูลจาก response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read all response failed: ", err)
		return "", err
	}

	return string(body), nil
}
