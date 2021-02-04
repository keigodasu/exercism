package flatten

func Flatten(array interface{}) []interface{}  {
	flattened := []interface{}{}
	for _, v := range array.([]interface{})	 {
		switch vv := v.(type) {
		case []interface{}:
			flattened = append(flattened, Flatten(vv)...)
		case int:
			flattened = append(flattened, vv)
		}
	}

	return flattened
}