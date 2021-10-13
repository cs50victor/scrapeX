import time, json
from datetime import datetime
from utils.makeRequests import MakeRequest
from utils.parseJson import requestData
import concurrent.futures

apis = [
    # coinbase top gainer -> resolution=day/week/month/year, (percentage_change * 100)%
    {"name":"coinbase-day","url":"https://www.coinbase.com/api/v2/assets/search?base=USD&country=US&filter=top_gainers&include_prices=true&limit=100&order=desc&page=1&query=&resolution=day&sort=percent_change"},
    {"name":"coinbase-week","url":"https://www.coinbase.com/api/v2/assets/search?base=USD&country=US&filter=top_gainers&include_prices=true&limit=100&order=desc&page=1&query=&resolution=week&sort=percent_change"},
    {"name":"coinbase-month","url":"https://www.coinbase.com/api/v2/assets/search?base=USD&country=US&filter=top_gainers&include_prices=true&limit=100&order=desc&page=1&query=&resolution=month&sort=percent_change"},
    {"name":"coinbase-year","url":"https://www.coinbase.com/api/v2/assets/search?base=USD&country=US&filter=top_gainers&include_prices=true&limit=100&order=desc&page=1&query=&resolution=year&sort=percent_change"},
    # binance , timeframe=7d/24h/30d
    {"name":"binance-day","url":"https://www.binance.com/bapi/composite/v1/public/promo/cmc/cryptocurrency/spotlight?dataType=2&timeframe=24h"},
    {"name":"binance-week","url":"https://www.binance.com/bapi/composite/v1/public/promo/cmc/cryptocurrency/spotlight?dataType=2&timeframe=7d"},
    {"name":"binance-month","url":"https://www.binance.com/bapi/composite/v1/public/promo/cmc/cryptocurrency/spotlight?dataType=2&timeframe=30d"},
]

requestStatuses = []

def parseJson(api):
    name, url = api.get("name"),api.get("url")
    
    r =  MakeRequest(url)
    requestStatuses.append({"Name": f'{name}',"Status": f'{r.response_data().get("status_code")}', "Agent": f'{r.response_data().get("request_Header").get("User-Agent")}'})
    mainJson = json.loads(r.response_data().get("main"))

    fileName = f"apiJson/{name}.json"
    with open(fileName, "w") as outputFile:
        json.dump(mainJson, outputFile, indent = 4, ensure_ascii=False)
    
    return

def getNewJson():
    #start timer
    start = time.perf_counter()

    # concurrently make request to api list and parse json responses
    with concurrent.futures.ThreadPoolExecutor() as executor:
        executor.map(parseJson, apis)
    
    # write the status of all request made to the json file 
    mainResJson = requestData()
    mainResJson["statuses"] = requestStatuses

    #end timer
    end = time.perf_counter() - start
    date = datetime.now().strftime("%d/%m/%Y %H:%M:%S")
    mainResJson["Time Taken"] = f"{end} | {date}"

    with open("./5am.json", "w") as outputFile:
        json.dump(mainResJson, outputFile, indent = 4, ensure_ascii=False)
    
    return 

if __name__ == "__main__":
    getNewJson()
