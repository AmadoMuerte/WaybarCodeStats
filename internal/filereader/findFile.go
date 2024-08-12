package filereader

import (
	"path/filepath"
)

func FindFile(filePATH string) (string, error) {
	absPath, err := filepath.Abs(filePATH)
	if err != nil {
		return "", err
	}

	return absPath, nil
}
