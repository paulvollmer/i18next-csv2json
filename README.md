# i18next-csv2json

`i18next-csv2json` is a tool to generate json files for [i18next]() out csv file(s)

## Installation
```sh
go get github.com/paulvollmer/i18next-csv2json
```

## Usage

To generate out of one csv file the json files you can use the `-i` flag

```sh
i18next-csv2json -i path/to/input/file.csv -o path/to/output/dir
```

To generate out of a directory with multiple csv files the json files you can use the `-d` flag

```sh
i18next-csv2json -d path/to/input/dir -o path/to/output/dir
```


## License
MIT License
