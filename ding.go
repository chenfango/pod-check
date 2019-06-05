package main

import (
	"net/http"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
)
// func Post(url, contentType string, body io.Reader) (resp *Response, err error)
//func PostForm(url string, data url.Values) (resp *Response, err error)
func main() {
	res, err := http.Get("https://rms.api.aixiangdao.com/actuator/health")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	type Text struct {
		content string
	}
	// type Msg struct {
	// 	msgtype string  
	// 	Text Text       "json: text"
	// }
	var text1 Text
	// var msg1 Msg
	text1.content = "我就是我, 是不一样的烟火"
	// msg1.msgtype = "text"
	// msg1.Text = text1
	msg := url.Values{}
	msg.Set("msgtype", "text")
	msg.Add("text", text1)

	m, err := json.Marshal(msg1)
	if err != nil {
		log.Fatal(err)
	}
	// dingData := url.Values{"msgtype": {"text"}}
	fmt.Printf(m)
	dingUrl := "https://oapi.dingtalk.com/robot/send?access_token=2e7ba0145b1bf885c2e75ba0e458ff2b204b2c41d641ebe7ee734fcb73ae88f0"
	dingresp, dingerr := http.PostForm(dingUrl, msg)

	if dingerr != nil {
		log.Fatal(err)
	}

	defer dingresp.Body.Close()
	body, err := ioutil.ReadAll(dingresp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(body))
}


