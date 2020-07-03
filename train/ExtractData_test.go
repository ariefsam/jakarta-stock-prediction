package train_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/jakarta-stock-prediction/train"
)

func TestExtractData(t *testing.T) {

	stockData := []train.StockData{
		train.StockData{
			1500, 100, 120, 100, 110, 12000,
		},
		train.StockData{
			1501, 110, 150, 100, 140, 15000,
		},
		train.StockData{
			1502, 140, 180, 120, 170, 13000,
		},
		train.StockData{
			1503, 170, 200, 150, 200, 13000,
		},
	}
	expected := []float64{
		(stockData[3][2] - stockData[2][4]) / stockData[2][4],
		(stockData[2][2] - stockData[1][4]) / stockData[1][4],
	}
	expected2 := []float64{
		(stockData[3][2] - stockData[2][4]) / stockData[2][4],
		(stockData[2][2] - stockData[1][4]) / stockData[1][4],
		(stockData[1][2] - stockData[0][4]) / stockData[0][4],
	}
	extractDataInput := train.ExtractHighToClose(stockData, 2)
	assert.Equal(t, expected, extractDataInput)
	extractDataInput = train.ExtractHighToClose(stockData, 3)
	assert.Equal(t, expected2, extractDataInput)
}
