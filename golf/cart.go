package golf

import (
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"syscall/js"
)

// Dset saves the data to the broser with the given name
func (e *Engine) Dset(name string, data []byte) error {
	if len(data) > 1024 {
		return errors.New("only 1024 bytes or less can be saved")
	}
	if len(name) > 1024 {
		return errors.New("cart data name must be 1024 characters or shorter")
	}

	match := regexp.MustCompile(`[a-zA-Z0-9_]+`).FindString(name)
	if len(match) != len(name) {
		return errors.New("cart data name can only contain alpha numeric characters or the underscore character")
	}

	save := base64.StdEncoding.EncodeToString(data)
	cookie := fmt.Sprintf("%s=%s;", name, save)
	js.Global().Get("document").Set("cookie", cookie)
	return nil
}

// Dget retrives the named data from the browser
func (e *Engine) Dget(name string) ([]byte, bool) {
	cookies := js.Global().Get("document").Get("cookie").String()
	allCookies := strings.Split(cookies, ";")
	for _, c := range allCookies {
		data := strings.Split(c, "=")
		n, d := strings.Trim(data[0], " "), strings.Trim(data[1], " ")
		if n == name {
			ret, err := base64.StdEncoding.DecodeString(d)
			if err != nil {
				return []byte{}, false
			}
			return ret, true
		}
	}
	return []byte{}, false
}
