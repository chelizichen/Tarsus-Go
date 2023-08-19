package main

import (
	"fmt"
	"regexp"
	_ "strings"
)

type ServerConfig struct {
	Project  string `json:"project"`
	Language string `json:"language"`
	Proto    string `json:"proto"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Weight   string `json:"weight"`
}

func main() {
	input := "@TarsusDemoProject/GoServer -l go -t @tarsus/ms -h 127.0.0.1 -p 12014 -w 10"
	serverConfig := parseServerConfig(input)
	fmt.Printf("Parsed Server Config: %+v\n", serverConfig)
}

func parseServerConfig(input string) ServerConfig {
	re := regexp.MustCompile(`@([^\/\s]+)\/([^\/\s]+) -l ([^\s]+) -t ([^\s]+) -h ([^\s]+) -p ([^\s]+) -w ([^\s]+)`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 8 {
		return ServerConfig{}
	}

	return ServerConfig{
		Project:  matches[1] + "/" + matches[2],
		Language: matches[3],
		Proto:    matches[4],
		Host:     matches[5],
		Port:     matches[6],
		Weight:   matches[7],
	}
}
