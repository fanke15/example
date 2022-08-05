package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/imroc/req/v3"
	"github.com/shopspring/decimal"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	data := ReadExcel(
		[]decimal.Decimal{decimal.NewFromFloat(0.94), decimal.NewFromFloat(0.95)},
		[]decimal.Decimal{decimal.NewFromFloat(0.96), decimal.NewFromFloat(10)},
	)
	Write4("block.csv", data)
}

func main1() {
	data, all, current, _ := GetOtherPriceHistoryByCoin(
		[]decimal.Decimal{decimal.NewFromFloat(0.95), decimal.NewFromFloat(0.96)},
		[]decimal.Decimal{decimal.NewFromFloat(0.97), decimal.NewFromFloat(0.98)},
	)

	Write("back_test.csv", data)
	Write2("all.csv", all)
	Write3("current.csv", current)
}

var (
	reqClient = &req.Client{}
)

func init() {
	opt := &req.DumpOptions{
		Output:         os.Stdout,
		RequestHeader:  false,
		ResponseBody:   false,
		RequestBody:    false,
		ResponseHeader: false,
		Async:          false,
	}
	reqClient = req.C().SetTimeout(120 * time.Second).SetCommonDumpOptions(opt).EnableDumpAll()
}

func InitReq() *req.Request {
	return reqClient.R()
}

type coingeckoPriceData struct {
	Stats   [][]decimal.Decimal `json:"stats"`
	Volumes [][]decimal.Decimal `json:"total_volumes"`
}

type PriceData struct {
	Unix  int64
	Price decimal.Decimal
}

type BackTest struct {
	Start PriceData
	End   PriceData
}

type CurrentDay struct {
	Unix int64
	Min  decimal.Decimal
	Max  decimal.Decimal
}

// 从coingecko获取令牌历史价格，交易量数据
func GetOtherPriceHistoryByCoin(min, max []decimal.Decimal) (data []BackTest, price []PriceData, current []CurrentDay, err error) {
	var (
		tempData  coingeckoPriceData
		tempPrice = make([]PriceData, 0)

		nextTempPrice = make([]PriceData, 0)

		currentMap = make(map[int64]CurrentDay)
	)

	resp, err := InitReq().SetResult(&tempData).Get("https://www.coingecko.com/price_charts/13442/eth/90_days.json")
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	if !resp.IsSuccess() {
		fmt.Println("resp err.")
		return
	}
	fmt.Println(111111, len(tempData.Stats))
	// 渲染响应结果
	for _, v := range tempData.Stats {
		price = append(price, PriceData{
			Unix:  v[0].IntPart(),
			Price: v[1],
		})

		// 渲染该天的最大最小值
		cDay := GetDateByNS(v[0].IntPart(), 8*60*60*1000)
		if ct, ok := currentMap[cDay]; ok {
			currentMap[cDay] = CurrentDay{
				Unix: cDay,
				Min:  decimal.Min(v[1], ct.Min),
				Max:  decimal.Max(v[1], ct.Max),
			}
		} else {
			currentMap[cDay] = CurrentDay{
				Unix: cDay,
				Min:  v[1],
				Max:  v[1],
			}
		}
	}
	for _, v := range currentMap {
		current = append(current, v)
	}
	sort.Slice(current, func(i, j int) bool {
		return current[i].Unix < current[j].Unix
	})

	var index = 0
	sort.Slice(price, func(i, j int) bool {
		return price[i].Unix < price[j].Unix
	})

	// 数据过滤&处理分析
	filter := min
	for i := 0; i < len(price); i++ {
		if price[i].Price.GreaterThanOrEqual(decimal.NewFromFloat(0.95)) && price[i].Price.LessThanOrEqual(decimal.NewFromFloat(0.97)) {

			index++

			nextTempPrice = append(nextTempPrice, price[i])
		}

		if isWithin(price[i].Price, filter) {
			tempPrice = append(tempPrice, price[i])
			if EqualAny(filter, min) {
				filter = max
			} else {
				filter = min
			}
		}
	}

	fmt.Println(11111, len(price), index)
	b, _ := json.Marshal(nextTempPrice)
	fmt.Println(string(b))

	// 处理结果
	realLen := len(tempPrice) / 2
	for i := 0; i < realLen; i++ {
		data = append(data, BackTest{
			Start: tempPrice[i*2],
			End:   tempPrice[i*2+1],
		})
	}
	return
}

// 范围限定
func isBTStatus(min, max []decimal.Decimal, price decimal.Decimal) bool {
	status1 := price.LessThan(min[1]) && price.GreaterThanOrEqual(min[0])
	status2 := price.LessThan(max[1]) && price.GreaterThanOrEqual(max[0])
	return status1 || status2
}

