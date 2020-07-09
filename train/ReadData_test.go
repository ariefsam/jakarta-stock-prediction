package train_test

import (
	"testing"

	"github.com/ariefsam/jakarta-stock-prediction/train"
	"github.com/stretchr/testify/assert"
)

func TestReadData(t *testing.T) {

	var stockData []train.StockData
	stockData = train.ReadData("./ICBP.json")
	expected := train.StockData{
		1593648000000, 9250.0000, 9550.0000, 9225.0000, 9550.0000, 7749600,
	}
	assert.Equal(t, expected, stockData[0])

}
