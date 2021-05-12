package libsetting

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"
)

var (
	userData    map[string]string
	dataLock    sync.RWMutex
	envData     map[string]string // readonly
	defaultData map[string]string // readonly
)

// getter

func getData(key string, datas ...map[string]string) (value string, ok bool) {
	if len(datas) > 0 {
		for _, data := range datas {
			if data != nil {
				dataLock.RLock()
				if v, ok := data[strings.ToUpper(key)]; ok {
					return v, true
				}
				dataLock.RUnlock()
			}
		}
	}
	return "", false
}

func GetString(key string) (value string, ok bool) {
	return getData(key, userData, envData, defaultData)
}

func GetInt(key string) (value int, ok bool) {
	if val, ok := GetString(key); ok {
		if i, err := strconv.Atoi(val); err == nil {
			return i, true
		} else {
			lastError = err
		}
	}
	return 0, false
}

func GetBool(key string) (value bool, ok bool) {
	if val, ok := GetString(key); ok {
		return strings.EqualFold(val, "true"), true
	}
	return false, false
}

func GetObject(key string) (value interface{}, ok bool) {
	if val, ok := GetString(key); ok {
		if err := json.Unmarshal([]byte(val), &value); err == nil {
			return value, true
		} else {
			lastError = err
		}
	}
	return nil, false
}

// setter

func setData(key, value string) {
	dataLock.Lock()
	userData[strings.ToUpper(key)] = value
	dataLock.Unlock()
}

func SetString(key, value string) {
	setData(key, value)
}

func SetInt(key string, value int) {
	setData(key, strconv.Itoa(value))
}

func SetBool(key string, value bool) {
	if value {
		setData(key, "true")
	} else {
		setData(key, "false")
	}
}

func SetObject(key string, value interface{}) {
	if data, err := json.Marshal(value); err == nil {
		SetString(key, string(data))
	} else {
		lastError = err
	}
}
