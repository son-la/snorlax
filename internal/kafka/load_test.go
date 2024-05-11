package kafka

import (
	"testing"
)

func TestIsTopicNameValid(t *testing.T) {
	topicNames := map[string]bool{
		"pedroTheRaccoon": true,
		"pedro123":        true,
		"Pedro@":          false,
		"Pedro...":        true,
		"":                false,
	}

	for name, want := range topicNames {
		got := isTopicNameValid(name)
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}

}
