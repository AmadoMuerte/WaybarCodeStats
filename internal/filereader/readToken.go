package filereader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadToken(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "api_key = ") {
			return strings.TrimSpace(line[10:]), nil
		}
	}

	return "", fmt.Errorf("api_key not found in file")
}
