package Debug

import "log"

func CheckErr(err error) bool {
	if err == nil {
		return true
	} else {
		log.Print(err)
		return false
	}
}
