package main

import (
	"net/http"
	"os"
	"io"
)

func wget(url string) error{
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	file, err := os.OpenFile("index.html", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := file.Truncate(0); err != nil {
		return err
	}
	if _, err := file.Write(body); err != nil {
		return err
	}
	return nil
}

func main() {
	wget("https://www.daniweb.com/programming/computer-science/code/495192/get-the-content-of-a-web-page-golang")
}