package storage

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const dbHomeVarName = "HOMEGROWN_DB_HOME"

func ReadRootPathEnv() (string, error) {
	home := os.Getenv(dbHomeVarName)
	if home == "" {
		return home, errors.New("env variable: " + dbHomeVarName + " is empty")
	} else {
		log.Printf("DB home path is set to: %s\n", home)
	}

	return home, nil
}

func SetRootPathEnv(rootPath string) error {
	return SetOsEnv(dbHomeVarName, rootPath)
}

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

func ClearRootPathEnv() error {
	return ClearOsEnv(dbHomeVarName)
}

func ClearOsEnv(envName string) error {
	homeDir, err := os.UserHomeDir()
	file, err := os.OpenFile(homeDir+"/"+".zprofile", os.O_RDWR, os.ModeType)
	if err == nil {
		return clearOsEnvAndCloseFile(envName, file)
	}
	file, err = os.OpenFile(homeDir+"/.bash_profile", os.O_RDWR, os.ModeType)
	if err == nil {
		return clearOsEnvAndCloseFile(envName, file)
	}

	return errors.New("could not file env file")
}

func clearOsEnvAndCloseFile(envName string, file *os.File) error {
	defer func() {
		err := file.Close()
		if err != nil {
			panic("unexpected err: " + err.Error())
		}
	}()
	return clearOsEnv(envName, file)
}

func clearOsEnv(envName string, file *os.File) error {
	stats, err := file.Stat()
	if err != nil {
		return err
	}
	rawFileData := make([]byte, stats.Size())
	_, err = file.Read(rawFileData)
	if err != nil {
		return err
	}
	fileData := string(rawFileData)

	index := strings.Index(fileData, "export "+envName)
	if index < 0 {
		return fmt.Errorf("file: %s does not contain env: %s", file.Name(), envName)
	}
	endIndex := index + strings.Index(fileData[index:], "\n")

	var newContent string
	if endIndex <= index {
		newContent = fileData[:index]
	} else {
		newContent = fileData[:index] + fileData[endIndex+1:]
	}

	_, err = file.WriteAt([]byte(newContent), 0)
	if err != nil {
		return err
	}

	err = file.Truncate(int64(len(newContent)))
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
	for index := strings.Index(str, "export "+envName); index >= 0; {
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
