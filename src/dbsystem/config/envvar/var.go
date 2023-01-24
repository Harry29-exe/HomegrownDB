package envvar

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func SetOsEnv(envName string, envValue string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(homeDir+"/"+".zprofile", os.O_RDWR, os.ModeType)
	if err == nil {
		return writeOsEnv(file, envName, envValue)
	}
	file, err = os.OpenFile(homeDir+"/.bash_profile", os.O_RDWR, os.ModeType)
	if err == nil {
		return writeOsEnv(file, envName, envValue)
	}

	return err
}

func writeOsEnv(file *os.File, envName, envValue string) error {
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	fileData := make([]byte, stat.Size())
	_, err = file.Read(fileData)
	if err != nil {
		return err
	}

	str := string(fileData)
	var outputData string
	if start, end := envExportExist(str, envName); start >= 0 {
		outputData = fmt.Sprintf("%s%s%s",
			str[:start],
			fmt.Sprintf("export %s=\"%s\"", envName, escapeChars(envValue)),
			str[end:],
		)
	} else {
		outputData = str + fmt.Sprintf("\nexport %s=\"%s\"", envName, envValue)
	}
	_, err = file.WriteAt([]byte(outputData), 0)
	return err
}

func envExportExist(envFile string, envName string) (startIndex, endIndex int) {
	str := envFile
	offset := 0
	for index := strings.Index(str, envName); index >= 0; {
		startIndex = strings.LastIndex(str[:index], "\n") + 1
		endIndex = strings.Index(str[index:], "\n")
		if endIndex < 0 {
			endIndex = len(str)
		}
		exportLine := removeSpace(str[startIndex:endIndex])
		if index = strings.Index(exportLine, fmt.Sprintf("export %s=\"", envName)); index == 0 {
			return startIndex + offset, endIndex + offset
		} else {
			offset += len(str[:endIndex])
			str = str[endIndex:]
		}
	}
	return -1, -1
}

func removeSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}

func escapeChars(str string) string {
	str = strings.ReplaceAll(str, "\n", "\\n")
	str = strings.ReplaceAll(str, "\"", "\\\"")
	return str
}
