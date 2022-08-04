## Velocidad

Tested with 15 topics, 140 videos each:

- 2100 videos in 29 minutes
- ~ 72 videos per minute

## Dependencies

1. Install Ruby:

	Linux:

	```bash
	sudo apk add ruby # alphine
	sudo apt install ruby # debian
	sudo xbps-install ruby # void-linux
	```

	Windows (untested): https://rubyinstaller.org/

2. Install Chrome / Chromium.



## Setup

- Clone 
- Install Ruby packages (gems):

	```bash
	bundle config set --local path 'vendor/bundle'
	bundle install
	```

- Define search topics in `topics.txt`.

## Run

This will start the crawling process:

```bash
bundle exec ruby crawl.rb
```

The output data will be saved in `data.json`.

## Formatting
You can execute this formating tool to remove some unwanted characters from the data:

```bash
bundle exec ruby formater.rb data.json
```

It'll create a new formated file called `data.json.formated`.

