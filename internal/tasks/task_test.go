package tasks

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nathan-wien/qkly/internal/filesystem"
	"github.com/nathan-wien/qkly/internal/judges"
	"github.com/spf13/afero"
)

func TestTaskDir(t *testing.T) {
	fsMgr := filesystem.NewFileSystemManager(afero.NewMemMapFs())

	t.Run("Normal case", func(t *testing.T) {
		xDir := "/project/dir/codeforces/abc123/a"
		task := Task{judges.CodeForces{}, nil, "abc123", "a", fsMgr, "/project/dir", ""}
		dir := task.Dir()
		if dir != xDir {
			t.Fatalf("expected %s, got %s\n", xDir, dir)
		}
	})
}

func TestCreateTask(t *testing.T) {
	fsMgr := filesystem.NewFileSystemManager(afero.NewMemMapFs())

	testContent := make(map[string]string, 4)
	testContent["1.in.txt"] = "4 5\n.....\n.###.\n.###.\n....."
	testContent["1.out.txt"] = "1"
	testContent["2.in.txt"] = "2 3\n...\n..."
	testContent["2.out.txt"] = "1"

	task := Task{
		Judge: judges.CodeForces{},
		Data: &TaskData{
			Url: `https://codeforces.com/problemset/problem/348/D`,
			Tests: []TestData{
				{
					Input:  testContent["1.in.txt"],
					Output: testContent["1.out.txt"],
				},
				{
					Input:  testContent["2.in.txt"],
					Output: testContent["2.out.txt"],
				},
			},
		},
		ContestId:   "348",
		TaskId:      "d",
		FsMgr:       fsMgr,
		TemplateDir: DefaultTemplateDir,
	}

	dir := "./codeforces/348/d"

	t.Run("TestWriteTestFiles", func(t *testing.T) {
		if err := task.writeTestFiles(); err != nil {
			t.Fatalf("Error: %v\n", err.Error())
		}

		if exists, err := afero.DirExists(fsMgr.Fs, dir); !exists {
			t.Fatalf("Error: Directory \"./codeforces/348/d\" was not created with err %v\n", err)
		}

		fileInfo, err := afero.ReadDir(fsMgr.Fs, dir)
		if err != nil {
			t.Fatalf("Error while trying to get all files in \"%s\"", dir)
		}

		for _, file := range fileInfo {
			content, err := afero.ReadFile(fsMgr.Fs, filepath.Join(dir, file.Name()))
			if err != nil {
				t.Fatalf("Error: cannot read file %s\n", file.Name())
			}
			if string(content) != testContent[file.Name()] {
				t.Fatalf("Error: content of file %s is incorrect\n", file.Name())
			}
		}
	})

	t.Run("TestCopyConfigFiles", func(t *testing.T) {
		if err := fsMgr.Fs.MkdirAll(task.TemplateDir, os.ModePerm); err != nil {
			t.Fatalf("Error: %s\n", err.Error())
		}

		f, err := fsMgr.Fs.Create(filepath.Join(task.TemplateDir, "main.cpp"))
		if err != nil {
			t.Fatalf("Error: %s\n", err.Error())
		}

		writeContent :=
			`
#include <iostream>

int main() {
	std::cout << "Hello\n";
	return 0;
}
`
		_, err = f.WriteString(writeContent)
		if err != nil {
			t.Fatalf("Error: %s\n", err.Error())
		}
		f.Close()

		task.copyTemplateFiles()

		readBytes, err := fsMgr.Util.ReadFile(filepath.Join(dir, "main.cpp"))
		readContent := string(readBytes)
		if err != nil {
			t.Fatalf("Error: %s\n", err.Error())
		}
		if readContent != writeContent {
			t.Fatalf("Error: template file content is incorrect")
		}
		f.Close()
	})
}
