package module

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/beego/beego/v2/core/config"
	beego "github.com/beego/beego/v2/server/web"
	"io"
	"sort"
	"strings"
)

type ManagerController struct {
	beego.Controller
}

// VerifyToken is used to verify the server address is valid
func (c *ManagerController) VerifyToken() {

	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")
	TOKEN, _ := config.String("TOKEN")

	if validateWeChatRequest(TOKEN, signature, timestamp, nonce) {
		c.Ctx.WriteString(echostr)
	} else {
		c.Ctx.WriteString("Signature verification failed")
	}
}

// validateWeChatRequest is used to validate the signature in the request
func validateWeChatRequest(token, signature, timestamp, nonce string) bool {
	strs := []string{token, timestamp, nonce}
	sort.Strings(strs) // Sort alphabetically
	str := strings.Join(strs, "")

	h := sha1.New()
	io.WriteString(h, str)
	sha1Hash := hex.EncodeToString(h.Sum(nil))

	return sha1Hash == signature
}
