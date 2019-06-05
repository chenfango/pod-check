package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


type Markdown struct {
	Title string `json:"title"`
	Text string `json:"text"`
}

type Message struct {
	Msgtype string `json:"msgtype"`
	Markdown `json:"markdown"`
}
//var user string
//var branch string
//var action string
//func init()  {
//	flag.StringVar(&user, "user", "Wangyijie", "build user")
//	flag.StringVar(&user, "u", "Wangyijie", "build user")
//	flag.StringVar(&branch, "branch", "develop", "build branch")
//	flag.StringVar(&branch, "b", "develop", "build branch")
//	flag.StringVar(&action, "action", "", "action task")
//	flag.StringVar(&action, "a", "", "action task")
//	flag.Parse()
//}

func send(namespace, pod, reason string)  {
	var st string
	st = template(namespace, pod, reason)
	markdown1 := Markdown{
		Title: "测试部署检查",
		Text: st,
	}
	message1 := Message{
		Msgtype:"markdown",
		Markdown: markdown1,
	}
	dingUrl := "https://oapi.dingtalk.com/robot/send?access_token=2e7ba0145b1bf885c2e75ba0e458ff2b204b2c41d641ebe7ee734fcb73ae88f0"
	b, err := json.Marshal(message1)
	if err != nil{
		fmt.Println("message err:", err)
	}
	body := bytes.NewBuffer([]byte(b))
	res, err := http.Post(dingUrl, "application/json;chartset=utf-8",body)
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)

}
func template(namespace, pod, reason string) string {
	st := namespace + " " + pod + " status Exception: \n" +
		"> kubectl get pods -n" + namespace + "\n\n" +
		"> kubectl logs "+ pod + " -n " + namespace + "\n\n" +
		"> " + reason

	return st
}

