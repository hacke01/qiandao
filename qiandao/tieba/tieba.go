package tieba

import (
	"log"
	"regexp"

	"github.com/headzoo/surf"
)

const (
	urlopen  = "https://wapp.baidu.com/mo/q----,sz@320_240-1-3---2/m?tn=bdFBW&tab=favorite"
	urlopen1 = "http://wapp.baidu.com/mo/q---7E9A5ED7F9B6BABD83D52D8A7E152590%3AFG%3D1-sz%40320_240%2C-1-3-0--2--wapp_1510041009828_519/sign?tbs=e77351fe5c0b710a1510041034&fid=94555&"
)

func regexp0(s, v string) []string { //正则寻找字符串
	reg := regexp.MustCompile(s)
	regs := reg.FindAllString(v, -1)
	return regs
}
func regexp1(s, v string) string { //正则删除字符
	reg := regexp.MustCompile(s)
	regs := reg.ReplaceAllString(v, "")
	return regs
}
func httpurl(s, cookie string) string {
	bow := surf.NewBrowser()
	bow.AddRequestHeader("Accept", "text/html")
	bow.AddRequestHeader("Accept-Charset", "utf8")
	bow.AddRequestHeader("Cookie", cookie)
	err := bow.Open(s)
	if err != nil {
		log.Println("网站打开失败！")
	}
	return bow.Body()
}

func Tieba(cookie string) {
	log.Println("**************正在签到百度贴吧**************")
	html := httpurl(urlopen, cookie) //利用cookie打开网页
	url1 := regexp0(`\.<a href="(?P<tmp>.+?)m\?kw=(?P<tmp>.+?)">([\\u4e00-\\u9fa5]+|[^\\x00-\\xff]+|\\w+)</a>`, html)
	t := len(url1)
	for i := 0; i < t; i++ {
		url2 := regexp1(`(\.<a href=\"//|\">|([\\u4e00-\\u9fa5]+|[^\\x00-\\xff]+|\\w+)</a>)`, url1[i]) //删除多余的元素，留下链接
		name := regexp1(`(\.<a href=\"//.+\">|</a>)`, url1[i])                                         //删除多余的元素，留下贴吧名字
		url3 := regexp0(`kw=.*`, url2)
		url4 := urlopen1 + url3[0]
		_ = httpurl(url4, cookie)
		log.Printf("%v-OK", name)

	}
	log.Printf("**************百度贴吧签到完成**************")

}
