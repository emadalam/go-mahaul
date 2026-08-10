package mahaul

import "reflect"

type EnvVarType string

const (
	StringType EnvVarType = "string"
	NumberType EnvVarType = "num"
	IntType    EnvVarType = "int"
	BoolType   EnvVarType = "bool"
	PortType   EnvVarType = "port"
	HostType   EnvVarType = "host"
	URIType    EnvVarType = "uri"
)

type EnvVarConf struct {
	Type    EnvVarType
	Default string
	Choices []any
}

// TODO: maybe have a value field that contains the boot time env values

type EnvVals struct{}

func (et EnvVarType) GoType() (gt string) {
	switch et {
	case StringType:
		gt = reflect.TypeOf(ParseString).Out(0).String()
	case NumberType:
		gt = reflect.TypeOf(ParseNum).Out(0).String()
	case IntType:
		gt = reflect.TypeOf(ParseInt).Out(0).String()
	case BoolType:
		gt = reflect.TypeOf(ParseBool).Out(0).String()
	case PortType:
		gt = reflect.TypeOf(ParsePort).Out(0).String()
	case HostType:
		gt = reflect.TypeOf(ParseHost).Out(0).String()
	case URIType:
		gt = reflect.TypeOf(ParseURI).Out(0).String()
	}

	return gt
}
