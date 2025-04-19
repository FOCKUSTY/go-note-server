package env

import "os"

func InitEnv() {
	os.Setenv("PORT", "8000")
}

func Get(name string) (string, bool) {
	return os.LookupEnv(name)
}
