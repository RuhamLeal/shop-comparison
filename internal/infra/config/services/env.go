package services

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnvironmentVariable(environmentVariable string, omitEmpty bool) string {
	value, exists := os.LookupEnv(environmentVariable)

	if !exists {
		panic(fmt.Sprintf("Missing [%s] environment variable", environmentVariable))
	}

	if !omitEmpty && value == "" {
		panic(fmt.Sprintf("Environment variable [%s] is not set", environmentVariable))
	}

	return value
}

func GetEnvironmentVariableAsBool(environmentVariable string, omitEmpty bool) bool {
	value := GetEnvironmentVariable(environmentVariable, omitEmpty)

	if value == "" {
		return false
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(fmt.Sprintf("Environment variable [%s] must be a boolean value (true/false, 1/0)", environmentVariable))
	}

	return boolValue
}
