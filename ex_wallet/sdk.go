package ex_wallet

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type WalletRestfulProxyImpl struct {
	walletRestHost string
	publicKey      string
	privateKey     string
	client         *http.Client
	X_SIGNATURE    string
	X_ACCESS_KEY   string
}

// 构造函数，初始化 `WalletRestfulProxyImpl` 结构体
func NewWalletRestfulProxyImpl(walletRestHost, publicKey, privateKey string) *WalletRestfulProxyImpl {
	if walletRestHost == "" || publicKey == "" || privateKey == "" {
		log.Fatal("没有配置默认钱包的相关配置无效")
	}

	return &WalletRestfulProxyImpl{
		walletRestHost: walletRestHost,
		publicKey:      publicKey,
		privateKey:     privateKey,
		client: &http.Client{
			Timeout: 10 * time.Second, // 设置超时
		},
		X_SIGNATURE:  "X_SIGNATURE",
		X_ACCESS_KEY: "X_ACCESS_KEY",
	}
}

// 发送 GET 请求并返回响应
func (p *WalletRestfulProxyImpl) doGet(url string, params map[string]interface{}) string {
	str := SignToString(params)
	log.Println(str)
	headers := p.buildHeader(str)

	req, err := http.NewRequest("GET", url+str, nil)
	if err != nil {
		log.Fatalf("请求创建失败: %v", err)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatalf("url:%s 请求异常: %s", url+str, string(body))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func (p *WalletRestfulProxyImpl) doDelete(url string, params map[string]interface{}) string {
	log.Println(params)
	str := SignToString(params)
	log.Println(str)
	headers := p.buildHeader(str)

	req, err := http.NewRequest("DELETE", url+str, nil)
	if err != nil {
		log.Fatalf("请求创建失败: %v", err)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatalf("url:%s 请求异常: %s", url+str, string(body))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// 发送 POST 请求并返回响应
func (p *WalletRestfulProxyImpl) doPost(url string, params map[string]interface{}, body interface{}) string {
	str := SignToString(params)
	headers := p.buildHeader(str)

	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", url+str, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("请求创建失败: %v", err)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	resp, err := p.client.Do(req)
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatalf("url:%s 请求异常: %s", url+str, string(body))
	}

	bodyResp, _ := ioutil.ReadAll(resp.Body)
	return string(bodyResp)
}

// 构建 HTTP 请求头
func (p *WalletRestfulProxyImpl) buildHeader(str string) map[string]string {
	if p.publicKey == "" || p.privateKey == "" {
		log.Fatal("不安全，并不合法的请求")
	}

	headers := make(map[string]string)
	signature, err := Sign([]byte(str), p.privateKey)
	if err != nil {
		log.Println(err)
	}
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"
	headers[p.X_SIGNATURE] = signature
	headers[p.X_ACCESS_KEY] = p.publicKey
	log.Println(signature)
	log.Println(p.publicKey)
	return headers
}

func (p *WalletRestfulProxyImpl) GetUnUsedAddress(param AddressParam) (CommonResponse[[]string], error) {
	url := p.walletRestHost + "/address/unused?"
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	body := p.doGet(url, params)

	var response CommonResponse[[]string]
	log.Println(body)
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) AddressVerify(param AddressVerifyParam) (CommonResponse[bool], error) {
	url := p.walletRestHost + "/address/verify?"
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	body := p.doGet(url, params)

	var response CommonResponse[bool]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) GetAllocatedAddress(param AddressParam) (CommonResponse[[]string], error) {
	url := p.walletRestHost + "/address/allocated?"
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	body := p.doGet(url, params)

	var response CommonResponse[[]string]
	log.Printf("aaa %s", body)
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) AddressBalanceList(param AddressBalanceParam) (CommonResponse[[]AddressBalanceVo], error) {
	url := p.walletRestHost + "/address/balance?"
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	body := p.doGet(url, params)

	var response CommonResponse[[]AddressBalanceVo]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) ListCoinConfig() (CommonResponse[[]CoinConfigVo], error) {
	url := p.walletRestHost + "/coin/list?"
	body := p.doGet(url, make(map[string]interface{}))

	var response CommonResponse[[]CoinConfigVo]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) FlowTransactionRecords(param TransactionFlowParam) (CommonResponse[[]WalletTransactionVo], error) {
	url := p.walletRestHost + "/transaction/low?"
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	body := p.doGet(url, params)

	var response CommonResponse[[]WalletTransactionVo]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) PageTransactionRecords(param TransactionPageParam) (CommonResponse[PageInfo[WalletTransactionVo]], error) {
	url := p.walletRestHost + "/transaction/page?"
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	body := p.doGet(url, params)

	var response CommonResponse[PageInfo[WalletTransactionVo]]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) FindByHash(param TransactionHashParam) (CommonResponse[WalletTransactionVo], error) {
	url := p.walletRestHost + "/transaction/record/hash?"
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	body := p.doGet(url, params)

	var response CommonResponse[WalletTransactionVo]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) CreateWithdraw(param WithdrawCreateParam) (CommonResponse[TxEntityVo], error) {
	url := p.walletRestHost + "/withdraw/create?"
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	body := p.doPost(url, params, param)

	var response CommonResponse[TxEntityVo]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) CancelWithdraw(param WithdrawCancelParam) (CommonResponse[any], error) {
	url := p.walletRestHost + "/withdraw/cancel?"
	log.Println(param)
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	log.Println(params)
	body := p.doDelete(url, params)

	var response CommonResponse[any]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) GetWithdrawRecord(param WithdrawGetParam) (CommonResponse[TxEntityVo], error) {
	url := p.walletRestHost + "/withdraw/record?"
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	body := p.doGet(url, params)

	var response CommonResponse[TxEntityVo]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) FlowWithdrawRecords(param WithdrawFlowParam) (CommonResponse[[]WalletTransactionVo], error) {
	url := p.walletRestHost + "/withdraw/low?"
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	body := p.doGet(url, params)

	var response CommonResponse[[]WalletTransactionVo]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) PageWithdrawRecords(param WithdrawPageParam) (CommonResponse[PageInfo[WalletTransactionVo]], error) {
	url := p.walletRestHost + "/withdraw/page?"
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	body := p.doGet(url, params)

	var response CommonResponse[PageInfo[WalletTransactionVo]]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) GetWalletPublicKey() (CommonResponse[string], error) {
	url := p.walletRestHost + "/service/config/wallet/publicKey"
	body := p.doGet(url, make(map[string]interface{}))

	var response CommonResponse[string]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) UpdSidPublicKey(param UpdPublicKeyParam) (CommonResponse[any], error) {
	url := p.walletRestHost + "/service/config/upd/publicKey"
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	body := p.doPost(url, params, nil)

	var response CommonResponse[any]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (p *WalletRestfulProxyImpl) WalletStatistic(param WalletStatisticParam) (CommonResponse[[]WalletStatisticVo], error) {
	url := p.walletRestHost + "wallet/statistic/by-sid?"
	params := BeanToMap(param) // MapToMap 是将结构体转换为 Map 的函数
	body := p.doGet(url, params)

	var response CommonResponse[[]WalletStatisticVo]
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
