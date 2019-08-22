package main

import (
	"encoding/json"
	"fmt"
	"io"
	"nrbf/ll_nrbf"
	"os"
)

func readRecords() []ll_nrbf.Record {
	fmt.Print("1. Opening save.drq...")
	f, err := os.Open("save.drq")
	if err != nil {
		fmt.Println("FAIL")
		panic(err)
	}
	defer f.Close()
	fmt.Println(" OK")

	fmt.Print("2. Creating nrbf decoder...")
	r := ll_nrbf.NewDecoder(f)
	fmt.Println(" OK")

	fmt.Print("3. Read records...")
	records := make([]ll_nrbf.Record, 0)
	err = nil
	for {
		rec, err := r.NextRecord()
		if err == io.EOF {
			return records
		}
		if err != nil {
			fmt.Println("FAIL")
			panic(err)
		}
		fmt.Println(" OK")

		data, err := json.MarshalIndent(&rec, "    ", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Print("    ")
		fmt.Println(string(data))

		records = append(records, rec)
	}
}

func writeFile(records []ll_nrbf.Record) {
	fmt.Print("4. Opening save_custom.drq...")
	f, err := os.Create("save_custom.drq")
	if err != nil {
		fmt.Println("FAIL")
		panic(err)
	}
	defer f.Close()
	fmt.Println(" OK")

	e := ll_nrbf.NewEncoder(f)
	for _, record := range records {
		err = e.WriteRecord(record)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	records := readRecords()
	writeFile(records)
}
