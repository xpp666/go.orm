package util

import (
	"os"
	"fmt"
	"encoding/json"
)

var path ="./conf"

type Config struct {   //配置文件要通过tag来指定配置文件中的名称
	Mysql_host string
	Mysql_port string
	Mysql_user string
	Mysql_password string
	Database_name string
	Max_idle_conns   string
	Max_open_conns string
}

func ConfigUtil() Config{
	file, _ := os.Open(path)
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return conf
}

