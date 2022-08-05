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

