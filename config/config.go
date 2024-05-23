package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type Database struct {
	Name     string `json:"name"`     // имя базы данных
	User     string `json:"user"`     // имя пользователя
	Password string `json:"password"` // пароль
	Address  string `json:"address"`  // адрес сервера базы данных
	Port     string `json:"port"`     // порт сервера базы данных
}

type Config struct {
	Tables       []string `json:"tables"`        // список таблиц для дампа
	SSHUser      string   `json:"ssh_user"`      // пользователь для SSH-соединения
	RemoteServer string   `json:"remote_server"` // адрес удаленного сервера
	LocalDB      Database `json:"local_db"`      // информация о подключении к локальной базе данных
	RemoteDB     Database `json:"remote_db"`     // информация о подключении к удаленной базе данных
	SSHKeyPath   string   `json:"ssh_key_path"`  // путь до ключа для SSH-соединения
	MaxRoutines  int      `json:"max_routines"`  // максимальное количество горутин
	CompressDump bool     `json:"compress_dump"` // необходимо ли сжимать данные true\false
}

func LoadConfig(filename string) Config {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	checkConfigParameters(cfg)

	if cfg.MaxRoutines == 0 {
		cfg.MaxRoutines = 5
	}

	return cfg
}

func checkConfigParameters(cfg Config) {
	missingParams := []string{}

	if len(cfg.Tables) == 0 {
		missingParams = append(missingParams, "Tables")
	}
	if cfg.SSHUser == "" {
		missingParams = append(missingParams, "SSHUser")
	}
	if cfg.RemoteServer == "" {
		missingParams = append(missingParams, "RemoteServer")
	}
	if cfg.SSHKeyPath == "" {
		missingParams = append(missingParams, "SSHKeyPath")
	}
	if cfg.LocalDB.Name == "" {
		missingParams = append(missingParams, "LocalDB.Name")
	}
	if cfg.LocalDB.User == "" {
		missingParams = append(missingParams, "LocalDB.User")
	}
	if cfg.LocalDB.Password == "" {
		missingParams = append(missingParams, "LocalDB.Password")
	}
	if cfg.LocalDB.Address == "" {
		missingParams = append(missingParams, "LocalDB.Address")
	}
	if cfg.RemoteDB.Name == "" {
		missingParams = append(missingParams, "RemoteDB.Name")
	}
	if cfg.RemoteDB.User == "" {
		missingParams = append(missingParams, "RemoteDB.User")
	}
	if cfg.RemoteDB.Password == "" {
		missingParams = append(missingParams, "RemoteDB.Password")
	}
	if cfg.RemoteDB.Address == "" {
		missingParams = append(missingParams, "RemoteDB.Address")
	}

	if len(missingParams) > 0 {
		log.Fatal("Все параметры в файле конфигурации должны быть заполнены. Незаполненные параметры: ", strings.Join(missingParams, ", "))
	}
}
