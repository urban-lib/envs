package envs

import (
	"strconv"
	"time"
)

func Get(name string) string {
	if e, ok := envs[name]; ok {
		return e.Value()
	} else {
		return ""
	}
}

func GetInt(name string) int {
	value, _ := strconv.Atoi(Get(name))
	return value
}

func GetDuration(name string) time.Duration {
	value := GetInt(name)
	return time.Duration(value)
}
