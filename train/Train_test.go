package train_test

import (
	"testing"

	"github.com/ariefsam/jakarta-stock-prediction/train"
)

func TestTrain(t *testing.T) {
	totalDay := 100
	totalSlide := 10
	train.Train("data/", train.StockList, "data/modelmodelTest.json", totalSlide, totalDay)
	// network := persist.FromFile("data/modelTest.json")
	// assert.NotNil(t, network.Enters[0])
}
