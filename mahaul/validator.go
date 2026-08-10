package mahaul

import (
	"fmt"
	"os"
)

func ValidateEnv(config map[string]EnvVarConf) (errs map[string]error) {
	for key, c := range config {
		_, err := c.LookupEnv(key)

		addError(&errs, key, err)
	}

	return errs
}

func WarnValidateEnv(config map[string]EnvVarConf) (errs map[string]error) {
	errs = ValidateEnv(config)

	for key, err := range errs {
		fmt.Fprintf(os.Stderr, "%v: %v\n", key, err)
	}
	return errs
}

func MustValidateEnv(config map[string]EnvVarConf) {
	errs := WarnValidateEnv(config)

	if errs != nil {
		os.Exit(1)
	}
}

func addError(errs *map[string]error, key string, err error) {
	if err != nil {
		if *errs == nil {
			*errs = make(map[string]error)
		}
		(*errs)[key] = err
	}
}
