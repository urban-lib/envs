package envs

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

var (
	envs map[string]*env
	once sync.Once
)

type Envs interface {
	Value()
}

type env struct {
	Name    string
	Require bool
	Default string
}

func (e *env) Value() string {
	result := os.Getenv(e.Name)
	if result == "" {
		return e.Default
	}
	return result
}

func NewEnv(name string, require bool, def string) {
	once.Do(func() {
		envs = make(map[string]*env)
	})
	envs[name] = &env{
		Name:    name,
		Require: require,
		Default: def,
	}
}

func CheckEnvironments() {
	keys := make([]string, 0, len(envs))
	for key := range envs {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Default", "Environment value"})
	for _, k := range keys {
		if envs[k].Require && len(os.Getenv(envs[k].Name)) == 0 && len(envs[k].Default) == 0 {
			log.Fatalf("Require environment is empty: %s", envs[k].Name)
		}
		var defaultValue, envValue string
		if strings.Contains(envs[k].Name, "PASS") || strings.Contains(envs[k].Name, "PASSWORD") {
			if envs[k].Default != "" {
				defaultValue = "******"
			}
			envValue = "******"
		} else {
			defaultValue = envs[k].Default
			envValue = os.Getenv(envs[k].Name)
		}
		t.AppendRow(table.Row{
			envs[k].Name,
			defaultValue,
			envValue,
		})
	}
	t.Render()
}

func formatLog(e *env, params ...interface{}) {
	if e != nil {

		log.Println("\t", e.Name, len(os.Getenv(e.Name)) > 0)
	} else {
		log.Println(params...)
	}
}
