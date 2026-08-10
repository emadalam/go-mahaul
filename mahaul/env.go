package mahaul

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

func (c EnvVarConf) LookupEnv(key string) (any, error) {
	val, isSet := os.LookupEnv(key)
	if !isSet && c.Default == "" {
		return nil, errors.New("not set")
	}

	if val == "" {
		val = c.Default
	}

	parsedVal, err := c.Parse(val)

	if err != nil {
		return nil, err
	}

	if c.Choices != nil && !slices.Contains(c.Choices, parsedVal) {
		return parsedVal, fmt.Errorf("value %v not in %v", parsedVal, c.Choices)
	}

	return parsedVal, nil
}

func (c EnvVarConf) Getenv(key string) any {
	val, _ := c.LookupEnv(key)
	return val
}

func GetenvAs[T any](c EnvVarConf, key string) T {
	var zero T

	v := c.Getenv(key)
	if v == nil {
		return zero
	}

	val, ok := v.(T)
	if !ok {
		return zero
	}

	return val
}
