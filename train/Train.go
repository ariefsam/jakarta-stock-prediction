package train

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/engine"
	"github.com/NOX73/go-neural/persist"
)

func Train(folderName string, stockList []string, fileName string, lastDayToCount int, totalDay int) {
	log.Println("Folder name:", folderName)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var day, lastDataToCount int

	totalInput := len(stockList) * (lastDataToCount)

	var network *neural.Network
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("File exists\n")
		network = persist.FromFile(fileName)

	} else {
		network = neural.NewNetwork(totalInput, []int{2, 2, len(stockList)})
		network.RandomizeSynapses()
		fmt.Printf("File does not exist\n")
	}

	log.Println("Total input:", totalInput)

	engine := engine.New(network)
	engine.Start()

	totalIteration := 100000
	go func() {
		for i := 0; i < totalIteration; i++ {
			if (i % 1000) == 0 {
				log.Println("=========\nIterasi ke", i)
			}
			day = 0
			for {
				if day > totalDay {
					break
				}

				inputNeural := []float64{}
				outputNeural := []float64{}

				for _, code := range stockList {
					stockData := ReadData(folderName + code + ".json")
					if len(stockData) < totalDay {
						log.Println("Saham ", code, " hanya sampai ", len(stockData), " hari")
					}
					extract := ExtractCloseToClose(stockData[:len(stockData)-day], (lastDataToCount + 1))

					outputNeural = append(outputNeural, extract[0])
					inputNeural = append(inputNeural, extract[1:]...)
				}

				engine.Learn(inputNeural, outputNeural, 1)
				// persist.ToFile(fileName, network)
				day++
			}
			persist.ToFile(fileName, network)
		}

		log.Println("Training DONE")
	}()

	time.Sleep(1 * time.Second)

	for {
		inputToPredict := []float64{}
		for _, code := range stockList {

			// log.Println(code)
			stockData := ReadData(folderName + code + ".json")
			extract := ExtractCloseToClose(stockData, (lastDataToCount))

			inputToPredict = append(inputToPredict, extract...)
		}
		out := engine.Calculate(inputToPredict)
		var maxout float64
		maxout = 0
		maxidx := ""
		naik := []string{}
		log.Print("\n\n\n\n\n")
		for index, code := range stockList {

			stockData := ReadData(folderName + code + ".json")
			// log.Println(stockData[0])
			// panic("x")
			if maxout < out[index] {
				maxout = out[index]
				maxidx = code
			}
			if out[index] >= 0.01 {
				naik = append(naik, code+" "+fmt.Sprintln(out[index]*float64(stockData[0][4])+float64(stockData[0][4]), "dari", stockData[0][4], "-", out[index]*100, "%"))
			}

			log.Println("Prediksi harga penutupan "+code+": ", out[index], "yaitu:", out[index]*float64(stockData[0][4])+float64(stockData[0][4]), "dari", stockData[0][4])
			// log.Println("Prediksi harga tertinggi "+code+": ", out[index], "yaitu:", out[index]*float64(stockData[0][4])+float64(stockData[0][4]), "dari", stockData[0][4])
			// log.Println(stockData)
		}
		log.Printf("naik di atas 1persen: \n%+v", naik)
		log.Println("belilah saham ", maxidx)

		time.Sleep(20 * time.Second)
	}
}
