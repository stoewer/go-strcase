// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package strcase

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
	data := map[string]string{
		"":                           "",
		"F":                          "f",
		"Foo":                        "foo",
		"FooB":                       "foo_b",
		"FooID":                      "foo_id",
		" FooBar\t":                  "foo_bar",
		"HTTPStatusCode":             "http_status_code",
		"ParseURL.DoParse":           "parse_url.do_parse",
		"Convert Space":              "convert_space",
		"Convert-dash":               "convert_dash",
		"Skip___MultipleUnderscores": "skip_multiple_underscores",
		"Skip   MultipleSpaces":      "skip_multiple_spaces",
		"Skip---MultipleDashes":      "skip_multiple_dashes",
	}

	for camel, snake := range data {
		converted := SnakeCase(camel)
		assert.Equal(t, snake, converted)
	}
}

func TestUpperSnakeCase(t *testing.T) {
	data := map[string]string{
		"":                           "",
		"F":                          "F",
		"Foo":                        "FOO",
		"FooB":                       "FOO_B",
		"FooID":                      "FOO_ID",
		" FooBar\t":                  "FOO_BAR",
		"HTTPStatusCode":             "HTTP_STATUS_CODE",
		"ParseURL.DoParse":           "PARSE_URL.DO_PARSE",
		"Convert Space":              "CONVERT_SPACE",
		"Convert-dash":               "CONVERT_DASH",
		"Skip___MultipleUnderscores": "SKIP_MULTIPLE_UNDERSCORES",
		"Skip   MultipleSpaces":      "SKIP_MULTIPLE_SPACES",
		"Skip---MultipleDashes":      "SKIP_MULTIPLE_DASHES",
	}

	for camel, snake := range data {
		converted := UpperSnakeCase(camel)
		assert.Equal(t, snake, converted)
	}
}

func TestDelimiterCaseTrailingDelimiters(t *testing.T) {
	converters := []struct {
		name      string
		convert   func(string) string
		delimiter string
		upper     bool
	}{
		{"snake", SnakeCase, "_", false},
		{"upper snake", UpperSnakeCase, "_", true},
		{"kebab", KebabCase, "-", false},
		{"upper kebab", UpperKebabCase, "-", true},
	}
	// The expected values use "_" as delimiter and lower case letters, they are
	// adjusted to the respective converter below.
	words := []struct {
		input string
		want  string
	}{
		{"Dolphin", "dolphin"},
		{"FooBar", "foo_bar"},
		{"HTTPServer", "http_server"},
		{"fooB", "foo_b"},
		{"foo1", "foo1"},
	}
	prefixes := []string{"", "_", "-", "__", "-_"}
	suffixes := []string{"_", "-", "__", "--", "_-", "-_", "_ ", " _"}

	for _, converter := range converters {
		t.Run(converter.name, func(t *testing.T) {
			adjust := func(s string) string {
				s = strings.ReplaceAll(s, "_", converter.delimiter)
				if converter.upper {
					s = strings.ToUpper(s)
				}
				return s
			}

			for _, suffix := range suffixes {
				assert.Equal(t, converter.delimiter, converter.convert(suffix), "only %q", suffix)

				for _, word := range words {
					for _, prefix := range prefixes {
						want := word.want + "_"
						if prefix != "" {
							want = "_" + want
						}
						input := prefix + word.input + suffix
						assert.Equal(t, adjust(want), converter.convert(input), "input %q", input)
					}
				}
			}
		})
	}
}
