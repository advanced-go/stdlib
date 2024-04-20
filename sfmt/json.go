package sfmt

// JsonString - Json format a string value
func JsonString(value string) string {
	if len(value) == 0 {
		return "null"
	}
	return "\"" + value + "\""
}
