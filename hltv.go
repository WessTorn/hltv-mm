package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type HLTV struct {
	path string

	server *http.Server
}

func NewHltv() *HLTV {
	return &HLTV{}
}

func (hltv *HLTV) Init() error {
	return nil
}

func (hltv *HLTV) GetPath() error {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return err
	}

	hltv.path = path

	return nil
}

func (hltv *HLTV) CheckHltvFiles() error {
	var requiredFiles = []string{
		"hltv",
		"filesystem_stdio.so",
		"proxy.so",
		"libsteam_api.so",
		"core.so",
		"hltv.cfg",
		"steamclient.so",
		"libsteam.so",
	}

	ldLibraryPath := os.Getenv("LD_LIBRARY_PATH")
	fmt.Println("LD_LIBRARY_PATH:", ldLibraryPath)

	searchPaths := []string{hltv.path}

	paths := strings.Split(ldLibraryPath, ":")
	for _, path := range paths {
		if path != "." && path != "" {
			searchPaths = append(searchPaths, path)
		}
	}

	fmt.Println("Checking files in directories:")
	for _, p := range searchPaths {
		fmt.Println(" -", p)
	}

	missingFiles := []string{}

	for _, file := range requiredFiles {
		found := false
		for _, path := range searchPaths {
			fullPath := filepath.Join(path, file)
			if _, err := os.Stat(fullPath); err == nil {
				found = true
				break
			}
		}
		if !found {
			missingFiles = append(missingFiles, file)
		}
	}

	if len(missingFiles) > 0 {
		fmt.Println("Missing required files:")
		for _, file := range missingFiles {
			fmt.Println(" -", file)
		}
		return fmt.Errorf("some required files are missing")
	}

	fmt.Println("All required files are present.")
	return nil
}
