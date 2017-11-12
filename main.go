package main

import (
	"github.com/Caeseason/qiandao/qiandao/bilibili"
	"github.com/Caeseason/qiandao/qiandao/smzdm"
	"github.com/Caeseason/qiandao/qiandao/tieba"

	"log"
	"os"
	"regexp"
	"time"
)

func openconf(s string) string { //读取本地文件
	z, err := os.Open(s)
	if err != nil {
		log.Println("没有cookie.conf文件！")
	}
	buf := make([]byte, 10240)
	n, _ := z.Read(buf)
	cookie := string(buf[:n])
	defer z.Close()
	return cookie
}
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
func date(cookie string) string {
	cookie = ""
	time.Sleep(1 * time.Second)
	return cookie
}
func main() {
	var cookie string
	cookie0 := openconf("cookie.conf") //读取本地conf
	c0 := regexp0("\"tieba\"=(.+|)", cookie0)
	c1 := regexp0("\"bilibili\"=(.+|)", cookie0)
	c2 := regexp0("\"smzdm\"=(.+|)", cookie0)
	cookie = regexp1(`(\"tieba\"=|\r)`, c0[0])
	if cookie != "" {
		tieba.Tieba(cookie) //把贴吧cookie传递到GO/qiandao/tieba.go包使用
		cookie = date(cookie)
	}

	cookie = regexp1("(\"bilibili\"=|\r)", c1[0])
	if cookie != "" {
		bilibili.Bilibili(cookie) //把贴吧cookie传递到GO/qiandao/bilibili.go包使用
		cookie = date(cookie)
	}
	cookie = regexp1("(\"smzdm\"=|\r)", c2[0])
	if cookie != "" {

		smzdm.SMZDM(cookie) //把贴吧cookie传递到GO/qiandao/smzdm.go包使用
		cookie = date(cookie)
	}

}
