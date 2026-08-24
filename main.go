package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type service struct {
	Name string
	Host string
	Port int
	Tags []string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: service-mesh-manager <registry.json> [list|register|route]")
		os.Exit(1)
	}
	registryPath := os.Args[1]
	data, _ := os.ReadFile(registryPath)
	var services []service
	if len(data) > 0 {
		json.Unmarshal(data, &services)
	}
	cmd := "list"
	if len(os.Args) > 2 {
		cmd = os.Args[2]
	}
	switch cmd {
	case "list":
		for _, s := range services {
			fmt.Printf("%-20s %s:%d %v\n", s.Name, s.Host, s.Port, s.Tags)
		}
		fmt.Printf("\n%d services registered\n", len(services))
	case "register":
		if len(os.Args) < 5 {
			fmt.Fprintln(os.Stderr, "need name host:port")
			os.Exit(1)
		}
		s := service{Name: os.Args[3], Host: os.Args[4]}
		services = append(services, s)
		b, _ := json.Marshal(services)
		os.WriteFile(registryPath, b, 0644)
		fmt.Println("registered", s.Name)
	case "route":
		if len(os.Args) < 4 {
			fmt.Fprintln(os.Stderr, "need service name")
			os.Exit(1)
		}
		name := os.Args[3]
		for _, s := range services {
			if s.Name == name {
				fmt.Printf("route: %s -> %s:%d\n", name, s.Host, s.Port)
				return
			}
		}
		fmt.Println("service not found")
		os.Exit(1)
	}
	_ = sort.IntSlice{}
}
