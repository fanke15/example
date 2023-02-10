// 通过爬虫&API获取信标链数据
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/imroc/req/v3"
	"github.com/shopspring/decimal"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	si := GetStaked(1)

	siB, _ := json.Marshal(si)
	fmt.Println("staked info:", string(siB))

	// GetTotalBySpider()
}

type StakedInfo struct {
	Epoch             int64           `json:"epoch"`
	ServiceName       string          `json:"service_name"`
	TotalETH          decimal.Decimal `json:"total_eth"`
	ActiveValidators  decimal.Decimal `json:"active_validators"`
	PendingValidators decimal.Decimal `json:"pending_validators"`
	Income            decimal.Decimal `json:"income"`
	PerIncome         decimal.Decimal `json:"per_income"`
	APR               decimal.Decimal `json:"apr"`
}

type (
	epochResp struct {
		Status string
		Data   Epoch
	}
	Epoch struct {
		Attestationscount       int             `json:"attestationscount"`
		Attesterslashingscount  int             `json:"attesterslashingscount"`
		Blockscount             int             `json:"blockscount"`
		Depositscount           int             `json:"depositscount"`
		Epoch                   int             `json:"epoch"`
		Missedblocks            int             `json:"missedblocks"`
		Orphanedblocks          int             `]json:"orphanedblocks"`
		Proposedblocks          int             `json:"proposedblocks"`
		Proposerslashingscount  int             `json:"proposerslashingscount"`
		Scheduledblocks         int             `json:"scheduledblocks"`
		Validatorscount         int             `json:"validatorscount"`
		Voluntaryexitscount     int             `json:"voluntaryexitscount"`
		Finalized               bool            `]json:"finalized"`
		Averagevalidatorbalance decimal.Decimal `json:"averagevalidatorbalance"`
		Eligibleether           decimal.Decimal `json:"eligibleether"`
		Globalparticipationrate decimal.Decimal `json:"globalparticipationrate"`
		Totalvalidatorbalance   decimal.Decimal `json:"totalvalidatorbalance"`
		Votedether              decimal.Decimal `json:"votedether"`
		Slots                   []Block         `json:"slots"`
	}
	Block struct {
		Attestationscount          int    `json:"attestationscount"`
		Attesterslashingscount     int    `json:"attesterslashingscount"`
		Depositscount              int    `json:"depositscount"`
		Epoch                      int    `json:"epoch"`
		Eth1DataDepositcount       int    `json:"eth1data_depositcount"`
		Proposer                   int    `json:"proposer"`
		Proposerslashingscount     int    `json:"proposerslashingscount"`
		Slot                       int    `json:"slot"`
		SyncaggregateParticipation int    `json:"syncaggregate_participation"`
		Voluntaryexitscount        int    `json:"voluntaryexitscount"`
		Blockroot                  string `json:"blockroot"`
		Eth1DataBlockhash          string `json:"eth1data_blockhash"`
		Eth1DataDepositroot        string `json:"eth1data_depositroot"`
		Graffiti                   string `json:"graffiti"`
		GraffitiText               string `json:"graffiti_text"`
		Parentroot                 string `json:"parentroot"`
		Randaoreveal               string `json:"randaoreveal"`
		Signature                  string `json:"signature"`
		Stateroot                  string `json:"stateroot"`
		Status                     string `json:"status"`
		SyncaggregateBits          string `json:"syncaggregate_bits"`
		SyncaggregateSignature     string `json:"syncaggregate_signature"`
	}
)

const (
	basePath = "https://beaconcha.in/api/v1"
	epochUrl = "/epoch/{epochNum}"
)

// 获取质押统计数据
func GetStaked(epochNum int64) []StakedInfo {
	var (
		sis = make([]StakedInfo, 0)
	)

	// 1. 从beaconcha-api获取某个epoch质押数据
	sis = append(sis, GetTotalByApi(epochNum))

	// 2. 爬虫获取每个池子对应的质押数据
	sis = append(sis, GetPoolsBySpider()...)
	return sis
}

