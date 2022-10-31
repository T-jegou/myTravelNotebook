package util

import "github.com/spf13/cast"

// SafeString safely cast to string
func SafeString(s interface{}) (ret string) {
	return cast.ToString(s)
}
