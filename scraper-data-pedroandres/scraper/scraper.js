const startBrowser = require('./puppeteer');
const fs = require('fs');
const { urls } = require('./queries');
const data = require('../data.json');

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

async function getVideosUrl(max, page) {
	let videos_length = 0;

	while (videos_length < max) {
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

async function getVideoData(browser, videoUrl) {
	console.log('🟩 Parsing video --> ', videoUrl);

	// Setup
	const page = await browser.newPage();
	await page.setViewport({ width: 1920, height: 1080, deviceScaleFactor: 1 });
	await page.goto(videoUrl);

	// Wait for the selectors
	await page.waitForSelector('h1 > yt-formatted-string');
	await page.waitForSelector('meta[name="keywords"]');
	await page.waitForSelector(
		'ytd-expander > #content > #description > yt-formatted-string'
	);

	// Get data
	const titleContainer = await page.$('h1 > yt-formatted-string');
	const title = await titleContainer.evaluate(
		(element) => element.textContent,
		titleContainer
	);

	const tagsContainer = await page.$('meta[name="keywords"]');
	const tags = await tagsContainer.evaluate(
		(element) => element.content,
		tagsContainer
	);

	const descriptionContainer = await page.$(
		'ytd-expander > #content > #description > yt-formatted-string'
	);
	const description = await descriptionContainer.evaluate(
		(element) => element.textContent,
		descriptionContainer
	);

	const thumbnail = `https://img.youtube.com/vi/${
		videoUrl.split('=')[1]
	}/maxresdefault.jpg`;

	const video = {
		url: videoUrl,
		title,
		description: description.replace(/\n/g, ' ').replace(/\s\s+/g, ' ').trim(),
		tags,
		thumbnail,
	};

	updateJson(video);

	console.log('🟩 End video --> ', videoUrl, '\n');

	await page.close();
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

	// Get information for each video
	for (let i = 0; i < links.length; i++) {
		console.log(`Topic: ${topic_url_pair.topic} Video number: ${counter}`);
		await getVideoData(browser, links[i]);
		counter++;
	}

	await browser.close();

	const endDate = new Date();
	const timeElapsed = Math.abs(endDate - startDate) / 60000;

	console.log(
		`⏳ Data from ${links.length} videos about ${topic_url_pair.topic} was scrapped in ${timeElapsed} minutes \n`
	);

	// Next topic
	if (list_index < urls.length - 1) {
		scrapper(++list_index);
	}
}

// Run
scrapper(0);
