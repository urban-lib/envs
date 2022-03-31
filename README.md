```go
package main

import "github.com/urban-lib/envs"

func init() {
	envs.NewEnv("APP_CONFIG", false, "default")
	envs.NewEnv("APP_DB_PASSWORD", true, "")
}

func main() {
	// check existing environments
	envs.PrintAppEnvironments()
	...
}
```