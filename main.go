package main

import (
	"log"
	"time"

	"github.com/LoveAndCode/golenium/chromeservice"

	"github.com/fedesog/webdriver"
)

func main() {
	service := chromeservice.GetNewChromeDriverService()
	err := service.Start()

	if err != nil {
		log.Println(err)
	}
	desired := webdriver.Capabilities{"platform": "Window"}
	required := webdriver.Capabilities{}

	session, err := service.NewSession(desired, required)

	if err != nil {
		log.Println(err)
	}

	err = session.Url("https://naver.com")

	if err != nil {
		log.Println(err)
	}

	time.Sleep(10 * time.Second)
	session.Delete()
	service.Stop()
}
