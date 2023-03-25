package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

const LIMIT int =  1000

func main()  {
	curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Read path errors: ", err.Error())
		return
	}

	fileDir := path.Join(curDir, "file-store")
	_ = os.Mkdir(fileDir, os.ModePerm)

	// 1. read file
	filename := "big_input_file.txt"
	inputFile, err := os.Open(path.Join(fileDir, filename))
	if err != nil {
		fmt.Println("Read inputfile error: ", err.Error())
		return
	}
	defer inputFile.Close()

	// 2. split inputFile into several pieces that every piece hold 100,000 lines
	filePieceArr := []string{}
	scanner := bufio.NewScanner(inputFile)
	piece := 1
	Outter:
		for  {
			outputFilename := "input_piece_" + strconv.Itoa(piece)
			outputFilePos := path.Join(fileDir, outputFilename)
			filePieceArr = append(filePieceArr, outputFilePos)
			outputFile, err := os.Create(outputFilePos)
			if err != nil {
				fmt.Println("Split inputFile error: ", err.Error())
				continue
			}
			defer outputFile.Close()

			for cnt := 0; cnt < LIMIT; cnt++ {
				if !scanner.Scan() {
					break Outter
				}
				_, err := outputFile.WriteString(scanner.Text() + "\n")
				if err != nil {
					fmt.Println("Split inputFile writing error: ", err.Error())
					return
				}
			}

			piece++
		}

		// 3. pass to master
		res := master.Handle(filePieceArr, fileDir)

		fmt.Println("res: ", res)
}