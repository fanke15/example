package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
)

type Data struct {
	Addr string
	Num  int
}

func main() {

	d := make([]Data, 0)
	a1 := Read("1.csv")
	a2 := Read("2.csv")
	a3 := Read("3.csv")

	for k, v := range a2 {
		if _, ok := a1[k]; ok {
			a1[k] = a1[k] + v
		} else {
			a1[k] = v
		}
	}
	for k, v := range a3 {
		if _, ok := a1[k]; ok {
			a1[k] = a1[k] + v
		} else {
			a1[k] = v
		}
	}

	for k, v := range a1 {
		d = append(d, Data{k, v})
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i].Num > d[j].Num
	})

	Write("addr.csv", d)

}

func Read(path string) map[string]int {
	f := make(map[string]int)
	//打开文件(只读模式)，创建io.read接口实例
	opencast, _ := os.Open(path)

	defer opencast.Close()

	//创建csv读取接口实例
	ReadCsv := csv.NewReader(opencast)

	//读取所有内容
	ReadAll, _ := ReadCsv.ReadAll() //返回切片类型：[[s s ds] [a a a]]

	for _, v := range ReadAll {
		if len(v[4]) > 20 {
			if _, ok := f[v[4]]; ok {
				f[v[4]] = f[v[4]] + 1
			} else {
				f[v[4]] = 1
			}
		}
	}
	return f
}

func Write(path string, data []Data) {
	File, _ := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	defer File.Close()

	//创建写入接口
	WriterCsv := csv.NewWriter(File)
	WriterCsv.Write([]string{"eth1 address", "transactions num"})
	for _, v := range data {
		//写入一条数据，传入数据为切片(追加模式)
		WriterCsv.Write([]string{v.Addr, strconv.Itoa(v.Num)})
	}

	WriterCsv.Flush() //刷新，不刷新是无法写入的
	log.Println("数据写入成功...")
}
