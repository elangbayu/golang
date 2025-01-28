package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte("Elang Bayu Segara"))
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err == nil {
		fmt.Println(decoded)
	} else {
		fmt.Println(err.Error())
	}

	// Read CSV
	rawCSV := "elang,bayu,segara\n" + "wiwit,widowati,segara"
	reader := csv.NewReader(strings.NewReader(rawCSV))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	// Write CSV
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Elang,Bayu,Segara"})
	_ = writer.Write([]string{"Wiwit,Widowati,Segara"})
	writer.Flush()
}
