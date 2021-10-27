package wechat

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
)

// GeneralSign 生成加密 signature
func generalSign(params ...string) string {

	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		_, _ = io.WriteString(h, s)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
