// Copyright 2022 Zinc Labs Inc. and Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/blugelabs/bluge/index"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: snp <PATH>")
		os.Exit(1)
	}

	rootPath := os.Args[1]
	snps, err := scanSnapshot(rootPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for name, snp := range snps {
		fmt.Printf("Found snapshot: %s\n", name)
		for _, segSnapshot := range snp.Segments() {
			var numDeleted int
			if segSnapshot.Deleted() != nil {
				numDeleted = int(segSnapshot.Deleted().GetCardinality())
			}
			docTimeMin, docTimeMax := segSnapshot.Timestamp()
			fmt.Printf("segment id: %012x num_docs: %8d, num_deleted: %4d, doc_time_min: %d, doc_time_max: %d, segment_size: %d\n",
				segSnapshot.ID(), segSnapshot.DocNum(), numDeleted, docTimeMin, docTimeMax, segSnapshot.SegmentSize())
		}
	}

	segs, err := scanSegment(rootPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, seg := range segs {
		fmt.Printf("Found segment: %s", seg)
		fmt.Printf("... checking... ")

		exists := false
		for _, snp := range snps {
			for _, segSnapshot := range snp.Segments() {
				id := fmt.Sprintf("%012x.seg", segSnapshot.ID())
				if id == seg {
					exists = true
					break
				}
			}
		}

		if exists {
			fmt.Printf("OK")
		} else {
			fmt.Printf("UNUSED... deleting... ")
			if err := os.Remove(path.Join(rootPath, seg)); err != nil {
				fmt.Printf("Err: %v", err)
			} else {
				fmt.Printf("OK")
			}
		}

		fmt.Println("")
	}
}

func scanSnapshot(rootPath string) (map[string]*index.Snapshot, error) {
	fs, err := os.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}

	ss := make(map[string]*index.Snapshot, 0)
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		if path.Ext(f.Name()) != ".snp" {
			continue
		}
		snp, err := readSnapshot(path.Join(rootPath, f.Name()))
		if err != nil {
			return nil, err
		}
		ss[f.Name()] = snp
	}

	return ss, nil
}

func readSnapshot(filepath string) (*index.Snapshot, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	br := bufio.NewReader(f)
	defer f.Close()

	snp := new(index.Snapshot)
	_, err = snp.ReadFrom(br)
	return snp, err
}

func scanSegment(rootPath string) ([]string, error) {
	fs, err := os.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}

	ss := make([]string, 0)
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		if path.Ext(f.Name()) != ".seg" {
			continue
		}
		ss = append(ss, f.Name())
	}

	return ss, nil
}
