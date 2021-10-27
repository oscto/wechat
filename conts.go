package wechat

import (
	"time"
)

const (
	ApiUrl                         = "https://api.weixin.qq.com"
	RequestNewTokenBeforeExpiresIn = time.Duration(60) * time.Second
	OfficialSubscribe              = "subscribe"
	OfficialUnsubscribe            = "unsubscribe"
	QRCodeUrl                      = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s"
)
