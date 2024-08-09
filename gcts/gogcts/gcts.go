// Package gogcts
// Wrote by yijian on 2024/08/03
package gogcts

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func GenerateCreateTableSqlFromFile(tableName, delimiter string, inputFile *os.File) (string, error) {
	return GenerateCreateTableSqlFromIoReader(tableName, delimiter, inputFile)
}

func GenerateCreateTableSqlFromString(tableName, delimiter string, inputString string) (string, error) {
	return GenerateCreateTableSqlFromIoReader(tableName, delimiter, strings.NewReader(inputString))
}

func GenerateCreateTableSqlFromIoReader(tableName, delimiter string, inputIoReader io.Reader) (string, error) {
	var result strings.Builder

	// delimiter 不能为空格、TAB符、单引号、双引号和反引号
	if delimiter == " " || delimiter == "\t" || delimiter == "'" || delimiter == "\"" || delimiter == "`" {
		return "", fmt.Errorf("delimiter must not be a space, tab, single quote, double quote, or backquote")
	}
	if len(delimiter) != 1 {
		return "", fmt.Errorf("delimiter must be a single character")
	}

	result.WriteString("DROP TABLE IF EXISTS `" + tableName + "`;\n")
	result.WriteString("CREATE TABLE `" + tableName + "` (\n")

	scanner := bufio.NewScanner(inputIoReader)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, delimiter)
		if len(columns) == 3 {
			fName := strings.TrimSpace(columns[0])
			fType := strings.TrimSpace(columns[1])
			fComment := strings.TrimSpace(columns[2])
			result.WriteString("  `" + fName + "` " + fType + " COMMENT '" + fComment + "',\n")
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	result.WriteString(");\n")
	return result.String(), nil
}
