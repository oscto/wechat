package wechat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// NewClient 初始化,Token 为后台设置的 Token
func NewClient(appId, appSecret, token string) (*Client, error) {

	return &Client{
		ApiUrl:     ApiUrl,
		AppId:      appId,
		AppSecret:  appSecret,
		Token:      token,
		HttpClient: &http.Client{},
	}, nil
}

func (c *Client) NewRequest(ctx context.Context, method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequestWithContext(ctx, method, url, buf)
}

// GetAccessToken /cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
func (c *Client) GetAccessToken(ctx context.Context) (*AccessTokenResponse, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s%s", c.ApiUrl, "/cgi-bin/token"), nil)
	if err != nil {
		return &AccessTokenResponse{}, err
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	q := url.Values{}
	q.Add("grant_type", "client_credential")
	q.Add("appid", c.AppId)
	q.Add("secret", c.AppSecret)
	req.URL.RawQuery = q.Encode()

	res := &AccessTokenResponse{}
	if err := c.Send(req, res); err != nil {
		return nil, err
	}

	if res.AccessToken != "" {
		c.AccessToken = res
		c.AccessTokenExpiresAt = time.Now().Add(time.Duration(res.ExpiresIn) * time.Second)
	}

	return res, err
}

func (c *Client) SetAccessToken(token string) {

	c.AccessToken = &AccessTokenResponse{
		AccessToken: token,
		ExpiresIn:   7200,
	}
	c.AccessTokenExpiresAt = time.Time{}
}

func (c *Client) SendWithAuth(req *http.Request, v interface{}) error {
	//c.Lock()
	//if c.AccessToken != nil {
	//	if c.AccessTokenExpiresAt.IsZero() && c.AccessTokenExpiresAt.Sub(time.Now()) < RequestNewTokenBeforeExpiresIn {
	//		if _, err := c.GetAccessToken(req.Context()); err != nil {
	//			c.Unlock()
	//			return err
	//		}
	//	}
	//	q := url.Values{}
	//	q.Add("access_token", c.AccessToken.AccessToken)
	//	req.URL.RawQuery = q.Encode()
	//}
	//
	//c.Unlock()

	token, err := c.GetAccessToken(context.Background())
	if err != nil {
		return err
	}
	c.AccessToken.AccessToken = token.AccessToken

	q := url.Values{}
	q.Add("access_token", c.AccessToken.AccessToken)
	req.URL.RawQuery = q.Encode()

	return c.Send(req, v)
}

func (c *Client) Send(req *http.Request, v interface{}) error {

	var err error
	var res *http.Response
	var data []byte

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "zh_CN")
	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}

	res, err = c.HttpClient.Do(req)
	c.log(req, res)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		errResp := &ErrorResponse{Response: res}
		data, err = ioutil.ReadAll(res.Body)
		if err == nil && len(data) > 0 {
			_ = json.Unmarshal(data, errResp)
		}

		return err
	}

	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, err := io.Copy(w, res.Body)
		return err
	}

	return json.NewDecoder(res.Body).Decode(v)
}

func (c *Client) log(r *http.Request, resp *http.Response) {

	var (
		reqDump  string
		respDump []byte
	)

	if r != nil {
		reqDump = fmt.Sprintf("%s %s. Data: %s", r.Method, r.URL.String(), r.Form.Encode())
	}
	if resp != nil {
		respDump, _ = httputil.DumpResponse(resp, true)
	}

	fmt.Println("-------------请求开始----------------")
	fmt.Println("reqDump", reqDump)
	fmt.Println("respDump", respDump)
	fmt.Println("-------------请求结束----------------")

	//if c.Logger != nil {
	//	var (
	//		reqDump  string
	//		respDump []byte
	//	)
	//
	//	if r != nil {
	//		reqDump = fmt.Sprintf("%s %s. Data: %s", r.Method, r.URL.String(), r.Form.Encode())
	//	}
	//	if resp != nil {
	//		respDump, _ = httputil.DumpResponse(resp, true)
	//	}
	//	c.Logger.Write([]byte(fmt.Sprintf("Request: %s\nResponse: %s\n", reqDump, string(respDump))))
	//}
}
