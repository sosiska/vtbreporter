package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		seleniumPath    = "/home/kirill/.cache/yay/selenium-server-standalone/selenium-server-standalone-3.141.59.jar"
		geckoDriverPath = "/usr/bin/geckodriver"
		port            = 8080
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(io.Discard),           // Output debug information to STDERR.
	}
	//selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	fmt.Println("Заходим на сайт")

	// Заходим на сайт
	if err := wd.Get("https://lk.broker.vtb.ru/login"); err != nil {
		panic(err)
	}

	fmt.Println("Ищем форму логина")

	// Ищем форму логина
	loginForm, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/main/div/div/div[1]/div/div[1]/section/div/form/div[1]/div/div/input")
	if err != nil {
		panic(err)
	}
	// Удаляем все лишнее
	if err := loginForm.Clear(); err != nil {
		panic(err)
	}

	fmt.Println("Вводим логин")

	// Вводим логин
	err = loginForm.SendKeys(``)
	if err != nil {
		panic(err)
	}

	// Ищем форму пароля
	passForm, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/main/div/div/div[1]/div/div[1]/section/div/form/div[2]/div/div/div/input")
	if err != nil {
		panic(err)
	}
	// Удаляем все лишнее
	if err := passForm.Clear(); err != nil {
		panic(err)
	}

	fmt.Println("Вводим пароль")

	// Вводим пароль
	err = passForm.SendKeys(``)
	if err != nil {
		panic(err)
	}

	fmt.Println("Нажимаем кнопку зайти")

	// Ищем кнопку зайти
	loginBtn, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/main/div/div/div[1]/div/div[1]/section/div/form/button")
	if err != nil {
		panic(err)
	}

	// тыкаем ее
	if err := loginBtn.Click(); err != nil {
		panic(err)
	}

	fmt.Println("Ждем загрузки сайта")

	time.Sleep(3 * time.Second)

	fmt.Println("Напрямую переходим в отчеты")

	err = wd.Get("https://lk.broker.vtb.ru/other/18")
	if err != nil {
		panic(err)
	}

	fmt.Println("Ждем загрузки сайта")

	time.Sleep(3 * time.Second)

	fmt.Println("Выбираем основной счет")

	list, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/main/div/div/div[2]/div/div[1]/div/div[3]/section/form/div[1]/div/div[1]/button/div")
	if err != nil {
		panic(err)
	}

	if err := list.Click(); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	mainSogl, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/main/div/div/div[2]/div/div[1]/div/div[3]/section/form/div[1]/div/div[2]/ul/div/div/div[1]/div/div/div[1]/div/div/div/div/span[1]")
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	if err := mainSogl.Click(); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("Выбираем тип периода")

	periodType, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/main/div/div/div[2]/div/div[1]/div/div[3]/section/form/div[2]/div/div[1]/div/div[1]/button/div/div/div")
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	if err := periodType.Click(); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	period, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/main/div/div/div[2]/div/div[1]/div/div[3]/section/form/div[2]/div/div[1]/div/div[2]/ul/div/div/div[1]/div/div/div[4]/div/div/div")
	if err != nil {
		panic(err)
	}

	if err := period.Click(); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("Выставляем сам период")

	getBtn, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/main/div/div/div[2]/div/div[1]/div/div[3]/section/form/button")
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	if err := getBtn.Click(); err != nil {
		panic(err)
	}

	bb, err := wd.Screenshot()
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("/home/kirill/123.jpg", bb, 0644)
	if err != nil {
		panic(err)
	}

	os.Exit(0)

	//// Wait for the program to finish running and get the output.
	//outputDiv, err := wd.FindElement(selenium.ByXPATH, "/html/body/main/div/div[3]")
	//if err != nil {
	//	panic(err)
	//}
	//
	//var output string
	//for {
	//	output, err = outputDiv.Text()
	//	if err != nil {
	//		panic(err)
	//	}
	//	if strings.Contains(output, "WebDriver") {
	//		break
	//	}
	//	time.Sleep(time.Millisecond * 100)
	//}
	//
	//fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))

}
