package entity

type Lang string

const (
	LangJa Lang = "japanese"
	LangEn Lang = "english"
)

func GetMatchLang(lang string) Lang {
	switch lang {
	case "ja":
		return LangJa
	case "en":
		return LangEn
	default:
		return LangJa
	}
}
