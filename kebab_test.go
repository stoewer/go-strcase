// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package strcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKebabCase(t *testing.T) {
	data := map[string]string{
		"":                           "",
		"F":                          "f",
		"Foo":                        "foo",
		"FooB":                       "foo-b",
		"FooID":                      "foo-id",
		" FooBar\t":                  "foo-bar",
		"HTTPStatusCode":             "http-status-code",
		"ParseURL.DoParse":           "parse-url.do-parse",
		"Convert Space":              "convert-space",
		"Convert-dash":               "convert-dash",
		"Skip___MultipleUnderscores": "skip-multiple-underscores",
		"Skip   MultipleSpaces":      "skip-multiple-spaces",
		"Skip---MultipleDashes":      "skip-multiple-dashes",
	}

	for camel, snake := range data {
		converted := KebabCase(camel)
		assert.Equal(t, snake, converted)
	}
}

func TestUpperKebabCase(t *testing.T) {
	data := map[string]string{
		"":                           "",
		"F":                          "F",
		"Foo":                        "FOO",
		"FooB":                       "FOO-B",
		"FooID":                      "FOO-ID",
		" FooBar\t":                  "FOO-BAR",
		"HTTPStatusCode":             "HTTP-STATUS-CODE",
		"ParseURL.DoParse":           "PARSE-URL.DO-PARSE",
		"Convert Space":              "CONVERT-SPACE",
		"Convert-dash":               "CONVERT-DASH",
		"Skip___MultipleUnderscores": "SKIP-MULTIPLE-UNDERSCORES",
		"Skip   MultipleSpaces":      "SKIP-MULTIPLE-SPACES",
		"Skip---MultipleDashes":      "SKIP-MULTIPLE-DASHES",
	}

	for camel, snake := range data {
		converted := UpperKebabCase(camel)
		assert.Equal(t, snake, converted)
	}
}
