import requests

def send_request():
    response = requests.get(
        url='https://app.scrapingbee.com/api/v1/',
        params={
            'api_key': '',
            'url': 'https://www.bestbuy.com/site/playstation-5/ps5-consoles/pcmcat1587395025973.c?id=pcmcat1587395025973', 
        },
        
    )
    print('Response HTTP Status Code: ', response.status_code)
    textFile = open("coin.html", "w+")
    textFile.write(str(response.content))
    textFile.close()
    print("Done writing to file")
send_request()