package chromeservice

import (
	"fmt"

	"github.com/fedesog/webdriver"
)

func GetNewChromeDriverService() *webdriver.ChromeDriver {
	defer fmt.Println("Create New Chrome Service.")
	return webdriver.NewChromeDriver("./chromedriver.exe")
}
