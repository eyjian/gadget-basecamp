// Package main
// Wrote by yijian on 2024/08/03
package main

import (
	"fmt"
	"os"
)
import (
	"github.com/eyjian/gadget-basecamp/gcts/gogcts"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s <table_name> <delimiter> <input_file>\n", os.Args[0])
		fmt.Printf("Example: %s 't_user' ',' 'input_file.txt'\n", os.Args[0])
		return
	}

	tableName := os.Args[1]
	delimiter := os.Args[2]
	inputFilePath := os.Args[3]

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer inputFile.Close()

	createTableSQL, err := gogcts.GenerateCreateTableSqlFromFile(tableName, delimiter, inputFile)
	if err != nil {
		fmt.Printf("Error generating create table SQL: %v\n", err)
		return
	}
	fmt.Println(createTableSQL)
}
