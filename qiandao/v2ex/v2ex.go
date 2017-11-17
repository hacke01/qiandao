package v2ex

import (
	"log"
	"regexp"
	"time"

	"github.com/headzoo/surf"
)

func httpurl(s, cookie string) string {
	bow := surf.NewBrowser()
	bow.AddRequestHeader("Accept", "text/html")
	bow.AddRequestHeader("Accept-Charset", "utf8")
	bow.AddRequestHeader("Cookie", cookie)
	err := bow.Open(s)
	if err != nil {
		log.Println("网站打开失败!重新打开")
		return ""
	}
	return bow.Body()
}
func regexp0(s, v string) []string {
	reg := regexp.MustCompile(s)
	regs := reg.FindAllString(v, -1)

	return regs
}
func V2ex(cookie string) {
	var html, s1 string
	log.Println("正在访问v2ex.")
	html = httpurl("https://www.v2ex.com/mission/daily", cookie)
	c1 := regexp0(`(/mission/daily/redeem\?once=\w+|每日登录奖励已领取)`, html)
	t := len(c1)
	for i := 0; i < t; i++ {
		if c1[i] == `每日登录奖励已领取` {
			log.Println("你已经签到过了拉！")
		} else {
			s1 = c1[i]
			c2 := "https://www.v2ex.com" + s1
			c3 := httpurl(c2, cookie)
			if c3 != "" {
				log.Println("V2EX签到成功!")
				time.Sleep(1 * time.Second)
			} else {
				c3 = httpurl(c3, cookie)
			}
		}
	}

}
