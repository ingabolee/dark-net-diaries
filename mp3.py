import re
import httplib2
from bs4 import BeautifulSoup, SoupStrainer
http = httplib2.Http()


with open('urls.txt', "a") as file:  
    for i in range(1, 128):
        status, response = http.request(f'https://darknetdiaries.com/episode/{i}')
    
        for url in BeautifulSoup(response, parse_only=SoupStrainer('script'), features="lxml"):
            if ".mp3" in str(url):
                index = re.findall(r'\d{10}', str(url))[0]
                file.write(f"https://traffic.megaphone.fm/ADV{index}.mp3")
                file.write("\n")
                print(index)