package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type fileInfo struct {
	name    string
	path    string
	content string
}

func createTestFilesWithPaths(files []fileInfo) (string, error) {
	dir, err := ioutil.TempDir(os.TempDir(), "tfsec")
	if err != nil {
		return "", err
	}

	for _, file := range files {
		path := filepath.Join(
			dir,
			file.path,
		)
		if err := os.MkdirAll(path, 0755|os.ModeDir); err != nil {
			return "", err
		} else if err := ioutil.WriteFile(
			filepath.Join(
				path,
				file.name,
			),
			[]byte(file.content),
			0755,
		); err != nil {
			return "", err
		}

	}

	return dir, nil
}

func Benchmark_ModuleParsing(b *testing.B) {

	path := createTestFileWithModule(`
module "my-mod" {
	source = "../module"
	input = "ok"
}

output "result" {
	value = module.my-mod.result
}
`,
		`
variable "input" {
	default = "?"
}

output "result" {
	value = var.input
}
`,
	)

	for i := 0; i < b.N; i++ {
		parser := New()
		if _, err := parser.ParseDirectory(path, nil, ""); err != nil {
			b.Fatal(err)
		}
	}

}

func Benchmark_MultiModuleParsing2(b *testing.B) {

	path, err := createTestFilesWithPaths(
		generateModuleContent(2),
	)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		parser := New()
		if _, err := parser.ParseDirectory(path, nil, ""); err != nil {
			b.Fatal(err)
		}
	}

}

func generateModuleContent(numModules int) []fileInfo {
	var files []fileInfo
	path := "entry"
	for i := 0; i < numModules; i++ {
		if i > 0 {
			path += fmt.Sprintf("/module%d", i)
		}
		files = append(files, fileInfo{
			path: path,
			name: "main.tf",
			content: fmt.Sprintf(`
module "my-mod" {
	source = "./module%d"
	input = "ok"
}
`, i+1),
		}, fileInfo{
			path: path,
			name: "variables.tf",
			content: `
variable "input" {
	default = "?"
}
`,
		}, fileInfo{
			path: path,
			name: "outputs.tf",
			content: `
output "result" {
	value = module.my-mod.result
}
`,
		})
	}

	// Add on the last module to complete the chain
	files = append(files, fileInfo{
		path: path + fmt.Sprintf("/module%d", numModules),
		name: "main.tf",
		content: `
variable "input" {
	default = "?"
}

output "result" {
	value = var.input
}
`,
	})
	return files
}

func Benchmark_MultiModuleParsing4(b *testing.B) {
	path, err := createTestFilesWithPaths(
		generateModuleContent(4),
	)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		parser := New()
		if _, err := parser.ParseDirectory(path, nil, ""); err != nil {
			b.Fatal(err)
		}
	}

}

func Benchmark_MultiModuleParsing8(b *testing.B) {
	path, err := createTestFilesWithPaths(
		generateModuleContent(8),
	)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		parser := New()
		if _, err := parser.ParseDirectory(path, nil, ""); err != nil {
			b.Fatal(err)
		}
	}

}

func Benchmark_MultiModuleParsing16(b *testing.B) {
	path, err := createTestFilesWithPaths(
		generateModuleContent(16),
	)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		parser := New()
		if _, err := parser.ParseDirectory(path, nil, ""); err != nil {
			b.Fatal(err)
		}
	}

}

func Benchmark_MultiModuleParsing32(b *testing.B) {
	path, err := createTestFilesWithPaths(
		generateModuleContent(32),
	)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		parser := New()
		if _, err := parser.ParseDirectory(path, nil, ""); err != nil {
			b.Fatal(err)
		}
	}

}
