package main

import "github.com/ariefsam/jakarta-stock-prediction/train"

func main() {
	totalDay := 10
	lastDayToCount := 5
	train.Train("train/data/", train.StockList, "train/data/modelTest5.json", lastDayToCount, totalDay)
}
