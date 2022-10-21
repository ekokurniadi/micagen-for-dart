package generator_v2

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GenerateBuildYaml() {
	createFileBuildYaml()
	writeFileBuildYaml()
}

func writeFileBuildYaml() (string, error) {
	filepath, err := filepath.Abs("././build.yaml")
	if err != nil {
		log.Fatal("error")
		return filepath, err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// //Write some text line-by-line to file.
	_, _ = file.WriteString("targets:\n")
	_, _ = file.WriteString("  $default:\n")
	_, _ = file.WriteString("    sources:\n")
	_, _ = file.WriteString("      - $package$\n")
	_, _ = file.WriteString("      - $lib/$lib$\n")
	_, _ = file.WriteString("      - lib/**.dart\n")
	_, _ = file.WriteString("      - test/**.dart\n")
	_, _ = file.WriteString("    builders:\n")
	_, _ = file.WriteString("      source_gen|combining_builder:\n")
	_, _ = file.WriteString("        options:\n")
	_, _ = file.WriteString("          ignore_for_file:\n")
	_, _ = file.WriteString("            - implicit_dynamic_parameter\n")
	_, _ = file.WriteString("      json_serializable:\n")
	_, _ = file.WriteString("        generate_for:\n")
	_, _ = file.WriteString("            - lib/**.codegen.dart\n")
	_, _ = file.WriteString("        options:\n")
	_, _ = file.WriteString("          field_rename: snake\n")
	_, _ = file.WriteString("      freezed:freezed:\n")
	_, _ = file.WriteString("        generate_for:\n")
	_, _ = file.WriteString("            - lib/**/blocs/**.dart\n")
	_, _ = file.WriteString("            - lib/**/cubit/**.dart\n")
	_, _ = file.WriteString("            - lib/**/codegen/**.dart\n")
	_, _ = file.WriteString("      mockito|mockBuilder:\n")
	_, _ = file.WriteString("        generate_for:\n")
	_, _ = file.WriteString("            - test/**.dart\n")

	return "", nil

}

func createFileBuildYaml() (string, error) {
	filepath, err := filepath.Abs("././build.yaml")

	if err != nil {
		log.Fatal("error")
		return filepath, err
	}
	_, err = os.Stat(filepath)
	if os.IsExist(err) {
		fmt.Println("your directory is already exist but it's ok")
		return "", err
	}

	filename, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err.Error())
		return filename.Name(), err
	}

	fmt.Printf("Create %s is successfully \n", filename.Name())
	return filepath, nil
}
