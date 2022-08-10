import json
import time

from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.chrome.service import Service
from webdriver_manager.chrome import ChromeDriverManager
import requests
from bs4 import BeautifulSoup

topics = ['technology',
          'productivity',
          'a day in the life of',
          'recipes',
          'cleaning recommendations',
          'news',
          'sport news',
          'student vlogs',
          'study advices',
          'tasting food']

keep_videos = []

inicio = time.time()

for word in topics:
    options = Options()
    options.add_argument("--headless")
    options.add_argument("--incognito")
    driver = webdriver.Chrome(service=Service(ChromeDriverManager().install()), options=options)
    url=f'https://www.youtube.com/results?search_query={word}'

    driver.get(url)

    soup = BeautifulSoup(driver.page_source, features="lxml")

    html = soup.find_all(id="video-title")

    while len(html) < 230:
          driver.execute_script("let scrollingElement = (document.scrollingElement || document.body);scrollingElement.scrollTop = scrollingElement.scrollHeight;")
          soup = BeautifulSoup(driver.page_source, features="lxml")
          html =  soup.find_all("a", id="video-title", class_="yt-simple-endpoint")

    for l in html:
        if l != None:
            link = l.get('href')
            response = requests.get(f'https://www.youtube.com{link}')

            soup = BeautifulSoup(response.text, features="lxml")

            keep_videos.append(
            {
            'url': f'https://www.youtube.com{link}',
            'title': (soup.find("meta", {"name": "title"})['content']).strip().replace("\n", " "),
            'description': (soup.find("meta", {"name": "description"})['content']).strip().replace("\n", " "),
            'tags': soup.find("meta", {"name": "keywords"})['content'],
            'thumbnail': soup.find("link", {"rel": "image_src"})['href']
            })

    fin = time.time()

    print((fin - inicio)/60)
    print(fin - inicio)

print((fin - inicio)/60)
print(fin - inicio)
with open('./scraper-data-silvia/data.json', 'w+', encoding='utf8') as json_file:
          json.dump({'videos': keep_videos}, json_file, ensure_ascii=False)
