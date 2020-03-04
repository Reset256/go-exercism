package flatten

func Flatten(input interface{}) []interface{} {
	slice := make([]interface{}, 0)
	result := &slice
	for _, value := range input.([]interface{}) {
		goDeep(value, result)
	}
	return *result
}

func goDeep(income interface{}, result *[]interface{}) {
	switch income.(type) {
	case []interface{}:
		for _, value := range income.([]interface{}) {
			goDeep(value, result)
		}
	default:
		if income != nil {
			*result = append(*result, income)
		}
	}
}
