package meister

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/skowrons/todo-meister/pkg/config"
)

// TODO: add more functions to this one
type FileIndex []string

// Walk will walk the provided path and return an index of all files including 
// their path. Each path will be relativ.
func Walk(config *config.Config, root string) FileIndex {
	index := FileIndex{}

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			levels := strings.Split(path, "/")
			dir := levels[len(levels)-1]
			
			_, ok := config.Ignores[dir]
			if ok {
				return filepath.SkipDir
			}

			return nil
		}
		

		index = append(index, path)

		return nil
	})
	
	return index
}

func NewCliConnector() {
	panic("unimplemented")
}

type Entry struct {
	KeyWord   string
	LineIndex int
	FilePath  string
	Comment string
}

func ScanFile(path string, config *config.Config) []Entry {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	entries := make([]Entry, 0)
	lineIndex := 0
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), config.Fields[0].Text) {
			entries = append(entries, Entry{
				KeyWord:   config.Fields[0].Text,
				LineIndex: lineIndex,
				FilePath:  path,
				Comment: strings.Split(scanner.Text(), config.Fields[0].Text)[1],
			})
		}
		lineIndex++
	}
	
	return entries
}

func SplitBefore(value string, substr string) []string {
	if i := strings.Index(value, substr); i >= 0 {
		return []string{value[:i], value[i:]}
	} else {
		return []string{value}
	}
}