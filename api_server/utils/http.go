package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"time"

	// "strconv"
	// "strings"

)

type Info struct {
	Code int64                    `json:"code"`
	Data []map[string]interface{} `json:"data"`
}
type DataInfo struct {
	Code int64                  `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}

var tr *http.Transport

// 所有client都共用一个transport
func init() {
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{ // https 请求跳过证书验证
			InsecureSkipVerify: true,
		},
		MaxIdleConns: 100,
	}
}

// params: contentType:"application/json" url 为请求url data 为请求参数
// tips: 使用前端返回参数请求 {"action_client_code":"PC",,"access_token":"MTY3MTU4NTUwMy4zMDU2NTc0OmU1Y2NlZDQzNWU4OGY4NDIxYTBhNTRjYjVjN2Y0YTBkY2M3MDU2MDc=","action_user_code":"USER2020042101",}
func PostRequest(url string, data interface{}, contentType string) (int, string, error) {
	client := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	jsonStr, _ := json.Marshal(data)
	// fmt.Println(bytes.NewBuffer(jsonStr))
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.46")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	// resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	result, _ := io.ReadAll(resp.Body)
	resCode := resp.StatusCode
	// fmt.Println(result)
	return resCode, string(result), nil
}

/*
desc：获取授权信息 data 数据返回示例

	[
	    {
	        "product_name": "MRIA",
	        "auth_status": 1,
	        "sn": "GNNEMM22GXVKWKYS",
	        "end_date": "2025-04-04 15:36:48",
	        "expiration_days": "360",
	        "last_update_date": "2024-04-09 15:37:04"
	    }
	]
*/
func GetAuthInfo(url string, data interface{}, contentType string) (Info, error) {
	var data_info Info
	_, result, err := PostRequest(url, data, "application/json")
	if err != nil {
		return data_info, err
	}
	json.Unmarshal(String2bytes(result), &data_info)
	if data_info.Code == int64(10000) {
		return data_info, nil
	} else {
		return data_info, nil
	}
}

