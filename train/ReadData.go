package train

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
)

var Cache map[string][]StockData

func init() {
	Cache = map[string][]StockData{}
}

func ReadData(filename string) (allData []StockData) {
	if val, ok := Cache[filename]; ok {
		allData = val
		return
	}
	stockPrice, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return
	}

	err = json.Unmarshal(stockPrice, &allData)
	reverseSlice(allData)
	if err != nil {
		log.Println(err)
		return
	}
	Cache[filename] = allData
	return
}

func reverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
