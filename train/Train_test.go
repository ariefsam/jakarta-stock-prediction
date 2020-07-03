package train_test

import (
	"testing"

	"github.com/ariefsam/jakarta-stock-prediction/train"
)

func TestTrain(t *testing.T) {
	totalDay := 10
	train.Train("data/", train.StockList, "data/modelTest.json", totalDay)
	// network := persist.FromFile("data/modelTest.json")
	// assert.NotNil(t, network.Enters[0])
}
