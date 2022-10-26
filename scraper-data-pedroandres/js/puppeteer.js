const puppeteer = require('puppeteer');

async function startBrowser() {
	let browser;

	try {
		browser = await puppeteer.launch({
			headless: true,
			args: ['--disable-setuid-sandbox'],
			ignoreHTTPSErrors: false,
		});
	} catch (err) {
		console.log('There was an error starting the browser: ', err);
	}

	return browser;
}

module.exports = startBrowser;
