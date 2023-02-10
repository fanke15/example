package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sort"

	"encoding/csv"
	"github.com/shopspring/decimal"
	"log"
	"os"
	"time"
)

func main() {

	data := getStEthData().DataList

	fmt.Println(1111, len(data))

	// 排序存储
	sort.Slice(data, func(i, j int) bool {
		return data[i].CurrentDay < data[j].CurrentDay
	})
	st := make([]stEth, 0)

	for k, v := range data {
		t := stEth{
			CurrentDay:       v.CurrentDay,
			Amount:           v.StakedEther,
			ChangeAmountAll:  v.StakedEther.Sub(data[0].StakedEther),
			Holders:          v.Holders,
			ChangeHoldersAll: v.Holders.Sub(data[0].Holders),
		}
		if k > 0 {
			t.ChangeAmountCurrentDay = v.StakedEther.Sub(data[k-1].StakedEther)
			t.ChangeHoldersCurrentDay = v.Holders.Sub(data[k-1].Holders)
		}

		st = append(st, t)
	}

	writeStEthData(st)

}

type MyData struct {
	ServiceStaked
}

type ServiceStaked struct {
	ID        int                 `gorm:"column:id" json:"id"`
	UUID      string              `gorm:"column:uuid;primaryKey" json:"uuid"`
	Name      string              `gorm:"column:name" json:"name"`
	TokenName string              `gorm:"column:token_name" json:"token_name"`
	Type      string              `gorm:"column:type" json:"type"`
	DataList  []ServiceStakedData `gorm:"foreignkey:Name;references:Name" json:"data_list"`
}

type ServiceStakedData struct {
	ID   int    `gorm:"column:id" json:"id"`
	UUID string `gorm:"column:uuid;primaryKey" json:"uuid"`
	NameEr
	StakedEther       decimal.Decimal `gorm:"column:staked_ether" json:"staked_ether"`
	Validators        decimal.Decimal `gorm:"column:validators" json:"validators"`
	ActiveValidators  decimal.Decimal `gorm:"column:active_validators" json:"active_validators"`   // 激活状态的验证者数量
	PendingValidators decimal.Decimal `gorm:"column:pending_validators" json:"pending_validators"` // 等待状态的验证者数量
	Income            decimal.Decimal `gorm:"column:income" json:"income"`                         // 总收入
	PerIncome         decimal.Decimal `gorm:"column:per_income" json:"per_column"`                 // ETH日收益
	Apr               decimal.Decimal `gorm:"column:apr" json:"apr"`                               // 年化收益
	CM                `gorm:"embedded"`
}

type NameEr struct {
	Name string `gorm:"column:name" json:"name"`
}
type CM struct {
	Apr     decimal.Decimal `gorm:"column:apr" json:"apr"`         // 年化收益
	Volume1 decimal.Decimal `gorm:"column:volume1" json:"volume1"` // 24小时交易量
	Holders decimal.Decimal `gorm:"column:holders" json:"holders"` // 24小时交易量
	CDer    `gorm:"embedded"`
}

type CDer struct {
	CurrentDay int64 `gorm:"column:current_day" json:"current_day"` // 当天时间戳
}

func (ps *ServiceStaked) TableName() string {
	return "postat_service_staked"
}

func (ps *ServiceStakedData) TableName() string {
	return "postat_service_staked_data"
}

// 连接数据库获取所有数据
func getStEthData() MyData {
	temp := MyData{}
	db, err := gorm.Open(mysql.Open("root:zkf123456@tcp(127.0.0.1:3306)/postat?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "postat_",
			SingularTable: true,
		},

		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		return temp
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return temp
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// 获取数据
	fmt.Println(db)

	err = db.Model(MyData{}).Where("name = ?", "Lido").Preload("DataList").Find(&temp).Error
	fmt.Println(err)

	return temp

}

type stEth struct {
	CurrentDay              int64
	Amount                  decimal.Decimal
	ChangeAmountCurrentDay  decimal.Decimal
	ChangeAmountAll         decimal.Decimal
	Holders                 decimal.Decimal
	ChangeHoldersCurrentDay decimal.Decimal
	ChangeHoldersAll        decimal.Decimal
}

// 写入表格
func writeStEthData(data []stEth) {
	File, _ := os.OpenFile("./steth.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	defer File.Close()

	//创建写入接口
	WriterCsv := csv.NewWriter(File)
	WriterCsv.Write([]string{"时间", "质押数量", "质押一天改变数量", "质押总改变数量", "holders", "holders一天改变数量", "holders总改变数量"})
	for _, v := range data {
		//写入一条数据，传入数据为切片(追加模式)

		cd := time.Unix(v.CurrentDay, 0).Format("2006-01-02")

		WriterCsv.Write([]string{cd,
			v.Amount.String(),
			v.ChangeAmountCurrentDay.String(),
			v.ChangeAmountAll.String(),
			v.Holders.String(),
			v.ChangeHoldersCurrentDay.String(),
			v.ChangeHoldersAll.String()})
	}

	WriterCsv.Flush() //刷新，不刷新是无法写入的
	log.Println("数据写入成功...")
}
