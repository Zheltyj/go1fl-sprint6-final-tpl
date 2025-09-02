package service

import "github.com/Zheltyj/go1fl-sprint6-final-tpl/pkg/morse"

func ConvertString(stringToConvert string) string {
	var isMorse bool
	var convertedString string

	for _, v := range stringToConvert {
		if v == 45 || v == 46 || v == 32 {
			isMorse = true
		} else {
			isMorse = false
			break
		}
	}
	if isMorse {

		convertedString = morse.ToText(stringToConvert)
	}
	if !isMorse {

		convertedString = morse.ToMorse(stringToConvert)
	}
	return convertedString
}
