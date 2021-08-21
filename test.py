import json
from utils.makeRequests import MakeRequest

x = MakeRequest("https://api.coinmarketcap.com/data-api/v3/cryptocurrency/spotlight?dataType=2&limit=10&rankRange=0&timeframe=24h")
print(x.response_data().get("status_code"))
print(x.response_data().get("request_Header"))
print(x.response_data().get("response_Header"))
print(x.response_data().get("main"))