package test

import (
	"github.com/orangbus/spider/bootstrap"
	"github.com/orangbus/spider/facades"
	"log"
	"testing"
)

func init() {
	bootstrap.Boot()
}

func TestPing(t *testing.T) {
	status := facades.Spider().BaseUrl("https://www.msnii.com/api/json.php").Ping()
	t.Log(status)
}

func TestGetCate(t *testing.T) {
	list, err := facades.Spider().BaseUrl("https://www.hongniuzy2.com/api.php/provide/vod").GetCateList()
	if err != nil {
		t.Logf("请求错误：%s", err.Error())
		return
	}
	t.Logf("%v", list)
}

func TestGetList(t *testing.T) {
	res, err := facades.Spider().BaseUrl("https://www.msnii.com/api/json.php").SetType(1).Debug().GetList(1)
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
	res, err := facades.Spider().BaseUrl("https://www.msnii.com/api/json.php").Search("斗罗", 1)
	if err != nil {
		t.Logf("请求错误：%s", err.Error())
		return
	}
	t.Logf("page:%v", res.Page)
	t.Logf("total:%d", res.Total)
	t.Logf("list:%v", res.List)
}

func TestDetail(t *testing.T) {
	res, err := facades.Spider().BaseUrl("https://www.msnii.com/api/json.php").Debug().Detail("1456")
	if err != nil {
		t.Logf("请求错误：%s", err.Error())
		return
	}
	log.Printf("%v", res)
}

func TestLivePintai(t *testing.T) {
	list, err := facades.Spider().Live().GetPinTai()
	if err != nil {
		t.Log(err)
		return
	}
	for _, v := range list {
		t.Log(v)
	}
}
func TestLiveZhubo(t *testing.T) {
	list, err := facades.Spider().Live().GetZhubo("jsonweishizhibo.txt")
	if err != nil {
		t.Log(err)
		return
	}
	for _, v := range list {
		t.Log(v)
	}
}
