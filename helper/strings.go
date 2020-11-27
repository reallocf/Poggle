package helper

import "github.com/segmentio/ksuid"

func ConvertKSUIDsToStrings(ksuids []ksuid.KSUID) []string {
	res := make([]string, len(ksuids))
	for i, ksuid := range ksuids {
		res[i] = ksuid.String()
	}
	return res
}
