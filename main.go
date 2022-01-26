package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	ps "github.com/mitchellh/go-ps"
)

type proc struct {
	Pid        int
	Executable string
}

func main() {
	procs, err := ps.Processes()
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	plist := make([]proc, len(procs))

	for i, p := range procs {
		plist[i].Pid = p.Pid()
		plist[i].Executable = p.Executable()
	}

	sort.SliceStable(plist, func(i, j int) bool {
		return plist[i].Pid < plist[j].Pid
	})

	err = plistToFile(plist, "ProcessList.txt")
	if err != nil {
		log.Fatalf("err: %s", err)
	}

}

func plistToFile(plist []proc, filename string) error {

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, p := range plist {
		process := fmt.Sprintf("Proccess ID: %6d\t Process execuatable: %s\n", p.Pid, p.Executable)
		_, err = f.WriteString(process)
		if err != nil {
			return err
		}

	}

	fmt.Println("done")
	return nil
}
