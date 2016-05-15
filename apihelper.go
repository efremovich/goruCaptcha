package goruCaptcha

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func sendGetRequest(apikey string, id string) ([]byte, error) {
	url := fmt.Sprintf("http://rucaptcha.com/res.php?key=%s&action=get&id=%s&json=true", apikey, id)

	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	return checkResult(resp)
}

func sendRequest(token string, name string, path string, params url.Values) ([]byte, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	params.Add("body", sEnc)
	url := "http://rucaptcha.com/in.php?json=true"
	client := &http.Client{}
	resp, err := client.PostForm(url, params)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return checkResult(resp)

}

func checkResult(resp *http.Response) ([]byte, error) {
	if resp.StatusCode != 200 {
		con, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("error:%s-%s", resp.Status, con))
	}
	jsonStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(jsonStr, &result)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("The server returned an invalid JSON response. %s-%s",
			resp.Status, resp.Body))
	}
	fmt.Println(string(jsonStr))
	// if result["status"] != 1.00 {
	// 	return nil, errors.New(fmt.Sprintf("gotelebot: Error.ErrorCode: %s-Description%s",
	// 		result["errorCode"], result["description"]))
	// }
	str, errs := json.Marshal(result["request"])
	if errs != nil {
		fmt.Println("Error encoding JSON")
		return nil, errors.New(fmt.Sprintln("gotelebot"))
	}
	return []byte(str), nil
}
