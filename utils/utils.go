package utils

import (
	"regexp"
)

func DerefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func FormatPrice(price *string) {
	r := regexp.MustCompile(`\$(\d+(\.\d+)?).*$`)

	newPrices := r.FindStringSubmatch(*price)

	if len(newPrices) > 1 {
		*price = newPrices[1]
	} else {
		*price = "Unknown"
	}

}

func FormatStars(stars *string) {
	if len(*stars) >= 3 {
		*stars = (*stars)[0:3]
	} else {
		*stars = "Unknown"
	}
}

func FormatURL(url *string) {
	if len(*url) > 0 {
		*url = "https://www.amazon.com" + *url
	} else {
		*url = "Unknown"
	}
}
