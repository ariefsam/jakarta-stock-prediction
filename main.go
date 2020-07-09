package main

import "github.com/ariefsam/jakarta-stock-prediction/train"

func main() {
	totalDay := 30
	lastDayToCount := 100
	train.Train("train/data/", train.StockList, "train/data/modelTest39.json", lastDayToCount, totalDay)
}
