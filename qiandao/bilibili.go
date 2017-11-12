package bilibili

import (
	"log"
	"net/url"

	"github.com/headzoo/surf"
)

func Bilibili(cookie string) {
	data := make(url.Values)
	data["allDays"] = []string{"1000000"}
	data["hadSignDays"] = []string{"1000000"}
	data["isBonusDay"] = []string{"0"}
	data["specialText"] = []string{""}
	data["text"] = []string{""}
	//把post表单发送给目标服务器
	bow := surf.NewBrowser()
	bow.AddRequestHeader("Accept", "text/html")
	bow.AddRequestHeader("Accept-Charset", "utf8")
	bow.AddRequestHeader("Cookie", cookie)
	err := bow.PostForm("https://api.live.bilibili.com/sign/doSign", data)
	if err != nil {
		log.Println("网站打开失败")
	}

	log.Println("**************哔哩哔哩签到完成**************")
}
