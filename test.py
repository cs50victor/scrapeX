import json, time
from utils.makeRequests import MakeRequest

def testFunc():
    

    start = time.perf_counter()
    
    x = MakeRequest("https://www.youtube.com/")
    stat = x.response_data().get("status_code")
    header = x.response_data().get("request_Header")
    res_header = x.response_data().get("response_Header")
    main = x.response_data().get("main")

    end = time.perf_counter() - start

    return f"Request Status: {stat}\nRequest Header:{header}\nResponse Header:{res_header}\nExecution Time:{end}"