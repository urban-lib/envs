```bash
export APP_CONFIG=develop
export APP_DB_PASSWORD=<password>
```

or `.env` if using `github.com/joho/godotenv`

```bash
# .env
APP_CONFIG=develop
APP_DB_PASSWORD=<password>
```

```go
package main

import "github.com/urban-lib/envs"

func init() {
	//      Environment name, required, default value  
	envs.NewEnv("APP_CONFIG", false, "default")
	envs.NewEnv("APP_DB_PASSWORD", true, "")
}

func main() {
	// check existing environments
	envs.PrintAppEnvironments()
	...
}

type Config struct {
	UseConf    string
	DbPassword string
}

func initConfig() *Config {
	return &Config{
		UseConf:    envs.Get("APP_CONFIG"),
		DbPassword: envs.Get("APP_DB_PASSWORD"),
	}
}
```