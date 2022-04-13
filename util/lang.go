package util

func getValidLanguages() map[string]string {
	lang := make(map[string]string)
	lang["EN"] = "EN"
	return lang
}

func IsLangValid(lang string) bool {
	s := getValidLanguages()[lang]
	return s != ""
}
