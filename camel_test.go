// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package strcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpperCamelCase(t *testing.T) {
	data := map[string]string{
		"":                      "",
		"f":                     "F",
		"foo":                   "Foo",
		"fooBar":                "FooBar",
		"FooBarBla":             "FooBarBla",
		"foo_barBla":            "FooBarBla",
		" foo_bar\n":            "FooBar",
		" foo-bar\t":            "FooBar",
		" foo bar\r":            "FooBar",
		"HTTP_status_code":      "HttpStatusCode",
		"skip   many spaces":    "SkipManySpaces",
		"skip---many-dashes":    "SkipManyDashes",
		"skip___many_underline": "SkipManyUnderline",
	}

	for in, out := range data {
		converted := UpperCamelCase(in)
		assert.Equal(t, out, converted)
	}
}

func TestLowerCamelCase(t *testing.T) {
	data := map[string]string{
		"":                      "",
		"F":                     "f",
		"foo":                   "foo",
		"FooBar":                "fooBar",
		"fooBarBla":             "fooBarBla",
		"foo_barBla":            "fooBarBla",
		" foo_bar\n":            "fooBar",
		" foo-bar\t":            "fooBar",
		" foo bar\r":            "fooBar",
		"HTTP_status_code":      "httpStatusCode",
		"skip   many spaces":    "skipManySpaces",
		"skip---many-dashes":    "skipManyDashes",
		"skip___many_underline": "skipManyUnderline",
	}

	for in, out := range data {
		converted := LowerCamelCase(in)
		assert.Equal(t, out, converted)
	}
}
