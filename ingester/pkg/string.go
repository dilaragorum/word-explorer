package pkg

func StringSliceToMap(values []string) map[string]struct{} {
	result := make(map[string]struct{}, len(values))

	for i := range values {
		result[values[i]] = struct{}{}
	}

	return result
}
