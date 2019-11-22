package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/Jeffail/gabs/v2"
)

func usage() {
	fmt.Println("Usage: i18next-csv2json [flag]")
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
}

func main() {
	flagInputFile := flag.String("i", "", "path to the csv input file")
	flagInputDir := flag.String("d", "", "path to the csv input directory")
	flagOutput := flag.String("o", "./", "path to the output directory")
	flag.Usage = usage
	flag.Parse()

	if *flagInputFile == "" && *flagInputDir == "" {
		fmt.Println("Missing -i flag")
		os.Exit(1)
	}

	if *flagInputFile != "" {
		Generate(*flagInputFile, *flagOutput)
	} else if *flagInputDir != "" {
		files, err := ioutil.ReadDir(*flagInputDir)
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}
		for i := 0; i < len(files); i++ {
			if !files[i].IsDir() {
				Generate(path.Join(*flagInputDir, files[i].Name()), *flagOutput)
			}
		}
	}
}

func GenerateFromBytes(source []byte) ([]string, [][]byte, error) {
	langs := make([]string, 0)
	jsonObjs := make([]*gabs.Container, 0)

	// read and process the csv data
	r := csv.NewReader(bytes.NewReader(source))
	r.Comma = ','
	r.Comment = '#'
	counter := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return langs, [][]byte{}, err
		}
		if counter == 0 {
			langs = record[1:]
			for i := 0; i < len(langs); i++ {
				jsonObjs = append(jsonObjs, gabs.New())
			}
		}
		if counter > 0 {
			for i := 1; i < len(record); i++ {
				jsonObjs[i-1].SetP(record[i], record[0])
			}
		}
		counter++
	}

	data := make([][]byte, len(langs))
	for i := 0; i < len(langs); i++ {
		data[i] = jsonObjs[i].Bytes()
	}
	return langs, data, nil
}

func Generate(input, output string) {
	csvRaw, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println("Read File Error:", err)
		os.Exit(1)
	}

	langs, data, err := GenerateFromBytes(csvRaw)
	if err != nil {
		fmt.Print("CSV Error:", err)
		os.Exit(1)
	}

	filename := strings.Replace(path.Base(input), ".csv", ".json", 1)
	for i := 0; i < len(langs); i++ {
		// make output directory
		err := os.MkdirAll(path.Join(output, langs[i]), 0755)
		if err != nil {
			fmt.Print("Create output directory error:", err)
			os.Exit(1)
		}
		// write the json file
		err = ioutil.WriteFile(path.Join(output, langs[i], filename), data[i], 0755)
		if err != nil {
			fmt.Print("Write JSON file error:", err)
			os.Exit(1)
		}
	}
}
