package envs

import (
	"log"
	"os"
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
	formatLog(nil, "*********** Project environments ***********")
	for _, e := range envs {
		formatLog(e)
	}
	formatLog(nil, "********************************************")
}

func formatLog(e *env, params ...interface{}) {
	if e != nil {
		if e.Require && len(os.Getenv(e.Name)) == 0 && len(e.Default) == 0 {
			log.Fatalf("Require environment is empty: %s", e.Name)
		}
		log.Println("\t", e.Name, len(os.Getenv(e.Name)) > 0)
	} else {
		log.Println(params...)
	}
}
