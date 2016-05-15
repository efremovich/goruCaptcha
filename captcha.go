package goruCaptcha

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

//RuCaptcha start settings struct
type RuCaptcha struct {
	apiKey   string
	phrase   string
	numeric  string
	language string
	answer   chan (string)
	id       string
	stopreq  bool
}

//InitruCaptcha is the function to create goruCaptcha instance.
func InitruCaptcha(apiKey string) *RuCaptcha {

	parseCaptcha := new(RuCaptcha)
	parseCaptcha.apiKey = apiKey
	parseCaptcha.phrase = "0"
	parseCaptcha.language = "0"
	parseCaptcha.language = "0"

	return parseCaptcha
}

//Parse типа текст
func (parseCaptcha *RuCaptcha) Parse(name string, path string) {

	params := url.Values{}
	params.Add("key", parseCaptcha.apiKey)
	params.Add("phrase", parseCaptcha.phrase)
	params.Add("numeric", parseCaptcha.numeric)
	params.Add("language", parseCaptcha.language)
	params.Add("method", "base64")
	answers, _ := sendRequest(parseCaptcha.apiKey, name, path, params)

	fmt.Printf("%#v", strings.Trim(string(answers), "\""))
	parseCaptcha.id = strings.Trim(string(answers), "\"")
}

//StartCheckingStatus comment
func (parseCaptcha *RuCaptcha) StartCheckingStatus() error {
	parseCaptcha.stopreq = false
	for {
		if parseCaptcha.stopreq == true {
			return nil
		}

		newUpdates, err := sendGetRequest(parseCaptcha.apiKey, parseCaptcha.id)
		if err != nil {
			return err
		}
		parseCaptcha.processNewUpdate(newUpdates)
	}
}

func (parseCaptcha *RuCaptcha) processNewUpdate(updates []byte) {

	parseCaptcha.answer <- string(updates)

}

func (parseCaptcha *RuCaptcha) ProcessGetAnswer() {
	newMsgChan := parseCaptcha.answer
	time.Sleep(time.Second * 1)
	for {
		m := <-newMsgChan // Get new messaage, when new message arrive.
		fmt.Printf("Get Message:%#v \n", m)
		if m != "" { // Check message is text message.

		}
	}
}
