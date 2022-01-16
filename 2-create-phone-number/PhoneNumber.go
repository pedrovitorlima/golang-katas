package phone_number

import "strconv"

func CreatePhoneNumber(numbers [10]uint) string {
	phone := "(" + substring(numbers[:3]) + ") " + substring(numbers[3:6]) + "-" + substring(numbers[6:])
	return phone
}

func substring(array []uint) string {
	var phrase = ""
	for _, number := range array {
		phrase = phrase + strconv.FormatUint(uint64(number), 10)
	}

	return phrase
}
