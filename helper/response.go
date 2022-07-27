package helper

func ResponseFailed(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "error",
		"message": msg,
	}
}

func ResponseSuccesNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
	}
}

func ResponseSuccesWithData(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
		"data":    data,
	}
}
