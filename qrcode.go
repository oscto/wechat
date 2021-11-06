package wechat

import (
	"context"
	"fmt"
)

// QRCodeCreate /cgi-bin/qrcode/create
func (c *Client) QRCodeCreate(ctx context.Context, req QRCodeCreateBody) (*QRCodeCreateResponse, error) {

	qrCode := &QRCodeCreateResponse{}
	res, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s%s", c.ApiUrl, "/cgi-bin/qrcode/create"), req)
	if err != nil {
		return nil, err
	}

	if err := c.SendWithAuth(res, qrCode); err != nil {
		return nil, err
	}
	qrCode.QRCodeUrl = fmt.Sprintf(QRCodeUrl, qrCode.Ticket)

	return qrCode, err
}
