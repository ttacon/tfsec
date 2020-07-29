package scanner

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/liamg/tfsec/internal/app/tfsec/parser"
)

type fileInfo struct {
	name    string
	path    string
	content string
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

func Benchmark_MultiModuleScanning2(b *testing.B) {
	path, err := createTestFilesWithPaths(
		generateModuleContent(2),
	)
	if err != nil {
		b.Fatal(err)
	}

	parsr := parser.New()
	blocks, err := parsr.ParseDirectory(path, nil, "")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		_ = New().Scan(blocks, nil)
	}
}

func Benchmark_MultiModuleScanning4(b *testing.B) {
	path, err := createTestFilesWithPaths(
		generateModuleContent(4),
	)
	if err != nil {
		b.Fatal(err)
	}

	parsr := parser.New()
	blocks, err := parsr.ParseDirectory(path, nil, "")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		_ = New().Scan(blocks, nil)
	}
}

func Benchmark_MultiModuleScanning8(b *testing.B) {
	path, err := createTestFilesWithPaths(
		generateModuleContent(8),
	)
	if err != nil {
		b.Fatal(err)
	}

	parsr := parser.New()
	blocks, err := parsr.ParseDirectory(path, nil, "")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		_ = New().Scan(blocks, nil)
	}
}

func Benchmark_MultiModuleScanning16(b *testing.B) {
	path, err := createTestFilesWithPaths(
		generateModuleContent(16),
	)
	if err != nil {
		b.Fatal(err)
	}

	parsr := parser.New()
	blocks, err := parsr.ParseDirectory(path, nil, "")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		_ = New().Scan(blocks, nil)
	}
}

func Benchmark_MultiModuleScanning32(b *testing.B) {
	path, err := createTestFilesWithPaths(
		generateModuleContent(32),
	)
	if err != nil {
		b.Fatal(err)
	}

	parsr := parser.New()
	blocks, err := parsr.ParseDirectory(path, nil, "")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		_ = New().Scan(blocks, nil)
	}
}
