package wechat

import "time"

const (
	ApiUrl                         = "https://api.weixin.qq.com"
	EventOfficialSubscribe         = "subscribe"
	EventOfficialUnsubscribe       = "unsubscribe"
	EventOfficialScan              = "SCAN"
	RequestNewTokenBeforeExpiresIn = time.Duration(60) * time.Second
	QRCodeUrl                      = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s"
)
