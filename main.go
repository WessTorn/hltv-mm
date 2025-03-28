package main

//CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -o gohltv.so -buildvcs=false -buildmode=c-shared

import (
	"fmt"

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

	err = metamod.SetMetaCallbacks(&metamod.MetaCallbacks{
		MetaInit:   MetaInit,
		MetaQuery:  MetaQuery,
		MetaAttach: MetaAttach,
		MetaDetach: func(_ int, _ int) int {
			fmt.Println("Called MetaDetach")

			return 1
		},
	})
	if err != nil {
		panic(err)
	}
}

func main() {}

func MetaInit() {
	fmt.Println()
	fmt.Println("called MetaInit")
	fmt.Println()
}

func MetaQuery() int {
	fmt.Println()
	fmt.Println("called MetaQuery")
	fmt.Println()

	var err error

	engineFuncs, err = metamod.GetEngineFuncs()
	if err != nil {
		fmt.Println("Failed to get engine funcs:", err)
	}

	return 1
}

func MetaAttach(_ int) int {
	fmt.Println()
	fmt.Println("called MetaAttach")
	fmt.Println()

	engineFuncs.AddServerCommand("hltv", func(argc int, argv ...string) {
		if argc < 2 {
			fmt.Println("Usage: hltv <text>")
			return
		}

		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("HLTV", argv[1])
		fmt.Println("=====================================")
		fmt.Println()
	})

	return 1
}
