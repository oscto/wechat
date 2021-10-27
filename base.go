package wechat

import (
	"encoding/xml"
)

func (c *Client) Validate(signature, timestamp, nonce string) bool {

	return signature == generalSign(c.Token, timestamp, nonce)
}

func (c *Client) ParseMessage(str string) (res *MessageResponse, err error) {

	err = xml.Unmarshal([]byte(str), &res)

	return
}
