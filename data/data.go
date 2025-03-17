package data

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var mtx sync.Mutex

var stockConfigs = map[string]bool{
	"AAPL": true,
	"AMZN": true,
	"GOOG": true,
	"META": true,
	"MSFT": true,
	"NFLX": true,
}

type StockPrice struct {
	Code      string
	Price     int64
	Timestamp time.Time
}

var stockPrices = map[string][]StockPrice{}

func init() {
	for code, isEnabled := range stockConfigs {
		if !isEnabled {
			continue
		}
		log.Printf("saham %s diaktifkan", code)
	}
}

func updateStock(code string) {
	for {
		mtx.Lock()
		if !stockConfigs[code] {
			mtx.Unlock()
			break
		}

		time.Sleep(1 * time.Second)
		current, exists := stockPrices[code]
		if !exists {
			hargaAwal := 1000
			waktuAwal := time.Now()

			stockPrices[code] = []StockPrice{
				{
					Code:      code,
					Price:     int64(hargaAwal),
					Timestamp: waktuAwal,
				},
			}
			log.Printf("saham %s baru di tambahwakn untuk pertama kali dengan harga %d", code, hargaAwal)
			mtx.Unlock()
			continue
		}

		itemTerakhir := current[len(stockPrices)-1]
		harga := randomizePrice(itemTerakhir.Price)

		stockPrices[code] = append(stockPrices[code], StockPrice{
			Code:      code,
			Price:     harga,
			Timestamp: time.Now(),
		})
		log.Printf("saham %s di update dengan harga %d", code, harga)
	}
	mtx.Unlock()
}

func randomizePrice(price int64) int64 {
	operasi := rand.Intn(2)
	jumlah := rand.Int63n(100)
	if operasi == 0 {
		return price - jumlah
	}
	return price + jumlah
}

func ToggleStock(code string, isEnabled bool) {
	mtx.Lock()
	if isEnabled == stockConfigs[code] {
		mtx.Unlock()
		return
	}

	log.Printf("Mengubah status %s menjadi %t", code, isEnabled)
	if !stockConfigs[code] && isEnabled {
		stockConfigs[code] = true

		mtx.Unlock()
		go updateStock(code)
	}
	stockConfigs[code] = false
	mtx.Unlock()
}

func GetStockConfig() map[string]bool {
	return stockConfigs
}

func GetStockPrices(code string) []StockPrice {
	return stockPrices[code]
}
