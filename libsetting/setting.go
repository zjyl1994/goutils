package libsetting

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	applicationName  string
	loadedConfigPath string
	lastError        error
)

func Init(appName string, defaultConfig []byte) error {
	applicationName = strings.ToUpper(appName)
	if defaultConfig != nil {
		if defaultLoadData, err := loadData(defaultConfig); err != nil {
			return err
		} else {
			defaultData = defaultLoadData
		}
	}
	envData = loadFromEnv()
	if confPath, ok := searchConf(); ok {
		return Load(confPath)
	}
	return nil
}

func Load(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if dataInFile, err := loadData(data); err != nil {
		return err
	} else {
		dataLock.Lock()
		userData = dataInFile
		dataLock.Unlock()
		return nil
	}
}

func Save() error {
	return ioutil.WriteFile(loadedConfigPath, saveData(userData), 0666)
}

func GetLastError() error {
	return lastError
}

type dumpInfo struct {
	Value string
	From  string
}

func PrintForDebug() {
	allData := make(map[string]dumpInfo)
	for k, v := range defaultData {
		allData[k] = dumpInfo{v, "default"}
	}
	for k, v := range envData {
		allData[k] = dumpInfo{v, "env"}
	}
	for k, v := range userData {
		allData[k] = dumpInfo{v, "user"}
	}
	for k, v := range allData {
		fmt.Printf("[libsetting:%s] %s => %s\n", v.From, k, v.Value)
	}
}
