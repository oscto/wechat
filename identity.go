package wechat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// GetUserInfo /cgi-bin/user/info
func (c *Client) GetUserInfo(ctx context.Context, openId string) (*UserInfo, error) {

	userInfo := &UserInfo{}
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s%s", c.ApiUrl, "/cgi-bin/user/info"), nil)
	if err != nil {
		return nil, err
	}

	token, err := c.GetAccessToken(context.Background())
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("access_token", token.AccessToken)
	q.Add("openid", openId)
	req.URL.RawQuery = q.Encode()
	if err := c.Send(req, userInfo); err != nil {
		return nil, err
	}

	return userInfo, err
}
