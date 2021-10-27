package wechat

import (
	"encoding/xml"
	"io"
	"net/http"
	"sync"
	"time"
)

type (
	Client struct {
		sync.Mutex
		ApiUrl               string               `json:"api_url"`
		AppId                string               `json:"app_id"`
		AppSecret            string               `json:"app_secret"`
		Token                string               `json:"token"`
		ExpiresIn            int                  `json:"expires_in"`
		AccessToken          *AccessTokenResponse `json:"access_token"`
		AccessTokenExpiresAt time.Time            `json:"access_token_expires_at"`
		HttpClient           *http.Client
		Logger               io.Writer
	}
	UserInfo struct {
		Subscribe      int    `json:"subscribe"`
		Openid         string `json:"openid"`
		Nickname       string `json:"nickname"`
		Sex            int    `json:"sex"`
		Language       string `json:"language"`
		City           string `json:"city"`
		Province       string `json:"province"`
		Country        string `json:"country"`
		Headimgurl     string `json:"headimgurl"`
		SubscribeTime  int    `json:"subscribe_time"`
		Unionid        string `json:"unionid"`
		Remark         string `json:"remark"`
		Groupid        int    `json:"groupid"`
		TagidList      []int  `json:"tagid_list"`
		SubscribeScene string `json:"subscribe_scene"`
		QrScene        int    `json:"qr_scene"`
		QrSceneStr     string `json:"qr_scene_str"`
	}
	AccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	ErrorResponse struct {
		Response        *http.Response        `json:"-"`
		Name            string                `json:"name"`
		DebugID         string                `json:"debug_id"`
		Message         string                `json:"message"`
		InformationLink string                `json:"information_link"`
		Details         []ErrorResponseDetail `json:"details"`
	}
	ErrorResponseDetail struct {
		Field string `json:"field"`
		Issue string `json:"issue"`
	}
	MessageResponse struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string   `xml:"ToUserName"`
		FromUserName string   `xml:"FromUserName"`
		CreateTime   int      `xml:"CreateTime"`
		MsgType      string   `xml:"MsgType"`
		Event        string   `xml:"Event"`
		EventKey     string   `json:"EventKey"`
	}

	QRCodeCreateBody struct {
		ExpireSeconds *int   `json:"expire_seconds"`
		ActionName    string `json:"action_name"`
		ActionInfo    struct {
			Scene struct {
				SceneId  int    `json:"scene_id"`
				SceneStr string `json:"scene_str"`
			} `json:"scene"`
		} `json:"action_info"`
	}
	QRCodeCreateResponse struct {
		Ticket        string `json:"ticket"`
		ExpireSeconds int    `json:"expire_seconds"`
		Url           string `json:"url"`
		QRCodeUrl     string `json:"qr_code_url"`
	}
)
