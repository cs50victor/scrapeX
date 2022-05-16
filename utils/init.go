package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

func getRandomUserAgent() string {
	rand.Seed(time.Now().Unix())

	userAgents := []string{

		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.2725; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.2313; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/5.0 (X11; U; OpenBSD i386; en-US; rv:1.9.1) Gecko/20090702 Firefox/3.5",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/174.78; U; ru) Presto/2.12.423 Version/12.16",
		"Palm750/v0005 Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.6)",
		"Mozilla/5.0 (X11; OpenBSD i386; rv:72.0) Gecko/20100101 Firefox/72.0",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.2766; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/4.0 (Macintosh)",
		"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Mobile Safari/537.36",
		"Mozilla/5.0 (Macintosh; U; PPC; en-US; rv:1.2b) Gecko/20021016",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/29.3709; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21595/36.1409; U; en) Presto/2.12.423 Version/12.16",
		"Mozilla/4.79 (Macintosh; U; PPC)",
		"Mozilla/5.0 (X11; OpenBSD amd64) AppleWebKit/538.15 (KHTML, like Gecko) Version/8.0 Safari/538.15",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.3950; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/5.0 (X11; OpenBSD amd64; rv:88.0) Gecko/20100101 Firefox/88.0",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.2359; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.2794; U; en) Presto/2.8.119 Version/11.10",
		"PalmCentro/v0001 Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/Palm-D061; Blazer/4.5) 16;320x320 UP.Link/6.3.1.17.0",
		"Mozilla/5.0 (X11; OpenBSD amd64; rv:69.0) Gecko/20100101 Firefox/69.0",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.3392; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/5.0 (X11; OpenBSD amd64; rv:43.0) Gecko/20100101 Firefox/43.0 SeaMonkey/2.40",
		"Mozilla/5.0 (X11; U; OpenBSD i386; en-US) AppleWebKit/533.3 (KHTML, like Gecko) Chrome/5.0.359.0 Safari/533.3",
		"Mozilla/5.0 (X11; OpenBSD amd64; rv:58.0) Gecko/20100101 Firefox/58.0",
		"Mozilla/4.0 (compatible; MSIE 5.2; Mac_PowerPC)",
		"Mozilla/5.0 (X11; OpenBSD amd64; rv:59.0) Gecko/20100101 Firefox/59.0",
		"Mozilla/5.0 (X11; U; OpenBSD i386; en-US; rv:1.9.2.8) Gecko/20101230 Firefox/3.6.8",
		"Mozilla/4.0 (compatible; MSIE 5.1b1; Mac_PowerPC)",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/Palm-D062; Blazer/4.5) Palm 690p 16;320x320",
		"Palm680/RC1 Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/Palm-D053; Blazer/4.5) 16;320x320 UP.Link/6.3.0.0.0",
		"Mozilla/5.0 (Linux; Android 8.1.0; CompalMediatekArgon1 Build/O11019; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.157 Safari/537.36",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.2443; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.22 (X11; OpenBSD i386; U; en)",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.2647; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/27.1940; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.3030; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/29.3417; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/4.0 (compatible;MSIE 6.0;Windows95;PalmSource) Netfront/3.0;8;320x320",
		"Mozilla/4.0 (compatible; MSIE 5.0; Mac_PowerPC; AtHome021)",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/Palm-TunX; Blazer/4.3) 16;320x448",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/30.3214; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.1977; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/29.3530; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/5.0 (X11; OpenBSD amd64; rv:62.0) Gecko/20100101 Firefox/62.0",
		"Mozilla/5.0 (X11; OpenBSD amd64; rv:68.0) Gecko/20100101 Firefox/68.0",
		"Mozilla/5.0 (X11; OpenBSD amd64; rv:82.0) Gecko/20100101 Firefox/82.0",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21595/28.3590; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/4.0 (compatible; MSIE 5.23; Mac_PowerPC; SV1)",
		"Mozilla/5.0 (Macintosh; U; PPC; en-US; rv:1.0.1) Gecko/20020823 Netscape/7.0 (OEM-SBC)",
		"Mozilla/5.0 (Macintosh; U; PPC;) Gecko DEVONtech",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.2173; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/hspr-H102; Blazer/4.2) 16;320x320",
		"Mozilla/5.0 (X11; OpenBSD amd64; rv:78.0) Gecko/20100101 Firefox/78.0",
		"Mozilla/4.0 (compatible; MSIE 67.0; Windows 98; PalmSource/hspr-H102; Blazer/4.0) 16;320x320(Linux LLC 1.2)",
		"Palm680/RC1 Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/Palm-D053; Blazer/4.5) 16;320x320",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.3126; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/Palm-TunX; Blazer/4.3) 16;320x320",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/30.3112; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/5.0 (Macintosh; U; PPC; fr-FR; rv:1.0.2) Gecko/20030208 Netscape/7.02",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/27.1486; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/5.0 (compatible; Konqueror/4.1; OpenBSD) KHTML/4.1.4 (like Gecko)",
		"Mozilla/5.0 (X11; OpenBSD amd64; rv:67.0) Gecko/20100101 Firefox/67.0",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/Palm-TunX; Blazer/4.1) 16;320x448",
		"Mozilla/4.0 (PDA; PalmOS/sony/model prmr/Revision:1.1.54 (en)) NetFront/3.0 ubunto/2.0.1.16",
		"Mozilla/5.0 (X11; U; OpenBSD i386; en-US; rv:1.8.0.5) Gecko/20060819 Firefox/1.5.0.5",
		"Mozilla/5.0 (X11; OpenBSD amd64; rv:30.0) Gecko/20100101 Firefox/30.0",
		"Mozilla/4.0 (compatible; MSIE 5.1b1; AOL 5.1; Mac_PowerPC)",
		"Mozilla/5.0 (X11; OpenBSD amd64; rv:66.0) Gecko/20100101 Firefox/66.0",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.2225; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/4.72 (Macintosh; U; PPC)",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.3182; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/5.0 (X11; OpenBSD x86_64; rv:76.0) Gecko/20100101 Firefox/76.0",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.1914; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.2075; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.3445; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/27.1993; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21595/25.657; U; en) Presto/2.5.25 Version/10.54",
		"Mozilla/5.0 (X11; U; OpenBSD i386; en-US; rv:1.7.10) Gecko/20050919 (No IDN) Firefox/1.0.6",
		"Mozilla/5.0 (Macintosh; U; PPC; en-US; rv:0.9.4) Gecko/20011022 Netscape6/6.2",
		"Mozilla/5.0 (X11; U; OpenBSD amd64; en-US; rv:1.9.0.1) Gecko/2008081402 Firefox/3.0.1",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/30.3061; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/5.0 (Linux; PalmOS 3.0) Chromium/53.0 Mobile",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21595/28.3692; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/29.3594; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/4.0 (compatible; MSIE 5.23; Macintosh; PPC) Escape 5.1.8",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/26.1424; U; ru) Presto/2.8.119 Version/10.54",
		"PalmCentro/v0001 Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/Palm-D061; Blazer/4.5) 16;320x320 UP.Link/6.3.1.17.06.3.1.17.0",
		"Mozilla/5.0 (X11; OpenBSD amd64; rv:76.0) Gecko/20100101 Firefox/76.0",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/27.1813; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/5.0 (Linux; Android 4.4.2; SPALM 7.8 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Safari/537.36",
		"Mozilla/5.0 (X11; U; OpenBSD i386; en-US; rv:1.7.0.13) Gecko/20060901",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.3234; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/5.0 (X11; OpenBSD x86_64; rv:77.0) Gecko/20100101 Firefox/77.0",
		"Mozilla/5.0 (Macintosh; U; PPC; en-US; rv:0.9.2) Gecko/20010726 Netscape6/6.1",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/Palm-D050; Blazer/4.3) 16;448x320",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.3692; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.2555; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.2144; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/5.0 (X11; OpenBSD; x86_64) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.2.2 Safari/605.1.15",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.3782; U; en) Presto/2.8.119 Version/11.10",
		"PalmCentro/v0001 Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/Palm-D061; Blazer/4.5) 16;320x320 UP.Link/6.3.1.20.0",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/29.3638; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows 984; PalmSource/hspr-H102; Blazer/4.0) 16;320x320(Linux LLC 1.2)",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/37.7011; U; fr) Presto/2.12.423 Version/12.16",
		"Mozilla/5.0 (Macintosh; U; PPC; ja-JP; rv:1.0.2) Gecko/20030208 Netscape/7.02",
		"Palm750/v0000 Mozilla/4.0 (compatible; MSIE 4.01; Windows CE; PPC; 240x320) UP.Link/6.3.0.0.0",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/30.3158; U; en) Presto/2.8.119 Version/11.10",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21595/28.3445; U; en) Presto/2.8.119 Version/11.10",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/hspr-H102; Blazer/4.0) 16;320x320(Linux LLC 1.2)",
		"Palm750/v0100 Mozilla/4.0 (compatible; MSIE 4.01; Windows CE; PPC; 240x320)",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/Palm-D050; Blazer/4.3) 16;320x320)",
		"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/22.387; U; ru) Presto/2.5.25 Version/10.54",
	}

	i := rand.Intn(len(userAgents))
	return userAgents[i]
}

