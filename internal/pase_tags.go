package internal

import "strings"

type tag struct {
	Key   string
	Value string
}

func ParseTag(fullTag string) []tag {
	tags := []tag{}

	splitTags := strings.Split(fullTag, ",")
	for _, splitTag := range splitTags {
		individualTags := strings.Split(splitTag, "=")
		if len(individualTags) == 2 {
			tags = append(tags, tag{Key: individualTags[0], Value: individualTags[1]})
		}
	}

	return tags
}
