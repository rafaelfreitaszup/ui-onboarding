package main

import (
	"time"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	windowRect := sciter.NewRect(250, 500, 635, 625)

	presentation, windowsGenerateionError := window.New(sciter.SW_TITLEBAR|sciter.SW_TOOL|sciter.SW_CONTROLS, windowRect)

	if windowsGenerateionError != nil {
		print("Failed to generate sciter window ", windowsGenerateionError.Error())
	}

	uiLoadingError := presentation.LoadFile("./ui/presentation/index.html")

	if uiLoadingError != nil {
		print("Failed to load ui file ", uiLoadingError.Error())
	}

	showSplashScreen(presentation)

	presentation.DefineFunction("load_tool_selection", func(...*sciter.Value) *sciter.Value {
		showToolSelection(presentation)
		return nil
	})

	rootEl, _ := presentation.GetRootElement()

	buttonEl, _ := rootEl.SelectById("continue")

	buttonEl.OnClick(func() {
		presentation.Call("close", sciter.NewValue("ok"))
	})

	presentation.Run()
}

func showSplashScreen(mainScreen *window.Window) {
	rect := sciter.NewRect(250, 500, 635, 625)

	splashScreen, windowsGenerateionError := window.New(sciter.SW_TOOL, rect)

	if windowsGenerateionError != nil {
		print("Failed to generate sciter window ", windowsGenerateionError.Error())
	}

	uiLoadingError := splashScreen.LoadFile("./ui/splash-screen/index.html")

	if uiLoadingError != nil {
		print("Failed to load ui file ", uiLoadingError.Error())
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

func showToolSelection(mainScreen *window.Window) {
	rect := sciter.NewRect(250, 500, 635, 625)

	toolSelection, windowsGenerateionError := window.New(sciter.SW_TITLEBAR|sciter.SW_TOOL|sciter.SW_CONTROLS, rect)

	if windowsGenerateionError != nil {
		print("Failed to generate sciter window ", windowsGenerateionError.Error())
	}

	uiLoadingError := toolSelection.LoadFile("./ui/tool-selection/index.html")

	if uiLoadingError != nil {
		print("Failed to load ui file ", uiLoadingError.Error())
	}

	// toolSelection.DefineFunction("load_main_screen", func(...*sciter.Value) *sciter.Value {
	// 	mainScreen.Show()
	// 	return nil
	// })

	toolSelection.Show()
}
