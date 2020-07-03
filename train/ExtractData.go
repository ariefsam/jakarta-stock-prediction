package train

func ExtractHighToClose(data []StockData, total int) (extract []float64) {
	length := len(data)
	lastIndex := length - total
	for i := (length - 1); i >= lastIndex; i-- {
		current := data[i]
		yesterday := data[i-1]
		diff := current[2] - yesterday[4]
		percent := diff / yesterday[4]
		extract = append(extract, percent)
	}
	return
}

func ExtractCloseToClose(data []StockData, total int) (extract []float64) {
	length := len(data)
	lastIndex := length - total
	for i := (length - 1); i >= lastIndex; i-- {
		current := data[i]
		yesterday := data[i-1]
		diff := current[4] - yesterday[4]
		percent := diff / yesterday[4]
		extract = append(extract, percent)
	}
	return
}
