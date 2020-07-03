package train

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
)

func ReadData(filename string) (allData []StockData) {
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

	return
}

func reverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
