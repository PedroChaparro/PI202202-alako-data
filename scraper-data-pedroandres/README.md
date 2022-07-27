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
npm run start
```

or

```shell script
node scraper.js
```

5. Finally, you can find the results on `data.json` file.

## Dataset

The provided dataset constains data extracted from `2178 youtube videos` about `15 different topics / categories`. Each JSON object contain the video title, video description, video tags, video url, and the url of the video thumbnail.