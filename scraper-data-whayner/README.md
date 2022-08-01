### Dependencies

1. Install Ruby:

	Linux:

	```bash
	sudo apk add ruby # alphine
	sudo apt install ruby # debian
	sudo xbps-install ruby # void-linux
	```

	Windows (untested): https://rubyinstaller.org/

2. Install Chrome / Chromium.



### Setup

1. Install Ruby packages (gems):

	```bash
	bundle config set --local path 'vendor/bundle'
	bundle install
	```

2. Define search topics in `topics.txt`.

### Run
```bash
bundle exec ruby craw.rb
```

### Output files
```bash
data.json
```
