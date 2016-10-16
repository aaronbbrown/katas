package main

import (
	"fmt"
	"os"
	"strconv"
)

// get an environment variable as an integer with a default
func GetEnvNDefault(key string, defValue int) (n int, err error) {
	s, exists := os.LookupEnv(key)
	if exists {
		n, err = strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("%s must be an integer.  Got %s", key, s)
		}
	} else {
		n = 10
	}
	return n, nil
}
