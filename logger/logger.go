package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

type logWriter struct {
	logPath string
}

func (w logWriter) Write(b []byte) (int, error) {
	file, err := os.OpenFile(w.logPath+"hltv"+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	wrt := io.MultiWriter(os.Stdout, file)
	defer file.Close()
	b = bytes.Replace(b, []byte(".go:"), []byte{':'}, -1)
	return wrt.Write(b)
}

func Init(basePath string) error {
	wrt := new(logWriter)
	wrt.logPath = basePath
	err := os.MkdirAll(wrt.logPath, os.ModePerm)
	if err != nil {
		return err
	}
	Info = log.New(wrt, "[HLTV] INFO: ", log.Lshortfile)
	Warning = log.New(wrt, "[HLTV] WARNING: ", log.Lshortfile)
	Error = log.New(wrt, "[HLTV] ERROR: ", log.Lshortfile)
	return nil
}
