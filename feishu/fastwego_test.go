package defer_study

import (
	"github.com/faabiosr/cachego/file"
	"github.com/fastwego/feishu"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestABC(t *testing.T) {

	// 内部应用 tenant_access_token 管理器
	Atm = &feishu.DefaultAccessTokenManager{
		Id:    FeishuConfig["AppId"],
		Cache: file.New(os.TempDir()),
		GetRefreshRequestFunc: func() *http.Request {
			payload := `{
            "app_id":"` + FeishuConfig["AppId"] + `",
            "app_secret":"` + FeishuConfig["AppSecret"] + `"
        }`
			req, _ := http.NewRequest(http.MethodPost, feishu.ServerUrl+"/open-apis/auth/v3/tenant_access_token/internal/", strings.NewReader(payload))
			return req
		},
	}

	// 创建 飞书 客户端
	FeishuClient = feishu.NewClient()

	// 调用 api 接口
	tenantAccessToken, _ := Atm.GetAccessToken()

	params := url.Values{}
	params.Add("page_size", "10")
	request, _ := http.NewRequest(http.MethodGet, feishu.ServerUrl+"/open-apis/meeting_room/building/list?"+params.Encode(), nil)
	resp, err := FeishuClient.Do(request, tenantAccessToken)
}