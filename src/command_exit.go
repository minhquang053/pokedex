package main

import "os"

func commandExit(cf *config) error {
	os.Exit(0)
	return nil
}
