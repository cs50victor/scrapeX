from requests_html import HTMLSession

session = HTMLSession()

url = input("Enter a url")
scrollDown = input("scrolldown int value: ")

r = session.get(url)

r.html.render(sleep=1, keep_page=True, scrolldown=scrollDown)

videos = r.html.find("#video-title")

for item in videos:
    video = {
        "title": item.text,
        "link": item.absolute_link
    }
    print(video)

