package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/elastic/beats/libbeat/generator/fields"
)

func main() {
	esBeatsPath := flag.String("es_beats_path", "..", "Path to elastic/beats")
	beatPath := flag.String("beat_path", ".", "Path to your Beat")
	flag.Parse()

	beatFieldsPath := flag.Args()
	name := filepath.Base(*beatPath)

	err := os.MkdirAll(filepath.Join(*beatPath, "_meta"), 0744)
	if err != nil {
		fmt.Printf("Cannot creata _meta dir for %s: %v\n", name, err)
		os.Exit(1)
	}

	if len(beatFieldsPath) == 0 {
		fmt.Println("No field files to collect")
		err = fields.AppendFromLibbeat(*esBeatsPath, *beatPath)
		if err != nil {
			fmt.Printf("Cannot generate global fields.yml for %s: %v\n", name, err)
			os.Exit(2)
		}
		return
	}

	if *beatPath == "" {
		fmt.Println("beat_path cannot be empty")
		os.Exit(1)
	}

	pathToModules := filepath.Join(*beatPath, beatFieldsPath[0])
	fieldFiles, err := fields.CollectModuleFiles(pathToModules)
	if err != nil {
		fmt.Printf("Cannot collect fields.yml files: %v\n", err)
		os.Exit(2)

	}

	err = fields.Generate(*esBeatsPath, *beatPath, fieldFiles)
	if err != nil {
		fmt.Printf("Cannot generate global fields.yml file for %s: %v\n", name, err)
		os.Exit(3)
	}

	fmt.Printf("Generated fields.yml for %s\n", name)
}
