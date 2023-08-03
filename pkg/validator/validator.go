package validator

import (
	"fmt"
	"strings"
)

func IsValidString(text string) error {

	if len(text) == 0 {
		return fmt.Errorf("необходимо ввести значение")
	}

	if strings.HasPrefix(text, " ") {
		return fmt.Errorf("текст не может начинаться с пробела")
	}

	return nil
}
