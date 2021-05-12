package libsetting

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func loadData(binData []byte) (data map[string]string, err error) {
	ret := make(map[string]string)
	bs := bufio.NewScanner(bytes.NewBuffer(binData))
	for {
		if !bs.Scan() {
			break
		}
		line := bs.Text()
		name, value := cutWithStr(line, "=")
		if len(value) > 0 {
			ret[strings.ToUpper(name)] = value
		}
	}
	return ret, nil
}

func saveData(data map[string]string) []byte {
	var bw bytes.Buffer
	for k, v := range data {
		bw.WriteString(k)
		bw.WriteString("=")
		bw.WriteString(v)
		bw.WriteString("\n")
	}
	return bw.Bytes()
}

func loadFromEnv() map[string]string {
	data := make(map[string]string)
	for _, v := range os.Environ() {
		name, value := cutWithStr(v, "=")
		upName := strings.ToUpper(name)
		if strings.HasPrefix(upName, applicationName) {
			_, envName := cutWithStr(upName, "_")
			if len(envName) > 0 {
				data[envName] = value
			}
		}
	}
	return data
}
