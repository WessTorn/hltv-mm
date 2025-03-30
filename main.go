package main

//CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -o gohltv.so -buildvcs=false -buildmode=c-shared

import (
	"fmt"
	"log/slog"

	metamod "github.com/et-nik/metamod-go"
)

var (
	engineFuncs *metamod.EngineFuncs
)

func init() {
	err := metamod.SetPluginInfo(&metamod.PluginInfo{
		InterfaceVersion: metamod.MetaInterfaceVersion,
		Name:             "Metamod Go HLTV",
		Version:          "v0.0.0",
		Date:             "-",
		Author:           "WessTorn",
		Url:              "https://github.com/WessTorn/hltv-mm",
		LogTag:           "MetamodGoHltv",
		Loadable:         metamod.PluginLoadTimeStartup,
		Unloadable:       metamod.PluginLoadTimeAnyTime,
	})
	if err != nil {
		panic(err)
	}

	hltv := NewHltv()

	err = metamod.SetApiCallbacks(&metamod.APICallbacks{
		GameDLLInit: func() metamod.APICallbackResult {
			err := hltv.Init()
			if err != nil {
				// TODO: logger

				return metamod.APICallbackResultHandled
			}

			err = hltv.GetPath()
			if err != nil {
				// TODO: logger
				return metamod.APICallbackResultHandled
			}

			err = hltv.CheckHltvFiles()
			if err != nil {
				// TODO: logger
				return metamod.APICallbackResultHandled
			}

			fmt.Println("HLTV initialized")

			return metamod.APICallbackResultHandled
		},
		ServerDeactivate: func() metamod.APICallbackResult {
			// err := plugin.Reset()
			// if err != nil {
			// 	slog.Error("Failed to reset plugin: ", "error", err)

			// 	return metamod.APICallbackResultHandled
			// }

			fmt.Println("Server deactivated")

			return metamod.APICallbackResultHandled
		},
	})

	if err != nil {
		panic(err)
	}

	err = metamod.SetMetaCallbacks(&metamod.MetaCallbacks{
		MetaQuery:  metaQueryFn(hltv),
		MetaDetach: metaDetachFn(hltv),
	})
	if err != nil {
		panic(err)
	}
}

func main() {}

func metaQueryFn(p *HLTV) func() int {
	return func() int {
		engineFuncs, err := metamod.GetEngineFuncs()
		if err != nil {
			slog.Error("Failed to get engine funcs: ", "error", err)

			return 0
		}
		// TODO: init logger

		gameDir := engineFuncs.GetGameDir()
		fmt.Println("HLTV game dir: ", gameDir)

		// TODO: load config

		// // Запускаем http сервер
		// go func() {
		// 	runtime.LockOSThread()

		// 	err := p.RunServer(gameDir)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// }()

		return 1
	}
}

func metaDetachFn(p *HLTV) func(now int, reason int) int {
	return func(now int, reason int) int {
		fmt.Println("HLTV meta detach")

		return 1
	}
}
