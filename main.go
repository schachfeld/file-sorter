package main

import (
	"fmt"
	"log"
	"os"
)

func sortFiles(dir string, outputDir string, parentDirName string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	fails := 0
	success := 0

	for _, e := range entries {
		fi, err := os.Stat(dir + "/" + e.Name())
		if err != nil {
			fmt.Println(err)
			return
		}

		if fi.IsDir() {
			sortFiles(dir+"/"+e.Name(), outputDir, fi.Name())
			continue
		}

		date := fi.ModTime().Format("2006-01-02")
		originalFilename := e.Name()

		os.MkdirAll(outputDir+"/"+date, os.ModePerm)

		// 2009-12-31--23-59-59
		err = os.Rename(dir+"/"+fi.Name(), outputDir+"/"+date+"/"+fi.ModTime().Local().Format("2006-01-02--15-04-05")+"_"+parentDirName+"_"+originalFilename)
		if err != nil {
			fmt.Println(err)
			fails++
			continue
		}
		success++
	}
	fmt.Println("Success:", success, "Fails:", fails)
}

func main() {
	startDir := "<your input dir>"
	outputDir := "<your output dir>"
	sortFiles(startDir, outputDir, "")
}
