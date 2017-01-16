package utm

import (
	"errors"
	"net/url"
	"strings"
)

// Utm структура меток ( https://yandex.ru/support/metrika/reports/tags.xml )
type Utm struct {
	// источник перехода. Например, direct.yandex.ru, begun.ru и др.
	Source string `json:"source"`
	// средство маркетинга. Например, cpc(или ppc), banner, email и др.
	Medium string `json:"medium"`
	// название проводимой рекламной кампании
	Campaign string `json:"campaign"`
	// дополнительная информация, которая помогает различать объявления
	Content string `json:"content"`
	// ключевая фраза
	Term string `json:"term"`

	// Расширенные метки Google.Adwords и Yandex.Директ
	Extra map[string]string
}

// ParseURL парсинг utm меток из url
func ParseURL(rawurl *string) (Utm, error) {

	if !strings.Contains(*rawurl, "utm_") {
		return Utm{}, errors.New("nope")
	}

	urlObj, err := url.Parse(*rawurl)

	if err != nil {
		return Utm{}, err
	}

	if !strings.Contains(urlObj.RawQuery, "utm_") {
		return Utm{}, errors.New("nope")
	}

	u, _ := url.ParseQuery(urlObj.RawQuery)

	return ParseValuesFull(&u)
}

// ParseValues парсинг utm меток из url.Values
func ParseValues(params *url.Values) (Utm, error) {
	return parseValues(params, false)
}

// ParseValuesFull парсинг utm меток из url.Values
func ParseValuesFull(params *url.Values) (Utm, error) {
	return parseValues(params, true)
}

func parseValues(params *url.Values, extra bool) (Utm, error) {

	utm := Utm{}

	if extra {
		utm.Extra = make(map[string]string)
	}

	for name, values := range *params {

		if len(values[0]) == 0 || len(name) < 5 {
			continue
		}

		if "utm_" == name[0:4] {

			switch name {
			case "utm_source":
				utm.Source = values[0]
			case "utm_campaign":
				utm.Campaign = values[0]
			case "utm_medium":
				utm.Medium = values[0]
			case "utm_content":
				utm.Content = values[0]
			case "utm_term":
				utm.Term = values[0]
			default:
				if extra {
					utm.Extra[name] = values[0]
				}
			}
		}
	}

	return utm, nil
}
