package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"nrbf/helpers"
	"nrbf/ll_nrbf"
	"nrbf/quick_tui"
	"os"
	"path/filepath"
	"time"
)

func readRecords(readPath string) ([]ll_nrbf.Record, error) {
	fmt.Print("  1. Opening save file...")
	f, err := os.Open(readPath)
	if err != nil {
		fmt.Println("FAIL")
		return nil, err
	}
	defer f.Close()
	fmt.Println(" OK")

	fmt.Print("  2. Creating nrbf decoder...")
	r := ll_nrbf.NewDecoder(f)
	fmt.Println(" OK")

	fmt.Println("  3. Reading nrbf records...")
	records := make([]ll_nrbf.Record, 0)
	err = nil
	for {
		rec, err := r.NextRecord()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			fmt.Println("FAIL")
			return nil, err
		}

		data, err := json.MarshalIndent(&rec, "    ", "  ")
		if err != nil {
			return nil, err
		}
		fmt.Print("    ")
		fmt.Println(string(data))
		fmt.Println()

		records = append(records, rec)
	}
}

func writeFile(dest string, records []ll_nrbf.Record) error {
	fmt.Print("  1. Opening the save file...")
	f, err := os.Create(dest)
	if err != nil {
		fmt.Println("FAIL")
		return err
	}
	defer f.Close()
	fmt.Println(" OK")

	fmt.Print("  2. Creating nrbf encoder...")
	e := ll_nrbf.NewEncoder(f)
	fmt.Println(" OK")

	fmt.Println("  3. Writing nrbf records...")
	for _, record := range records {
		fmt.Println("   => Writing record", record.Record)
		err = e.WriteRecord(record)
		if err != nil {
			return err
		}
	}
	fmt.Println("OK")

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	quick_tui.SetReader(reader)

	darqSavePath := ""

	// 1] Check if there is a provided file (drag drop on the icon)
	if len(os.Args) == 2 {
		if _, err := os.Stat(os.Args[1]); err == nil {
			fmt.Println("File path provided:", os.Args[1])
			darqSavePath = os.Args[1]
		} else {
			fmt.Println("Wtf, this is not supposed to happen")
			time.Sleep(3600 * time.Second)
			return
		}
	}

	// 2] If not, write a welcome message
	if len(os.Args) != 2 {
		fmt.Println("Hello there!")
		fmt.Println()
		fmt.Println("It seems you haven't provided a DARQ save file.")
		fmt.Println()
		fmt.Println("You can do this by dragging & dropping your 'save.drq' file")
		fmt.Println(" directly over the icon of this program.")
		fmt.Println("(Just like you do when you put a file in a folder, except you drop on the program icon)")
		fmt.Println()
		fmt.Println("DARQ's save file is usually stored at the following location:")
		fmt.Println("  %USERPROFILE%\\AppData\\LocalLow\\Unfold Games\\DARQ\\save.drq")

		// If a darq save file exists, ask the user if he wants us to open it instead
		expectedPath := filepath.Join(
			os.Getenv("USERPROFILE"),
			"AppData",
			"LocalLow",
			"Unfold Games",
			"DARQ",
			"save.drq",
		)

		if _, err := os.Stat(expectedPath); err == nil {
			fmt.Println()
			fmt.Println("Do you want me to go and fetch it for you instead?")

			if !quick_tui.InputYesOrNo() {
				quick_tui.OkBye()
				return
			}

			darqSavePath = expectedPath
		} else {
			time.Sleep(3600 * time.Second)
			return
		}
	}

	// Try and decode the save file
	time.Sleep(350 * time.Millisecond)
	records, err := readRecords(darqSavePath)
	time.Sleep(350 * time.Millisecond)
	if err != nil {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("I wasn't able to read the provided save file.")
		fmt.Println("Are you sure this is a valid DARQ save file?")
		fmt.Println("Please check and try again.")

		time.Sleep(3600 * time.Second)
		return
	}

	// The save file was successfully parsed, let's search for the resolution
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("The 'save.drq' save file was successfully loaded.")
	fmt.Println("Searching for the resolution...")
	var (
		widthIndex  = -1
		heightIndex = -1
	)
	cmt := records[2].Record.(ll_nrbf.ClassWithMembersAndTypes) // TODO: fix risky cast
	for i, key := range cmt.ClassInfo.MemberNames {
		fmt.Printf(" %d", i)
		if key == "screenWidth" {
			widthIndex = i
			fmt.Print(" [found width]")
			time.Sleep(175 * time.Millisecond)
		} else if key == "screenHeight" {
			heightIndex = i
			fmt.Print(" [found height]")
			time.Sleep(175 * time.Millisecond)
		}

		if widthIndex != -1 && heightIndex != -1 {
			break
		}
	}
	fmt.Println()
	if widthIndex == -1 || heightIndex == -1 {
		fmt.Println("The resolution could not be found.")
		fmt.Println("Are you sure this is a valid DARQ save file?")
		fmt.Println("Please check and try again.")
		time.Sleep(3600 * time.Second)
		return
	}

	time.Sleep(150 * time.Millisecond)
	fmt.Println()
	fmt.Printf("FOUND! Current resolution:  %dx%d\n", cmt.Values[widthIndex], cmt.Values[heightIndex])
	fmt.Println()
	time.Sleep(150 * time.Millisecond)

	// Ask the user to enter the new resolution (width first, then height)
	newWidth := quick_tui.InputNumberResolution("width")
	fmt.Println(" => New width successfully confirmed:", newWidth)
	fmt.Println()

	newHeight := quick_tui.InputNumberResolution("height")
	fmt.Println(" => New height successfully confirmed:", newHeight)
	fmt.Println()

	// Replace the resolution in the config file IN MEMORY ONLY FOR NOW
	cmt.Values[widthIndex] = int32(newWidth)
	cmt.Values[heightIndex] = int32(newHeight)
	entry := records[2]
	entry.Record = cmt
	records[2] = entry

	// Everything is ready, ask the user if he wants to overwrite the file
	fmt.Println()
	fmt.Println()
	fmt.Println("Everything is ready. Do you want to overwrite the save file?")
	fmt.Println("(A backup of the previous one will be made anyway, no matter what)")
	fmt.Println()

	if !quick_tui.InputYesOrNo() {
		quick_tui.OkBye()
		return
	}

	// Doing a backup of the previous save file
	fmt.Println()
	fmt.Println()
	time.Sleep(120 * time.Millisecond)
	copyPath := darqSavePath + "__backup__" + time.Now().Format("2006-01-02_15-04-05")
	if helpers.CopyFile(darqSavePath, copyPath) != nil {
		fmt.Println("Error when doing the backup")
		fmt.Println("Please check is a program is locking the file or something.")
		time.Sleep(3600 * time.Second)
		return
	}
	fmt.Println(" => Backup created at the following location:")
	fmt.Println("    ", copyPath)
	time.Sleep(333 * time.Millisecond)

	// Actually overwriting the save
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Overwriting the save file...")
	time.Sleep(150 * time.Millisecond)
	err = writeFile(darqSavePath, records)
	if err != nil {
		fmt.Println("Error when overwriting the save file.")
		fmt.Println("Please check is a program is locking the file or something.")
		time.Sleep(3600 * time.Second)
		return
	}
	time.Sleep(320 * time.Millisecond)

	// Done !
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Done!")
	fmt.Printf("  Your new resolution is:  %dx%d\n", cmt.Values[widthIndex], cmt.Values[heightIndex])
	time.Sleep(3600 * time.Second)
}
