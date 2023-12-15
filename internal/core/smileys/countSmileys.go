package smileys

import (
	"strings"
)

func Process(smileysList []string) int {
	res := 0

	for _, smileys := range smileysList {
		var flag bool
		smiley := strings.Split(smileys, "") // Split string in array

		switch {
		case len(smileys) == 2: // กรณี ไม่มีจมูก
			flag = validateEyes(smiley[0]) && validateMouth(smiley[1])
		case len(smileys) == 3: // กรณี มีจมูก
			flag = validateEyes(smiley[0]) && validateNose(smiley[1]) && validateMouth(smiley[2])
		default:
			flag = false // กรณีไม่เข้าเงื่อนไข
		}

		if flag { // count smiley
			res++
		}
	}

	return res
}

func validateEyes(str string) bool {
	if str == ":" || str == ";" {
		return true
	}
	return false
}

func validateNose(str string) bool {
	if str == "-" || str == "~" {
		return true
	}
	return false
}

func validateMouth(str string) bool {
	if str == ")" || str == "D" {
		return true
	}
	return false
}
