package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	var urls []string
	file, err := os.Open("urls.txt")
	handleErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	err = scanner.Err()
	handleErr(err)

	for index, URL := range urls {
		go download(URL, index)
		time.Sleep(500 * time.Millisecond)
	}

	time.Sleep(3600 * time.Second)
}

func download(URL string, index int) {

	fmt.Println(URL)
	file, err := os.Create("EPISODE_" + fmt.Sprint(index) + "_.mp3")
	handleErr(err)
	defer file.Close()
	get, err := http.Get(URL)
	handleErr(err)
	_, err = io.Copy(file, get.Body)
	handleErr(err)

}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Unable to get file")
	}
}