func Hello(name string) string {
	var id int = rand.Intn(100)
	message := fmt.Sprintf("Hi, %v.(Id:%d ) Welcome! 🚀", name, id)
	return message
}

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func SendRequest(wg *sync.WaitGroup, client *http.Client, url string, c chan<- []byte) {
	defer wg.Done()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Referer", `https://www.coinbase.com/`)
	req.Header.Set("User-Agent", getRandomUserAgent())

	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	// Close the connection to reuse it
	defer resp.Body.Close()

	fmt.Println(resp.Status + ":\n\t" + url)
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	c <- body
}

func CommonCoinbaseTkns() []TokenData {

	coinbaseTokenUrls := []string{
		"https://www.coinbase.com/api/v2/assets/search?base=USD&country=US&filter=top_gainers&include_prices=true&limit=10&order=desc&page=1&query=&resolution=" + "day" + "&sort=percent_change",
		"https://www.coinbase.com/api/v2/assets/search?base=USD&country=US&filter=top_gainers&include_prices=true&limit=10&order=desc&page=1&query=&resolution=" + "week" + "&sort=percent_change",
		"https://www.coinbase.com/api/v2/assets/search?base=USD&country=US&filter=top_gainers&include_prices=true&limit=10&order=desc&page=1&query=&resolution=" + "month" + "&sort=percent_change",
		"https://www.coinbase.com/api/v2/assets/search?base=USD&country=US&filter=top_gainers&include_prices=true&limit=10&order=desc&page=1&query=&resolution=" + "year" + "&sort=percent_change",
	}

	// var numOfUrls int = len(coinbaseTokenUrls)

	var topTokens []TokenData

	client := httpClient()

	tknIds := make(map[string]gjson.Result)

	urlChan := make(chan []byte)
	tokenPricesChan := make(chan []byte)

	go func() {
		defer close(tokenPricesChan)
		var toknPricesWg sync.WaitGroup
		for tokenInfo := range urlChan {
			result := gjson.ParseBytes(tokenInfo).Get("data")

			//
			for _, tkn := range result.Array() {

				isListed := tkn.Get("listed").Bool()
				percent_change := math.Round((tkn.Get("percent_change").Float()) * 100)

				if isListed && (percent_change >= 100) {
					uuid := tkn.Get("id").String()

					_, exists := tknIds[uuid]
					if !exists {
						toknPricesWg.Add(1)
						tknIds[uuid] = tkn
						go SendRequest(&toknPricesWg, client, getTknPriceUrl(uuid), tokenPricesChan)
					}
				}
			}
		}
		toknPricesWg.Wait()
	}()

	go func() {
		var requestWg sync.WaitGroup
		for _, coinbaseUrl := range coinbaseTokenUrls {
			requestWg.Add(1)
			go SendRequest(&requestWg, client, coinbaseUrl, urlChan)
		}
		requestWg.Wait()
		close(urlChan)
	}()

	for {
		select {
		case _, open := <-urlChan:
			if !open {
				urlChan = nil
			}
		case tokenPrices, open := <-tokenPricesChan:
			tkn := ParseCoinbaseTknInfo(tokenPrices, tknIds)
			if !tkn.IsEmpty(){
				topTokens = append(topTokens, tkn)
			}
			
			if !open {
				tokenPricesChan = nil
			}
		}

		if urlChan == nil && tokenPricesChan == nil {
			break
		}
	}

	fmt.Println("here")
	return topTokens
}

func getTknPriceUrl(uuid string) string {
	return "https://www.coinbase.com/graphql/query?&operationName=useGetPricesForAssetPageQuery&extensions={%22persistedQuery%22:{%22version%22:1,%22sha256Hash%22:%2259282a0565bfbdc0477f69ad3ae4b687c93d75c808445386bfbfa70be7b4a976%22}}&variables={%22skip%22:false,%22uuid%22:%22" + uuid + "%22,%22currency%22:%22USD%22}"
}
