import requests
from random import choice
from .userAgentsArray import userAgentList
from requests.exceptions import Timeout
#from requests_ip_rotator import ApiGateway, EXTRA_REGIONS

# things to consider - headers, proxies, ip address, JS rendering and captchas, speed/concorrency, 

s = requests.Session()

def randHeader():
    return {
        "Accept": "application/json,text/plain,text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
        "Accept-encoding": "gzip, deflate, br",
        "Sec-Fetch-Dest":	"document",
        "Sec-Fetch-Mode":	"navigate",
        "Sec-Fetch-Site":	"none",
        "DNT":"1",
        "Referer":"https://www.nytimes.com/",
        "sec-ch-ua":"'Chromium';v='92', ' Not A;Brand';v='99', 'Google Chrome';v='92'",
        "sec-ch-ua-mobile":"?1",
        "Sec-Fetch-User": "?1",
        "Upgrade-Insecure-Requests":"1",
        "Accept-language": "en-US,en;q=0.5",
        "User-Agent": choice(userAgentList),
    }

class MakeRequest():
    def __init__(self, url):
        self.url = url
        self.defaultTimeout=(10,13) # (connect timeout, response waiting time after client has established a connection)

        # gateway = ApiGateway(url, regions=EXTRA_REGIONS, access_key_id="ID", access_key_secret="SECRET")
        # gateway.start()
        # s.mount(url, gateway)
        self.response = s.get(self.url,headers=randHeader(),timeout=self.defaultTimeout, stream=True)
        self.ip,self.port = self.response.raw._connection.sock.getpeername()
        self.status, self.reason, self.headers, self.cookies= self.response.status_code, self.response.reason, self.response.headers, self.response.cookies
        self.connection, self.history, self.payload =self.response.connection, self.response.history, self.response.content
        self.req_header = self.response.request.headers
        # gateway.shutdown()

    def response_data(self):
        return {"status_code":self.status,"reason":self.reason, "request_Header":self.req_header,
                "response_Header":self.headers,"cookies":self.cookies, "ip":(self.ip,self.port),
                "connection":self.connection, "history":self.history, "main":self.payload}
         
