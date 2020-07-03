package main

import "github.com/ariefsam/jakarta-stock-prediction/train"

func main() {
	totalDay := 100
	lastDayToCount := 100
	train.Train("train/data/", train.StockList, "train/data/modelTest19.json", lastDayToCount, totalDay)
}
