package utils

import (
	"math"
	"time"
	"reflect"
	"github.com/tidwall/gjson"
)

type bestTokenTiming struct {
	BestBuy struct {
		Price float64         `json:"price"`
		Time  time.Time       `json:"time"`
	} `json:"bestBuy"`
	BestSell struct {
		Price float64         `json:"price"`
		Time  time.Time       `json:"time"`
	} `json:"bestSell"`
	MaxProfit       float64   `json:"maxProfit"`
	PercentIncrease float64   `json:"percentIncrease"`
}

type TokenData struct {
	Uuid        string        `json:"uuid"`
	Name        string        `json:"name"`
	Symbol      string        `json:"symbol"`
	Image       string        `json:"image"`
	Price       float64       `json:"price"`
	Color       string       `json:"color"`
	Time        time.Time     `json:"timestamp"`
	CreatedOn   string        `json:"created"`
	MoreInfo    []interface{} `json:"moreInfo"`
	Description string        `json:"description"`
	Supply      float64       `json:"circulatingSupply"`
	Volume24h   float64       `json:"volume24h"`
	MarketCap   float64       `json:"marketCap"`
	BestTimes   struct {      
		Hour  bestTokenTiming  `json:"hour"`
		Day   bestTokenTiming  `json:"day"`  
		Week  bestTokenTiming  `json:"week"`
		Month bestTokenTiming  `json:"month"`
		Year  bestTokenTiming  `json:"year"`
		All   bestTokenTiming  `json:"all"`
	} `json:"bestTimes"`
}
func (t TokenData) IsEmpty() bool {
	return reflect.DeepEqual(t,TokenData{})
}


func ParseCoinbaseTknInfo(tokenPrices []byte, tknIds map[string]gjson.Result) TokenData {

	
	var listing TokenData

	if tokenPrices != nil {
		moreTknInfo := gjson.ParseBytes(tokenPrices).Get("data.assetByUuid")
		uuid := moreTknInfo.Get("uuid").String()

		token, uuidExists := tknIds[uuid]
		if(uuidExists){
			listing.Uuid = token.Get("id").String()
			listing.Name = token.Get("name").String()
			listing.Symbol = token.Get("symbol").String()
			listing.Color = token.Get("symbol").String()
			listing.Image = token.Get("image_url").String()
			listing.MarketCap = token.Get("market_cap").Float()
			listing.Volume24h = token.Get("volume_24h").Float()
			listing.MoreInfo = token.Get("resource_urls").Value().([]interface{})
			listing.CreatedOn = token.Get("launched_at").String()
			listing.Description = token.Get("description").String()
			listing.Supply = token.Get("circulating_supply").Float()
	
			listing.Price = moreTknInfo.Get("latestQuote.price").Float()
			listing.Time = moreTknInfo.Get("latestQuote.timestamp").Time()
	
			listing.BestTimes.Hour = maxProfit(moreTknInfo.Get("priceDataForHour.quotes").Array())
			listing.BestTimes.Day = maxProfit(moreTknInfo.Get("priceDataForDay.quotes").Array())
			listing.BestTimes.Week = maxProfit(moreTknInfo.Get("priceDataForWeek.quotes").Array())
			listing.BestTimes.Month = maxProfit(moreTknInfo.Get("priceDataForMonth.quotes").Array())
			listing.BestTimes.Year = maxProfit(moreTknInfo.Get("priceDataForYear.quotes").Array())
			listing.BestTimes.All = maxProfit(moreTknInfo.Get("priceDataForAll.quotes").Array())
		}
	}

	return listing
}

func maxProfit(quotes []gjson.Result) bestTokenTiming {

	var timing bestTokenTiming
	timing.MaxProfit = 0.0
	timing.BestBuy.Price = math.Inf(1)

	for _, quote := range quotes {
		currPrice := quote.Get("price").Float()
		currPriceTime := quote.Get("timestamp").Time()

		timing.BestBuy.Price = math.Min(timing.BestBuy.Price, currPrice)
		if currPrice == timing.BestBuy.Price {
			timing.BestBuy.Time = currPriceTime
		}
		currProfit := currPrice - timing.BestBuy.Price
		timing.MaxProfit = math.Max(timing.MaxProfit, currProfit)
		if currProfit == timing.MaxProfit {
			timing.BestSell.Price = currPrice
			timing.BestSell.Time = currPriceTime
		}
	}

	timing.PercentIncrease = 100 * ((timing.BestSell.Price - timing.BestBuy.Price) / timing.BestBuy.Price)

	return timing
}
