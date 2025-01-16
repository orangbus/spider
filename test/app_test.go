package test

import (
	spider "github.com/orangbus/spider"
	"log"
	"testing"
)

var client *spider.Spider

func init() {
	client = spider.NewSpider()
}

func TestPing(t *testing.T) {
	status := client.BaseUrl("https://www.msnii.com/api/json.php").Ping()
	t.Log(status)
}

func TestGetCate(t *testing.T) {
	list, err := client.BaseUrl("https://www.hongniuzy2.com/api.php/provide/vod").GetCateList()
	if err != nil {
		t.Logf("请求错误：%s", err.Error())
		return
	}
	t.Logf("%v", list)
}

func TestGetList(t *testing.T) {
	res, err := client.BaseUrl("https://www.msnii.com/api/json.php").GetList(1)
	if err != nil {
		t.Logf("请求错误：%s", err.Error())
		return
	}
	log.Printf("page:%v", res.Page)
	log.Printf("total:%d", res.Total)
	for _, v := range res.List {
		log.Println(v)
	}
}

func TestSearch(t *testing.T) {
	res, err := client.BaseUrl("https://www.msnii.com/api/json.php").Search("斗罗", 1)
	if err != nil {
		t.Logf("请求错误：%s", err.Error())
		return
	}
	t.Logf("page:%v", res.Page)
	t.Logf("total:%d", res.Total)
	t.Logf("list:%v", res.List)
}

func TestDetail(t *testing.T) {
	res, err := client.BaseUrl("https://www.msnii.com/api/json.php").Detail("1")
	if err != nil {
		t.Logf("请求错误：%s", err.Error())
		return
	}
	log.Printf("%v", res.List)
}
