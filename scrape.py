import requests
from bs4 import BeautifulSoup
import typer

app = typer.Typer()

@app.command()
def scrape(url:str):
    text = ""
    try:
        response = requests.get(url)
        if response.status_code == 200:
            soup = BeautifulSoup(response.text, 'html.parser')
            text = soup.get_text()
        else: print(f"Failed to retrieve the page. Status code: {response.status_code}")
    except: pass
    typer.echo(text)


if __name__ == "__main__":
    app()



# python scrape.py <url>