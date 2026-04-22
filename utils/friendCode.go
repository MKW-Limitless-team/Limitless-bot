package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func NormalizeFriendCode(friendCode string) string {
	friendCode = strings.ReplaceAll(friendCode, "-", "")
	friendCode = strings.TrimSpace(friendCode)
	return friendCode
}

func FormatFriendCode(friendCode string) string {
	friendCode = NormalizeFriendCode(friendCode)
	if len(friendCode) != 12 {
		return friendCode
	}

	return fmt.Sprintf("%s-%s-%s", friendCode[0:4], friendCode[4:8], friendCode[8:12])
}

func FriendCodeToPID(friendCode string) (uint32, error) {
	friendCode = NormalizeFriendCode(friendCode)
	if len(friendCode) != 12 {
		return 0, fmt.Errorf("invalid friend code length")
	}

	value, err := strconv.ParseUint(friendCode, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint32(value & 0xffffffff), nil
}
