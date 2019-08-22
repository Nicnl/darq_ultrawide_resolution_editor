package main

import (
	"encoding/json"
	"fmt"
	"nrbf/ll_nrbf"
	"os"
)

func main() {
	fmt.Print("1. Opening save.drq...")
	f, err := os.Open("save.drq")
	if err != nil {
		fmt.Println("FAIL")
		panic(err)
	}
	fmt.Println(" OK")

	fmt.Print("2. Creating nrbf decoder...")
	r := ll_nrbf.NewDecoder(f)
	fmt.Println(" OK")

	fmt.Print("3. Parsing nrbf header...")
	h, err := r.ParseHeader()
	if err != nil {
		fmt.Println("FAIL")
		panic(err)
	}
	fmt.Println(" OK")

	data, err := json.MarshalIndent(&h, "    ", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Print("    ")
	fmt.Println(string(data))

	fmt.Print("4. Read records...")
	for {
		rec, err := r.NextRecord()
		if err != nil {
			fmt.Println("FAIL")
			panic(err)
		}
		fmt.Println(" OK")

		data, err = json.MarshalIndent(&rec, "    ", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Print("    ")
		fmt.Println(string(data))
	}
}
