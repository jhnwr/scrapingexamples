import requests
from bs4 import BeautifulSoup
import csv


def get_fnumbers():
    fnums_list = []
    url = "https://federbridge.it/SocSp/ChkFS.asp?FS=F"
    response = requests.get(url)
    soup = BeautifulSoup(response.text, 'html.parser')
    fnumbers = soup.select('span.COLviolaChiaro')
    for fnum in fnumbers:
        if fnum.text[0] == "F":
            fnums_list.append(fnum.text)
    return fnums_list


def detail_page(fnumber):
    response = requests.get(f"https://www.federbridge.it/regioni/dettAss.asp?codice={fnumber}")
    soup = BeautifulSoup(response.text, 'html.parser')
    name = soup.select_one("td.FNTbase12").text
    email = soup.select_one("td.ALLbaseR a[href^='mailto']").text
    return name, email


def save_to_csv(output):
    with open("results.csv", "w") as f:
        csv_writer = csv.writer(f)
        for item in output:
            csv_writer.writerow(item)


def main():
    results = []
    numbers = get_fnumbers()
    for num in numbers[:10]:
        results.append(detail_page(num))
        print(num)
    save_to_csv(results)


if __name__ == '__main__':
    main()
