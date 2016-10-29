package wechat

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"nicosoft.org/wechat/utils"
	"sort"
	"strings"
)

func CheckData(c *gin.Context) ([]string, string, string) {

	signature := c.Param("signature")
	timestamp := c.Param("timestamp")
	nonce := c.Param("nonce")
	echoStr := c.Param("echostr")
	token := new(utils.Cache).Get(utils.WECHAT_TOKEN)

	return sort.StringSlice{token, timestamp, nonce}, signature, echoStr
}

func Check(data []string, signature string) bool {

	if len(data) == 0 || signature == "" {
		return false
	}

	sort.Strings(data)
	str := strings.Join(data, "")
	h := sha1.New()
	h.Write([]byte(str))

	return fmt.Sprintf("%x", h.Sum(nil)) == signature

}
