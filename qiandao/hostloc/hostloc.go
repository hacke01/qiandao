package hostloc

import (
	"log"
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

func Hostloc(cookie string) {
	var html string
	log.Println("正在访问hostloc.")
	var url1 = []string{"http://www.hostloc.com/forum.php", "http://www.hostloc.com/space-uid-25650.html", "http://www.hostloc.com/space-uid-7436.html", "http://www.hostloc.com/space-uid-22176.html", "http://www.hostloc.com/space-uid-23376.html", "http://www.hostloc.com/space-uid-132.html", "http://www.hostloc.com/space-uid-26477.html", "http://www.hostloc.com/space-uid-25285.html", "http://www.hostloc.com/space-uid-26532.html", "http://www.hostloc.com/space-uid-25728.html", "http://www.hostloc.com/space-uid-26440.html", "http://www.hostloc.com/space-uid-18756.html", "http://www.hostloc.com/space-uid-12368.html", "http://www.hostloc.com/space-uid-26564.html"}
	t := len(url1)
	for i := 0; i < t; i++ {
		html = httpurl(url1[i], cookie) //利用cookie打开网页
		if html != "" {
			time.Sleep(1 * time.Second)
		} else {

			html = httpurl(url1[i], cookie) //利用cookie打开网页
		}
	}
	log.Println("hostloc赚取金币成功~")
}
