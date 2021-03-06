package dbt

import (
	"fmt"
	"github.com/nikogura/gomason/pkg/gomason"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestWriteCreateTool(t *testing.T) {
	// create workspace
	lang, _ := gomason.GetByName("golang")
	goPath, err := lang.CreateWorkDir(tmpDir)
	if err != nil {
		log.Printf("Error creating gopath in %q: %s", tmpDir, err)
		t.FailNow()
	}

	fmt.Printf("Created gopath: %s\n", goPath)

	pkgDir := fmt.Sprintf("%s/src", goPath)

	err = os.MkdirAll(pkgDir, 0755)
	if err != nil {
		fmt.Printf("failed to create package directory %s", pkgDir)
		t.Fail()
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed to get current working dir: %s", err)
		t.Fail()
	}

	err = os.Chdir(pkgDir)
	if err != nil {
		fmt.Printf("failed to cwd to %s", pkgDir)
		t.Fail()
	}

	err = WriteConfigFiles(testToolName(), testPackageName(), testDescription(), testAuthor(), testToolRepo())
	if err != nil {
		fmt.Println(fmt.Sprintf("Error writing config files: %s", err))
		t.Fail()
	}

	location := fmt.Sprintf("%s/src/%s", goPath, testToolName())

	files, err := FilesForTool(location, testToolName(), testPackageName(), testDescription(), testAuthor(), testToolRepo())
	if err != nil {
		fmt.Println(fmt.Sprintf("Error generating file list: %s", err))
		t.Fail()
	}

	testFiles, err := testFilesForTool(location, testToolName(), testPackageName(), testDescription(), testAuthor(), testToolRepo())
	if err != nil {
		fmt.Println(fmt.Sprintf("Error generating test list: %s", err))
		t.Fail()
	}
	dirs := DirsForPackageName(location, testToolName())

	// Check that we have dirs
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Println(fmt.Sprintf("Expected dir %q does not exist.", dir))
			t.Fail()
		}
	}

	for index, file := range files {
		// Check that files exist.
		if _, err := os.Stat(file.Name); os.IsNotExist(err) {
			fmt.Println(fmt.Sprintf("Expected file %q does not exist.", file.Name))
			t.Fail()
		}

		// Verify that the contents is what we expect

		fmt.Printf("Verifying contents of %s\n", file.Name)

		fileContents, err := ioutil.ReadFile(file.Name)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error reading file %q: %s", file.Name, err))
			t.Fail()
		}

		assert.Equal(t, testFiles[index].Content, string(fileContents), fmt.Sprintf("Written file %q contents does not meet expectations.", file.Name))
	}

	err = os.Chdir(cwd)
	if err != nil {
		fmt.Printf("failed to cwd to %s", cwd)
		t.Fail()
	}
}