func isWithin(price decimal.Decimal, filter []decimal.Decimal) bool {
	return price.LessThan(filter[1]) && price.GreaterThanOrEqual(filter[0])
}

// 判断两个值是否一致
func EqualAny(p1, p2 interface{}) bool {
	b1, _ := json.Marshal(p1)
	b2, _ := json.Marshal(p2)
	return string(b1) == string(b2)
}

func Write(path string, data []BackTest) {
	File, _ := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	defer File.Close()

	//创建写入接口
	WriterCsv := csv.NewWriter(File)
	_ = WriterCsv.Write([]string{"startDate", "startPrice", "endDate", "endPrice"})
	for _, v := range data {
		//写入一条数据，传入数据为切片(追加模式)
		_ = WriterCsv.Write([]string{
			time.Unix(v.Start.Unix/1000, 10).Format("2006-01-02 15:04"),
			v.Start.Price.Round(4).String(),
			time.Unix(v.End.Unix/1000, 10).Format("2006-01-02 15:04"),
			v.End.Price.Round(4).String(),
		})
	}

	WriterCsv.Flush() //刷新，不刷新是无法写入的
	log.Println("数据写入成功...")
}

func Write2(path string, data []PriceData) {
	File, _ := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	defer File.Close()

	//创建写入接口
	WriterCsv := csv.NewWriter(File)
	_ = WriterCsv.Write([]string{"时间", "价格（ETH）"})
	for _, v := range data {
		//写入一条数据，传入数据为切片(追加模式)
		_ = WriterCsv.Write([]string{
			time.Unix(v.Unix/1000, 10).Format("2006-01-02 15:04"),
			v.Price.Round(4).String(),
		})
	}

	WriterCsv.Flush() //刷新，不刷新是无法写入的
	log.Println("数据写入成功...")
}

func Write3(path string, data []CurrentDay) {
	File, _ := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	defer File.Close()

	//创建写入接口
	WriterCsv := csv.NewWriter(File)
	_ = WriterCsv.Write([]string{"时间", "最小价格（ETH）", "最大价格（ETH）"})
	for _, v := range data {
		//写入一条数据，传入数据为切片(追加模式)
		_ = WriterCsv.Write([]string{
			time.Unix(v.Unix, 10).Format("2006-01-02 15:04"),
			v.Min.Round(4).String(),
			v.Max.Round(4).String(),
		})
	}

	WriterCsv.Flush() //刷新，不刷新是无法写入的
	log.Println("数据写入成功...")
}

func Write4(path string, data []BackTest) {
	File, _ := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	defer File.Close()

	//创建写入接口
	WriterCsv := csv.NewWriter(File)
	_ = WriterCsv.Write([]string{"startBlock", "startPrice", "endBlock", "endPrice"})
	for _, v := range data {
		//写入一条数据，传入数据为切片(追加模式)
		_ = WriterCsv.Write([]string{
			strconv.Itoa(int(v.Start.Unix)),
			v.Start.Price.Round(4).String(),
			strconv.Itoa(int(v.End.Unix)),
			v.End.Price.Round(4).String(),
		})
	}

	WriterCsv.Flush() //刷新，不刷新是无法写入的
	log.Println("数据写入成功...")
}

const DaySecond = 24 * 60 * 60

// 获取毫秒时间戳的当天整点时间戳
func GetDateByNS(i int64, o int64) int64 {
	bn := int64(DaySecond * 1000)
	return ((i/bn)*bn - o) / 1000
}

func ReadExcel(min, max []decimal.Decimal) []BackTest {
	var (
		data = make([]BackTest, 0)

		all       = make([]PriceData, 0)
		tempPrice = make([]PriceData, 0)
	)

	f, err := excelize.OpenFile("./3m.xlsx")
	if err != nil {
		fmt.Println(err)
		return data
	}
	rows := f.GetRows("steth3MonthPrice")
	for k, v := range rows {
		if k == 0 {
			continue
		}
		i, _ := strconv.Atoi(v[0])
		all = append(all, PriceData{
			Unix:  int64(i),
			Price: StrToDecimal(v[1]),
		})
	}
	// 分析
	filter := min
	for i := 0; i < len(all); i++ {
		if isWithin(all[i].Price, filter) {
			tempPrice = append(tempPrice, all[i])
			if EqualAny(filter, min) {
				filter = max
			} else {
				filter = min
			}
		}

	}
	// 处理结果
	realLen := len(tempPrice) / 2
	for i := 0; i < realLen; i++ {
		data = append(data, BackTest{
			Start: tempPrice[i*2],
			End:   tempPrice[i*2+1],
		})
	}
	return data
}

func StrToDecimal(str string) decimal.Decimal {
	v, _ := decimal.NewFromString(str)
	return v
}
