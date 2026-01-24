package test

import (
	"log"
	"testing"

	"github.com/orangbus/spider/bootstrap"
	"github.com/orangbus/spider/facades"
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
	res, err := facades.Spider().BaseUrl("https://www.jingpinx.com/api.php/provide/vod", "https://spider.orangbus.cn?url=").Debug().GetList(1)
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
	res, err := facades.Spider().BaseUrl("https://www.msnii.com/api/json.php").SetKeyword("斗罗").Search(1)
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

func TestGetIdsList(t *testing.T) {
	res, err := facades.Spider().BaseUrl("https://www.msnii.com/api/json.php").Debug().GetIdsList("253197,253195,253193")
	if err != nil {
		t.Logf("请求错误：%s", err.Error())
		return
	}
	log.Printf("%v", res)
}

/*
*
获取直播平台
*/
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

/*
*
获取直播平台直播列表
*/
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
