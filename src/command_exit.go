package main

import "os"

func commandExit(cf *config, args ...string) error {
	os.Exit(0)
	return nil
}
