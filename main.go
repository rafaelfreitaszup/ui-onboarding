package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"gopkg.in/yaml.v2"
)

type Setup map[string][]string

func main() {
	windowRect := sciter.NewRect(250, 500, 635, 625)

	presentation, windowsGenerateionError := window.New(sciter.SW_TITLEBAR|sciter.SW_TOOL|sciter.SW_CONTROLS, windowRect)

	if windowsGenerateionError != nil {
		log.Println("Failed to generate sciter window Presentation", windowsGenerateionError.Error())
	}

	uiLoadingError := presentation.LoadFile("./ui/presentation/index.html")

	if uiLoadingError != nil {
		log.Println("Failed to load ui file ./ui/presentation/index.html ", uiLoadingError.Error())
	}

	showSplashScreen(presentation)

	presentation.DefineFunction("load_tool_selection", func(args ...*sciter.Value) *sciter.Value {
		if args[0].IsString() == true {
			showToolSelection(args[0])
		}

		return nil
	})

	presentation.Run()
}

func showSplashScreen(mainScreen *window.Window) {
	rect := sciter.NewRect(250, 500, 635, 662)

	splashScreen, windowsGenerateionError := window.New(sciter.SW_TOOL, rect)

	if windowsGenerateionError != nil {
		log.Println("Failed to generate sciter window SplashScreen", windowsGenerateionError.Error())
	}

	uiLoadingError := splashScreen.LoadFile("./ui/splash-screen/index.html")

	if uiLoadingError != nil {
		log.Println("Failed to load ui file ./ui/splash-screen/index.html ", uiLoadingError.Error())
	}

	splashScreen.DefineFunction("load_main_screen", func(...*sciter.Value) *sciter.Value {
		mainScreen.Show()
		return nil
	})

	splashScreen.Show()

	go func() {
		time.Sleep(6 * time.Second)
		splashScreen.Call("close_splash", sciter.NewValue("ok"))
	}()
}

func showToolSelection(key *sciter.Value) {
	rect := sciter.NewRect(250, 500, 635, 625)

	toolSelection, windowsGenerateionError := window.New(sciter.SW_TITLEBAR|sciter.SW_TOOL|sciter.SW_CONTROLS, rect)

	if windowsGenerateionError != nil {
		log.Println("Failed to generate sciter window ToolSelection", windowsGenerateionError.Error())
	}

	uiLoadingError := toolSelection.LoadFile("./ui/tool-selection/index.html")

	if uiLoadingError != nil {
		log.Println("Failed to load ui file ./ui/tool-selection/index.html ", uiLoadingError.Error())
	}

	toolSelection.Call("selectElement", sciter.NewValue(key))
	toolSelection.Call("addElementsConfig", sciter.NewValue(string(parseJSON(key))))

	toolSelection.DefineFunction("changeConfig", func(args ...*sciter.Value) *sciter.Value {
		if args[0].IsString() == true {
			toolSelection.Call("addElementsConfig", sciter.NewValue(string(parseJSON(args[0]))))
		}

		return nil
	})

	toolSelection.DefineFunction("load_status", func(...*sciter.Value) *sciter.Value {
		showStatus()

		return nil
	})

	toolSelection.Show()
}

func showStatus() {
	rect := sciter.NewRect(250, 500, 635, 625)

	status, windowsGenerateionError := window.New(sciter.SW_TITLEBAR|sciter.SW_TOOL|sciter.SW_CONTROLS, rect)

	if windowsGenerateionError != nil {
		log.Println("Failed to generate sciter window Status", windowsGenerateionError.Error())
	}

	uiLoadingError := status.LoadFile("./ui/status/index.html")

	if uiLoadingError != nil {
		log.Println("Failed to load ui file ./ui/status/index.html ", uiLoadingError.Error())
	}

	status.Show()
}

func parseJSON(key *sciter.Value) []byte {
	parsedFile := getConfigs(key)

	json, _ := json.Marshal(parsedFile)

	return json
}

func getConfigs(key *sciter.Value) Setup {
	yamlFile, err := os.Open(fmt.Sprintf("./config/%s.yml", key))

	if err != nil {
		log.Println(err)
	}

	defer yamlFile.Close()

	parsedFile := Setup{}

	byteValue, _ := ioutil.ReadAll(yamlFile)

	err = yaml.Unmarshal(byteValue, &parsedFile)

	if err != nil {
		log.Println(err)
	}

	return parsedFile
}
