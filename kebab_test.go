// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package strcase_test

import (
	"testing"

	"github.com/stoewer/go-strcase"
	"github.com/stretchr/testify/assert"
)

func TestKebabCase(t *testing.T) {
	data := map[string]string{
		"":                           "",
		"F":                          "f",
		"Foo":                        "foo",
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
		converted := strcase.KebabCase(camel)
		assert.Equal(t, snake, converted)
	}
}
