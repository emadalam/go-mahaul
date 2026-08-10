package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/emadalam/go-mahaul/mahaul"
)

func main() {
	configPath := flag.String("config", "env.config.json", "Path to config file")
	outputPath := flag.String("out", "env_generated.go", "Output file name")
	packageName := flag.String("pkgName", "env", "Generated package name")
	flag.Parse()

	cf, err := os.OpenFile(*configPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer cf.Close()

	cc, err := io.ReadAll(cf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	var config = make(map[string]mahaul.EnvVarConf, 0)
	json.Unmarshal(cc, &config)

	var generator = mahaul.Generator{}
	generator.Make(mahaul.GeneratorCnf{PackageName: *packageName, Config: &config})

	contents, err := generator.Generate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(*outputPath, contents, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
