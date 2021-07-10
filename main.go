package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
)

var (
	showsVersion = flag.Bool("v", false, "show version")

	version = "develop"
)

func main() {
	flag.Parse()

	if *showsVersion {
		fmt.Println(version)
		return
	}

	GenerateUUIDs(os.Stdin, os.Stdout)
}

// GenerateUUIDs generates UUID strings from each string line
func GenerateUUIDs(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := scanner.Text()
		// ignore empty lines for readability
		if line == "" {
			fmt.Println("")
			continue
		}

		uuid := NewUUID(line)
		fmt.Fprintf(out, "%s\n", uuid)
	}
}

// NewUUID generates UUIDv5 from string
func NewUUID(s string) uuid.UUID {
	return uuid.NewSHA1(uuid.NameSpaceOID, []byte(s))
}
