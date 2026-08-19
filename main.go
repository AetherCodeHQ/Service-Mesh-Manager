package main

import (
	"fmt"
	"os"
)

// service_mesh_manager - Manage service mesh configs
func service_mesh_manager(path string) {
	fmt.Println("========================================")
	fmt.Println("  Service-Mesh-Manager")
	fmt.Println("  Manage service mesh configs")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	service_mesh_manager(path)
}
