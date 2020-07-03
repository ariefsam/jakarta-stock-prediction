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
		1286409600000, 3000.0000, 3100.0000, 2850.0000, 2975.0000, 570465000,
	}
	assert.Equal(t, expected, stockData[0])

}
