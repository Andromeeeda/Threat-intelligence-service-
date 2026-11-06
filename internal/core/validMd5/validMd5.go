package core_validMd5

import "regexp"

func IsValidMD5(md5 string) bool {
	matched, _ := regexp.MatchString(`^[a-fA-F0-9]{32}$`, md5)
	return matched
}

