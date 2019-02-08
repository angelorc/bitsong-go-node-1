package utils

import (
	"flag"
	"os"
	"path/filepath"
)

var (
	BitsongHome = flag.String("home", "", "Path to bitsong data directory")
)

func init() {
	flag.Parse()
}

func GetBitsongHome() string {
	if *BitsongHome != "" {
		return *BitsongHome
	}

	home := os.Getenv("BITSONGHOME")

	if home != "" {
		return home
	}

	return os.ExpandEnv(filepath.Join("$HOME", ".bitsongd"))
}
