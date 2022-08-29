# Youtube web scraper using puppeteer 🔎

## Installation and usage

1. Move to `./scraper`

```shell script
cd ./scraper
```

2. Install dependencies

```shell script
npm i
```

3. Enter your queries into `queries.js` file. Each query must contain a topic (Used for console logging) and a youtube search url.

4. Run start command or manually execute the scraper.js file

```shell script
npm start
```

or

```shell script
node scraper.js
```

5. Finally, you can find the results on `data.json` file **(Average time needed is 2.45 minutes to get data from 145 videos), but this can vary depending on your machine specs and your internet connection**

## Dataset

The provided dataset contains data extracted from `2175 youtube videos` about `15 different topics / categories`. Each JSON object contain the video title, video description, video tags, video url, and the url of the video thumbnail. Data is distributed as following (This can variate slightly (Between 1 or 10 more or less videos) due to recent updates):

| Topic                     | N. of data extracted |
| ------------------------- | -------------------- |
| Development methodologies | 145                  |
| Song mix playlist         | 147                  |
| Data structures           | 142                  |
| Ancient history           | 147                  |
| Football highlights       | 148                  |
| Digital art               | 146                  |
| Movie trailer             | 149                  |
| Tedx talks                | 143                  |
| Web development           | 145                  |
| Gameplay walkthrough      | 145                  |
| What would happen         | 141                  |
| Data mining               | 148                  |
| Machine learning          | 146                  |
| Traveling vlog            | 147                  |
| Colombian touristic       | 146                  |
