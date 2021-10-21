package conf

import (
	"fmt"
	"os"

	toml "github.com/pelletier/go-toml"
)

type API struct {
	IP   string
	Port int64
}

func getEnv() string {
	if os.Getenv("APP_ENV") == "" {
		return "DEV"
	}

	return os.Getenv("APP_ENV")
}

func readConf() (string, int64) {
	var ip string
	var port int64

	doc, err := toml.LoadFile("config.toml")

	env := getEnv()

	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		ip = doc.Get(fmt.Sprintf("API.%s.IP", env)).(string)
		port = doc.Get(fmt.Sprintf("API.%s.Port", env)).(int64)
	}

	return ip, port
}

// func main() {
func GetAPIConf() (string, int64) {
	// Get a greeting message and print it.
	env := getEnv()
	// message := "Gladys"
	fmt.Println(env)

	// x, y := readConf()
	return readConf()
}
