const startBrowser = require('./puppeteer');
const fs = require('fs');
const axios = require('axios').default;

const { urls } = require('./queries');
const data = require('../data.json');

// Control variables
let timesAcc = 0;
let lengthAcc = 0;

function updateJson(video) {
	// Get JSON
	const data = fs.readFileSync('data.json');
	const parsedJSON = JSON.parse(data);

	// Add data
	parsedJSON['videos'].push(video);

	// Write json
	fs.writeFile('data.json', JSON.stringify(parsedJSON), (err) => {
		// error checking
		if (err) throw err;
	});
}

async function getVideosUrl(min, page) {
	let videos_length = 0;

	while (videos_length < min) {
		// Scroll to the end of the document (Charge more videos)
		await page.evaluate((_) => {
			window.scrollBy(0, window.innerHeight);
		});

		// Update videos length
		const videos = await page.$$('ytd-video-renderer');
		videos_length = videos.length;
	}

	// Get links
	const links = await page.$$eval('#video-title[href]', (anchor_list) =>
		anchor_list.map((anchor) => anchor.href)
	);

	return links;
}

async function getVideoData(videoUrl) {
	console.log('🟩 Parsing video --> ', videoUrl);

	// Make the axios http request
	const html = await await (await axios.get(videoUrl)).data;

	try {
		// Get data
		let description = html.match(/"shortDescription":"(.*?)"/)[0]; // Match regexp
		description = description.slice(
			description.indexOf(':') + 2,
			description.length - 1
		);

		const title = html.match(/<meta name="title"[^>]+content="(.*?)"/)[1];

		const keywords = html.match(/<meta name="keywords"[^>]+content="(.*?)"/)[1];

		const thumbnail = html.match(/<link rel="image_src"[^>]+href="(.*?)"/)[1];

		// Build final object
		const video = {
			url: videoUrl,
			title: title.replace(/\\+n/g, ' ').replace(/\s\s+/g, ' ').trim(),
			description: description.replace(/\\+n/g, ' ').replace(/\s\s+/g, ' ').trim(),
			tags: keywords,
			thumbnail,
		};

		updateJson(video);

		console.log('🟩 End video --> ', videoUrl, '\n');
	} catch (error) {
		console.log('🟥 There was an error with --> ', videoUrl, '\n');
	}
}

async function scrapper(list_index) {
	const topic_url_pair = urls[list_index];

	let counter = 1;

	// Start the browser and navigate to the page
	const startDate = new Date();

	const browser = await startBrowser();

	const page = await browser.newPage();
	await page.setViewport({ width: 1920, height: 1080, deviceScaleFactor: 1 });
	await page.goto(topic_url_pair.url);

	// Get links array
	const links = await getVideosUrl(140, page);
	await browser.close();

	// Get information for each video
	for (let i = 0; i < links.length; i++) {
		console.log(`Topic: ${topic_url_pair.topic} Video number: ${counter}`);
		await getVideoData(links[i]);
		counter++;
	}

	await browser.close();

	const endDate = new Date();
	const timeElapsed = Math.abs(endDate - startDate) / 60000;

	timesAcc = timesAcc + timeElapsed;
	lengthAcc = lengthAcc + links.length;

	console.log(
		`⏳ Data from ${links.length} videos about ${topic_url_pair.topic} was scrapped in ${timeElapsed} minutes \n`
	);

	// Next topic
	if (list_index < urls.length - 1) {
		scrapper(++list_index);
	} else {
		// Print final message
		console.log(
			`🏁 Averages: ${lengthAcc / urls.length} videos were scrapped in ${
				timesAcc / urls.length
			} minutes`
		);
	}
}

// Run
scrapper(0);
