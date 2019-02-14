package main

import (
    "os"
    "fmt"
    "encoding/json"
)

type Config struct {
    Host                string `json:"host"`
    Port                string `json:"port"`
}

func LoadConfig(file string) Config {
    var config Config
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
        return config
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    //fmt.Printf("%+v\n", config)
    return config
}

var config = LoadConfig("config.json")
