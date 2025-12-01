package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {

	cache := NewCache(1 * time.Second)

	cases := []struct {
		input struct {
			key string
			val []byte
		}
		expected struct {
			key string
			val []byte
		}
	}{
		{
			input: struct {
				key string
				val []byte
			}{key: "x", val: []byte("x")},
			expected: struct {
				key string
				val []byte
			}{key: "x", val: []byte("x")},
		},
		{
			input: struct {
				key string
				val []byte
			}{key: "", val: []byte("empty")},
			expected: struct {
				key string
				val []byte
			}{key: "", val: []byte("empty")},
		},
		{
			input: struct {
				key string
				val []byte
			}{key: "#$%%%%", val: []byte("some message")},
			expected: struct {
				key string
				val []byte
			}{key: "#$%%%%", val: []byte("some message")},
		},
	}

	for _, testCase := range cases {
		cache.Add(testCase.input.key, testCase.input.val)
		acctual, _ := cache.Get(testCase.input.key)

		if string(testCase.expected.val) != string(acctual) {
			t.Errorf("input %s does not match cache value %s", testCase.input, string(acctual))
		}

		time.Sleep(2 * time.Second)
		_, exists := cache.Get("a")
		if exists {
			t.Error("Entry should not exist")
		}

	}

}
