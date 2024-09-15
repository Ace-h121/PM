package tree

import (
	"fmt"
	"os"
	"path"
	"sort"
)

type Counter struct {
	dirs  int
	files int
}

func (counter *Counter) index(path string) {
	stat, _ := os.Stat(path)
	if stat.IsDir() {
		counter.dirs += 1
	} else {
		counter.files += 1
	}
}

func (counter *Counter) output() string {
	return fmt.Sprintf("\n%d directories, %d files", counter.dirs, counter.files)
}

func dirnamesFrom(base string) []string {
	file, err := os.Open(base)
	if err != nil {
		fmt.Println(err)
	}

	names, _ := file.Readdirnames(0)
	file.Close()

	sort.Strings(names)
	return names
}

func tree(counter *Counter, base string, prefix string) {
	names := dirnamesFrom(base)

	for index, name := range names {
		if name[0] == '.' {
			continue
		}
		subpath := path.Join(base, name)
		counter.index(subpath)

		if index == len(names)-1 {
			fmt.Println(prefix+"└──", name)
			tree(counter, subpath, prefix+"    ")
		} else {
			fmt.Println(prefix+"├──", name)
			tree(counter, subpath, prefix+"│   ")
		}
	}
}

func Run() error {
	dir, err := os.UserHomeDir()
	if err !=nil {
		return err
	}
	dir = dir + "/PM/"
	counter := new(Counter)
	tree(counter, dir, "")
	fmt.Println(counter.output())
	return nil
}
