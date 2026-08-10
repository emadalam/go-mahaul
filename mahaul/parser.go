package mahaul

import (
	"errors"
	"net/url"
	"strconv"
)

func (c EnvVarConf) Parse(val string) (pVal any, err error) {
	switch c.Type {
	case StringType:
		pVal, err = ParseString(val)
	case NumberType:
		pVal, err = ParseNum(val)
	case IntType:
		pVal, err = ParseInt(val)
	case BoolType:
		pVal, err = ParseBool(val)
	case PortType:
		pVal, err = ParsePort(val)
	case HostType:
		pVal, err = ParseHost(val)
	case URIType:
		pVal, err = ParseURI(val)
	}
	return pVal, err
}

func ParseString(val string) (string, error) {
	return val, nil
}

func ParseNum(val string) (float64, error) {
	fVal, err := strconv.ParseFloat(val, 64)

	if err != nil {
		return 0, err
	}

	return fVal, nil
}

func ParseInt(val string) (int, error) {
	intVal, err := strconv.Atoi(val)

	if err != nil {
		return 0, err
	}

	return intVal, nil
}

func ParseBool(val string) (bool, error) {
	boolVal, err := strconv.ParseBool(val)

	if err != nil {
		return false, err
	}

	return boolVal, nil
}

func ParsePort(val string) (uint16, error) {
	port, err := strconv.ParseUint(val, 10, 16)

	if err != nil {
		return 0, err
	}

	return uint16(port), nil
}

func ParseHost(val string) (*url.URL, error) {
	url, err := url.Parse(val)

	if err != nil {
		return nil, err
	}
	if url.Host == "" {
		return nil, errors.New("invalid host name")
	}
	if url.Path != "" {
		return nil, errors.New("invalid host name, must not contain paths")
	}
	if url.Scheme != "" {
		return nil, errors.New("invalid host name, must not contain uri scheme")
	}

	return url, nil
}

func ParseURI(val string) (*url.URL, error) {
	url, err := url.Parse(val)

	if err != nil {
		return nil, err
	}
	if url.Host == "" {
		return nil, errors.New("invalid uri host name")
	}
	if url.Scheme == "" {
		return nil, errors.New("invalid uri scheme name")
	}

	return url, nil
}
