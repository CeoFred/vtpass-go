package utils

func MergeHeaders(baseHeaders map[string]string, headers map[string]string) map[string]string {
	result := make(map[string]string)
	for key, value := range baseHeaders {
		result[key] = value
	}
	for key, value := range headers {
		result[key] = value
	}
	return result
}

func HeadersToMap(headers ...map[string]string) map[string]string {
	headerMap := make(map[string]string)
	for _, hdrs := range headers {
		for key, value := range hdrs {
			headerMap[key] = value
		}
	}
	return headerMap
}
