// В этом пакете вы реализуете функцию автоматического определения кода Морзе или обычного текста из переданной строки.
// Если передан обычный текст,функция должна переконвертировать его в код Морзе и вернуть;
//  и наоборот — если был передан код Морзе, функция должна переконвертировать его в обычный текст и вернуть.
// Для реализации этой функции придётся обратиться к стандартной библиотеке, а именно — к пакету strings.
// В этом пакете есть хорошие примеры, которые демонстрируют, как можно решить эту задачу.
// Не забудьте обработать ошибки и вернуть их.

package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoConvertor(reqStr string) (string, error) {

	if len(reqStr) <= 0 {
		err := errors.New("an empty string is not allowed.")
		return "", err
	}

	processedStr := strings.ReplaceAll(reqStr, ".", "")
	processedStr = strings.ReplaceAll(processedStr, "-", "")
	processedStr = strings.ReplaceAll(processedStr, " ", "")

	if len(processedStr) > 0 {
		return morse.ToMorse(reqStr), nil
	}

	return morse.ToText(reqStr), nil

}
