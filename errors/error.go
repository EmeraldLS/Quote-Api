package myerrors

func DefinedErrors() map[string]string {
	myErrors := map[string]string{
		"EnvError":         "An error occured while loading the '.env' file. Error detail is listed below",
		"ClientOptionsErr": "An error occured while parsing the client options",
		"ClientErr":        "An error occured while parsing the mongo client",
	}
	return myErrors
}
