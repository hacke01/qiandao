package smzdm

import (
	"log"
	"net/url"

	"github.com/headzoo/surf"
)

func SMZDM(cookie string) {
	bow := surf.NewBrowser()
	bow.AddRequestHeader("cookie", cookie)
	bow.AddRequestHeader("Connection", "keep-alive")
	bow.AddRequestHeader("Host", "zhiyou.smzdm.com")
	bow.AddRequestHeader("Accept", "*/*")
	bow.AddRequestHeader("Accept-Language", "zh-CN,zh;q=0.8,ja;q=0.6")
	bow.AddRequestHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	bow.AddRequestHeader("Referer", "https://www.smzdm.com/")

	data := make(url.Values)
	data["add_point"] = []string{"0"}
	data["checkin_num"] = []string{"10000"}
	data["exp"] = []string{"10000"}
	data["gold"] = []string{"0"}
	data["point"] = []string{"10000"}
	data["prestige"] = []string{"0"}
	data["rank"] = []string{"0"}
	data["slogan"] = []string{"<div class=\"signIn_data\">今日已领<span class=\"red\">15</span>积分，再签到<span class=\"red\">5</span>天可领<span class=\"red\">20</span>积分</div>"}

	err := bow.PostForm("https://zhiyou.smzdm.com/user/checkin/jsonp_checkin?callback=jQuery112406552515165320463_1510470629864&_=1510470629867", data)
	if err != nil {
		panic("网站打开失败！")
	}

	log.Println("*************什么值得买签到完成*************")

}
