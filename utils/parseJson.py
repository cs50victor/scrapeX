import json 

def parseWebull():
    webullJsonList =  ["5am/webull-day.json","5am/webull-week.json","5am/webull-month.json","5am/webull-year.json"]
    payload = {}
    count = 1
    for webull in webullJsonList:
        f = open(webull)
        data = json.load(f).get("data")
        
        if count > 1:
            for stock in data:
                stockSymbol = stock["ticker"]["symbol"]
                if payload.get(stockSymbol):
                    payload[stockSymbol]["consistency"]+=1

        else:
            for stock in data[0:10]:
                listing = {}
                stockData = stock["ticker"]
                stockValues = stock["values"]

                listing["name"],listing["symbol"]= stockData["name"],stockData["symbol"]
                listing["image"] = "https://images.unsplash.com/photo-1593672755342-741a7f868732?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1050&q=80"
                listing["buyHere"]= ["WeBull"]
                listing["price"] = stockData["close"]
                listing["percentageChange"]= round(float(stockValues["changeRatio"])*100,2)
                listing["backers"]= ["WeBull"]
                listing["consistency"]= 1
                listing["reason"]="----"
                listing["moreInfo"]= f"This stock has increased by +${stockData['change']} in the past 24 hours.\nMarketValue: ${stockData['marketValue']}\nVolume: {stockData['volume']}\nTurn Over Rate: {round(float(stockData['turnoverRate'])*100,2)}%\nYear-High: ${stockData.get('fiftyTwoWkHigh')}\nYear-Low: ${stockData.get('fiftyTwoWkLow')}\nLowest-Price Today:${stockData['low']}\nHighest-Price Today: ${stockData['high']}."
                payload[stockData["symbol"]]=listing        
        count+=1

    with open("5am/res/stocks.json", "w") as outputFile:
        json.dump({"listings":[payload[stock] for stock in payload]}, outputFile, indent = 4, ensure_ascii=False)

def parseCoinbase():

    coinbaseList =  ["5am/coinbase-day.json","5am/coinbase-week.json","5am/coinbase-month.json","5am/coinbase-year.json"]
    payload = {}
    count = 1
    for coin in coinbaseList:
        f = open(coin)
        data = json.load(f).get("data")
        
        if count > 1:
            for stock in data:
                stockSymbol = stock["symbol"]
                if payload.get(stockSymbol):
                    payload[stockSymbol]["consistency"]+=1

        else:
            for stock in data[0:10]:
                listing = {}

                listing["name"],listing["symbol"]= stock["name"],stock["symbol"]
                listing["image"] = stock["image_url"]
                listing["buyHere"]= ["Coinbase"] if (stock["listed"]) else ["not on Coinbase"]
                listing["price"] = stock["latest"]
                listing["percentageChange"]= round(float(stock["percent_change"])*100,2)
                listing["backers"]= ["Coinbase"]
                listing["consistency"]= 1
                listing["reason"]="----"
                lastestPrice = stock['latest_price']['percent_change']
                listing["moreInfo"]= f"This stock was launched {stock['launched_at']}\nDescription: {stock['description']}\nMarket-Cap: ${stock['market_cap']}\nVolume(24hrs): ${stock['volume_24h']}\nCirculating Supply: ${round(float(stock['circulating_supply'])*100,2)}\n.Percentage Changes:\n Hour: {round(float(lastestPrice['hour'])*100,5)}%\n Day: {round(float(lastestPrice['day'])*100,5)}%\n Week: {round(float(lastestPrice['week'])*100,5)}%\n Month: {round(float(lastestPrice['month'])*100,5)}%\n Day: {round(float(lastestPrice['year'])*100,5)}%\n All Time: {round(float(lastestPrice['all'])*100,5)}%"
                payload[stock["symbol"]]=listing        
        count+=1

    """ 
    print("\n\nPayload size-->",len(payload))
    for load in payload:
        print(payload[load]["symbol"],payload[load]["consistency"])
    """

    with open("5am/res/stable-crypto.json", "w") as outputFile:
        json.dump({"listings":[payload[stock] for stock in payload]}, outputFile, indent = 4, ensure_ascii=False)

def parseBinance():
    binanceList =  ["5am/binance-day.json","5am/binance-week.json","5am/binance-month.json"]
    payload = {}
    count = 1
    for coin in binanceList:
        f = open(coin)
        data = json.load(f)["data"]["body"]["data"]
        
        if count > 1:
            for stock in data["gainerList"]:
                stockSymbol = stock["symbol"]
                if payload.get(stockSymbol):
                    payload[stockSymbol]["consistency"]+=1

            for stock in data["loserList"]:
                stockSymbol = stock["symbol"]
                if payload.get(stockSymbol):
                    payload[stockSymbol]["consistency"]-=2

        else:
            for stock in data["gainerList"][0:10]:
                listing = {}

                listing["name"],listing["symbol"]= stock["name"],stock["symbol"]
                listing["image"] = f"https://s2.coinmarketcap.com/static/img/coins/64x64/{stock['id']}.png"
                listing["buyHere"]= ["Binance / Uniswap"]
                pricing = stock['priceChange']
                listing["price"] = pricing["price"]
                listing["percentageChange"]= round(float(pricing["priceChange24h"]),2)
                listing["backers"]= ["Binance"]
                listing["consistency"]= 0
                listing["reason"]="----"
                
                listing["moreInfo"]= f"More Pricing Details:\nVolume(24hrs): ${round(float(pricing['volume24h']),2)}\n.Percentage Changes:\n Hour: {round(float(pricing['priceChange1h']),5)}%\n Day: {round(float(pricing['priceChange24h']),5)}%\n Week: {round(float(pricing['priceChange7d']),5)}%\n Month: {round(float(pricing['priceChange30d']),5)}%"
                payload[stock["symbol"]]=listing        
        count+=1

    with open("5am/res/unstable-crypto.json", "w") as outputFile:
        json.dump({"listings":[payload[stock] for stock in payload]}, outputFile, indent = 4, ensure_ascii=False)
