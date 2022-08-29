# Steps
1. Open Anaconda and create an enviroment (or create it according with your preferences)
2. From that enviroment open vscode or the IDE of your preference and open the folder of this project and locate in the folder: `scraper-data-silvia`

3. From a terminal of the folder, execute this to intall the dependencies:

```
    pip install -r requirements.txt
```
5. Finally execute the scraper:

```
    python -u ".\scrapper.py"
```

*NOTE: If there is a problem with lxml, just delete ", features="lxml"" in lines 37, 43 and 51*

6. Approximate time for each topic: 3-4 min
