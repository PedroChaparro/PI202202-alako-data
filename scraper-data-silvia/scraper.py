import json
import time
from pathlib import Path

from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.chrome.service import Service
from webdriver_manager.chrome import ChromeDriverManager
import requests
from bs4 import BeautifulSoup
import re

topics = ['technology news',
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
    print("----------------------------- Scrolling Youtube for videos -----------------------------")
    while len(html) < 215:
          driver.execute_script("let scrollingElement = (document.scrollingElement || document.body);scrollingElement.scrollTop = scrollingElement.scrollHeight;")
          soup = BeautifulSoup(driver.page_source, features="lxml")
          html =  soup.find_all("a", id="video-title", class_="yt-simple-endpoint")
    print("----------------------------- Starting scraping of topic -----------------------------", word)
    for l in html:
        link = l.get('href')
        if link != None:
            response = requests.get(f'https://www.youtube.com{link}')

            soup = BeautifulSoup(response.text, features="lxml")
            patron_description = re.compile('shortDescription":"(.*)","isCrawlable')

            keep_videos.append(
            {
                'url': f'https://www.youtube.com{link}',
                'title': (soup.find("meta", {"name": "title"})['content']).strip().replace('\\n', ' ').replace('\\', '').replace('  ', ' '),
                'description': (patron_description.findall(str(soup))[0]).strip().replace('\\n', ' ').replace('\\', '').replace('  ', ' '),
                'tags': soup.find("meta", {"name": "keywords"})['content'],
                'thumbnail': soup.find("link", {"rel": "image_src"})['href']
            })
            print(link, "---had been scrapped successfully---")
    print("----------------------------- Finishing scraping of topic -----------------------------", word)
    fin = time.time()

    print((fin - inicio)/60, "minutos parciales")
    print(fin - inicio, "segundos parciales")

print((fin - inicio)/60, "minutos totales")
print(fin - inicio, "segundos totales")

with open(Path(".\scraper-data-silvia\data.json"), 'w+', encoding='utf-8') as json_file:
          json.dump({'videos': keep_videos}, json_file, ensure_ascii=False)
