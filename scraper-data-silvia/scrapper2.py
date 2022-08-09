import json
import time

from selenium.webdriver.common.by import By
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.chrome.service import Service
from webdriver_manager.chrome import ChromeDriverManager
import requests
from bs4 import BeautifulSoup

def get_links(url, driver):
     driver.get(url)

     link_videos =  driver.find_elements(By.ID, "video-title")

     while len(link_videos) < 205:
          driver.execute_script("let scrollingElement = (document.scrollingElement || document.body);scrollingElement.scrollTop = scrollingElement.scrollHeight;")
          link_videos =  driver.find_elements(By.ID, "video-title")
     links = [l.get_attribute("href") for l in link_videos if l.get_attribute("href") != None]

     return links

topics = ['technology',
          'productivity',
          'song mix study',
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

     for link in get_links(url, driver):

          response = requests.get(link)

          soup = BeautifulSoup(response.text)

          keep_videos.append(
          {
          'url': link,
          'title': soup.find("meta", {"name": "title"})['content'],
          'description': soup.find("meta", {"name": "description"})['content'],
          'tags': soup.find("meta", {"name": "keywords"})['content'],
          #'thumbnail': driver.find_element(By.XPATH, "//link[@rel='image_src']").get_attribute("href"),
          })

     fin = time.time()

     print((fin - inicio)/60)
     print(fin - inicio)

print((fin - inicio)/60)
print(fin - inicio)
with open('./scraper-data-silvia/data.json', 'w+') as json_file:
          json.dump(keep_videos, json_file)
