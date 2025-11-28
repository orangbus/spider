# An Movie Spider For A Goravel Extend Package

## 快速入门

安装
```bash
go get -u github.com/orangbus/spider@latest
```

注册
打开 `config/app.go`
```go
import "github.com/orangbus/spider"
```
```go
"providers": []foundation.ServiceProvider{
	...
    &spider.ServiceProvider{},
},

```
使用
```go
import "github.com/orangbus/spider/facades"
```
```go
func (r *WebSpider) Ping(ctx http.Context) http.Response {
    status := facades.Spider().BaseUrl("https://xxx.com").Ping()
    return resp.Data(ctx, status)
}
```
## 主要功能
- [x] 视频列表
- [x] 视频详情
- [x] 根据ids获取视频列表
- [x] 生成视频下载文件
- [x] 直播平台
- [x] 直播列表

完整测试请查看`test` 目录