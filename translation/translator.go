package translation

import "strings"

func Translate(word string, language string) string {
	language = sanitizeInput(language)

	switch language {
	case "english":
		return "hello"
	case "german":
		return "hallo"
	case "finnish":
		return "hei"
	case "french":
		return "bonjour"
	default:
		return ""
	}
}

func sanitizeInput(s string) string {
	s = strings.ToLower(s)
	return strings.TrimSpace(s)
}
