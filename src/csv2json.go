package main

import "fmt"
import "os"
import "encoding/csv"
import "encoding/json"
import "strconv"
import "bytes"
import "strings"
import "regexp"

func main() {
	r := csv.NewReader(os.Stdin)
	headers, _ := r.Read()
	data := map[string]float64{}

	row, err := r.Read()
	if nil != err {
		fmt.Printf("%#v\n", err)
		return
	}

	for index, element := range row {
		data[CamelCase(strings.ToLower(headers[index]))], err = strconv.ParseFloat(element, 64)
	}

	json, _ := json.Marshal(data)

	fmt.Println(string(json))
}

var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

func CamelCase(src string) string {
	byteSrc := []byte(src)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		if idx > 0 {
			chunks[idx] = bytes.Title(val)
		}
	}
	return string(bytes.Join(chunks, nil))
}
