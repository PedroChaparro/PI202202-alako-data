## Velocidad

Probado con 15 temas, 140 registros cada uno:

- 2100 registros en 29 minutos
- ~ 72 registros por minuto

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

1. Install Ruby packages (gems):

	```bash
	bundle config set --local path 'vendor/bundle'
	bundle install
	```

2. Define search topics in `topics.txt`.

## Run
```bash
bundle exec ruby crawl.rb
```

## Output files
```bash
data.json
```