// 接口获取总质押数据
func GetTotalByApi(epochNum int64) StakedInfo {
	var (
		reqC    = apiReq()
		er      = epochResp{}
		totalSI = StakedInfo{}
	)
	res, err := reqC.SetPathParam("epochNum", strconv.FormatInt(epochNum, 10)).SetResult(&er).Get(basePath + epochUrl)
	if err != nil {
		log.Println("tagErr1:", err)
		return totalSI
	}
	if !res.IsSuccess() {
		log.Println("tagErr2:", "请求成功，响应异常")
		return totalSI
	}
	// 数据渲染
	totalSI.ActiveValidators = decimal.NewFromInt(int64(er.Data.Validatorscount))
	totalSI.TotalETH = er.Data.Eligibleether.Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(9)))
	totalSI.ServiceName = "ALL"
	totalSI.Epoch = epochNum

	return totalSI
}

// 爬虫获取总质押数据
func GetTotalBySpider() StakedInfo {
	var (
		spiderC = spiderReq()
	)

	spiderC.OnHTML(".pt-3", func(e *colly.HTMLElement) {
		//active := e.ChildTexts(".col-md-4[1] .p-2[0] span")
		e.ForEach(".col-md-4", func(i int, ee *colly.HTMLElement) {
			if i == 1 {
				fmt.Println(ee.ChildTexts(".p-2 h5 span")[0])
				fmt.Println(ee.ChildTexts(".p-2 h5 span")[1])

				fmt.Println(3333, ee.ChildAttr(".p-2 h5 span", "value"))

			}

			fmt.Println("index:", i)
		})
	})
	err := spiderC.Visit("https://beaconcha.in")
	if err != nil {
		log.Println("tagErr4:", err)
	}

	return StakedInfo{}
}

type PoolInfo struct {
	ServiceName string
	Validators  string
	PerIncome   string
}

// 爬虫获取池子质押数据
func GetPoolsBySpider() []StakedInfo {
	var (
		spiderC = spiderReq()
		poolSIs = make([]StakedInfo, 0)
	)

	spiderC.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, ee *colly.HTMLElement) {
			active, _ := decimal.NewFromString(ee.ChildTexts("td")[1])
			per, _ := decimal.NewFromString(getFiguresByStr(ee.ChildTexts("td span")[2]))

			poolSIs = append(poolSIs, StakedInfo{
				Epoch:             0,
				ServiceName:       ee.ChildTexts("td")[0],
				TotalETH:          active.Mul(decimal.NewFromInt(32)),
				ActiveValidators:  active,
				PendingValidators: decimal.Decimal{},
				Income:            decimal.Decimal{},
				PerIncome:         per,
				APR:               decimal.Decimal{},
			})
		})
	})
	err := spiderC.Visit("https://beaconcha.in/pools")
	if err != nil {
		log.Println("tagErr3:", err)
	}
	return poolSIs
}

//---------------------------内部私有方法---------------------------//

const timeout = 60 * time.Second // 超时时间设置

// api请求封装
func apiReq() *req.Request {
	return req.C().SetTimeout(timeout).SetCommonDumpOptions(&req.DumpOptions{
		Output:         os.Stdout,
		RequestHeader:  false, // 是否打印请求头信息
		ResponseBody:   false, // 是否打印响应结果信息
		RequestBody:    false, // 是否打印请求参数信息
		ResponseHeader: false, // 是否打印响应头信息
		Async:          false,
	}).EnableDumpAll().R()
}

const beaconApiDomain = "beaconcha.in"

// spider请求封装
func spiderReq() *colly.Collector {
	return colly.NewCollector(
		colly.AllowedDomains(beaconApiDomain), // 限定爬取数据源
	)
}

// 提取字符串中的浮点数
func getFiguresByStr(str string) string {
	re := regexp.MustCompile("[0-9]\\d*\\.?\\d*")
	return strings.Join(re.FindAllString(str, -1), "")
}
