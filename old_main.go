package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/mvdan/xurls"
)

func buscaUrls(body string) []string {
	return xurls.Strict.FindAllString(body, -1)
}

func getBody(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		println("deu erro")
		return ""
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(body)
	return stringBody
}

func main() {
	url := "http://g1.globo.com/"
	body := getBody(url)
	urls := buscaUrls(body)

	var stringUrls string
	for i := 0; i < len(urls); i++ {
		stringUrls = stringUrls + "\n" + urls[i]
	}

	arquivo := "urls.txt"
	f, err := Open(arquivo)
	bytesWritten, err := f.WriteString(stringUrls)
	fmt.Printf("%d bytes\n", bytesWritten)
	if err != nil {
		fmt.Println("Vish, deu problema!")
	}
	defer f.Close()
	urlsRetornadas := GetUrlsFromFile(arquivo)
	fmt.Println(urlsRetornadas)

	for j := 0; j < len(urlsRetornadas); j++ {
		if strings.HasSuffix(urlsRetornadas[j], "jpeg") {

			nomeArquivo := strconv.Itoa(j) + ".jpeg"

			f, err = Open(nomeArquivo)
			if err != nil {
				fmt.Println("Deu ruim!")
			} else {
				bytesWritten, err = f.WriteString(getBody(urlsRetornadas[j]))
				fmt.Printf("%d bytes\n", bytesWritten)
				if err != nil {
					fmt.Println("Vish, deu problema!")
					fmt.Println(err)
				}
				defer f.Close()
			}
		}
	}

}
