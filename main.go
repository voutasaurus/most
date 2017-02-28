package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

var newline = []byte("\n")

var flagPercent = flag.Int("p", 50, "percentage of line numbers to account for in listed output")

func main() {
	flag.Parse()
	if *flagPercent < 0 || *flagPercent > 100 {
		flag.Usage()
		return
	}
	if err := most(*flagPercent); err != nil {
		log.Fatal(err)
	}
}

func most(p int) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("s: error getting working directory: %v", err)
	}
	rec := &recorder{}
	if err := filepath.Walk(dir, rec.recordWalk()); err != nil {
		return fmt.Errorf("s: error walking filesystem %v", err)
	}
	sort.Slice(rec.rr, func(i, j int) bool {
		return rec.rr[i].length > rec.rr[j].length
	})
	cum := 0
	for _, r := range rec.rr {
		cum += r.length
		fmt.Printf("%8d, %s\n", r.length, r.path[len(dir)+1:])
		if cum > (p*rec.sum)/100 && cum != rec.sum {
			fmt.Println("...")
			return nil
		}
	}
	return nil
}

type recorder struct {
	rr  []*record
	sum int
}

type record struct {
	length int
	path   string
}

func (r *recorder) recordWalk() filepath.WalkFunc {
	f := func(path string, info os.FileInfo, err error) error {
		n := info.Name()
		ignore := n == "" || n[0] == '.' || n[0] == '_' || n == "vendor"

		if info.IsDir() {
			if ignore {
				return filepath.SkipDir
			}
			return nil
		}
		if !ignore {
			return r.record(path)
		}
		return nil
	}
	return filepath.WalkFunc(f)
}

func (r *recorder) record(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("subf: error opening file: %v", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("subf: error reading file: %v", err)
	}
	n := bytes.Count(b, newline)
	r.rr = append(r.rr, &record{
		path:   path,
		length: n,
	})
	r.sum += n
	return nil
}
