package gonbt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {
	type jsonTestCase struct {
		name string
		nbt  Tag
		want map[string]*string
	}

	cases := []jsonTestCase{}

	for _, v := range NBTValidTestCases {
		c := jsonTestCase{
			name: v.name,
			nbt:  v.nbt,
			want: map[string]*string{
				"default": v.json,
				"compact": v.compactJson,
				"pretty":  v.prettyJson,
			},
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, *c.want["default"], Json(&c.nbt))
			assert.Equal(t, *c.want["compact"], CompactJson(&c.nbt))
			assert.Equal(t, *c.want["pretty"], PrettyJson(&c.nbt, " ", "    "))
		})
	}
}

func TestTagName_Json(t *testing.T) {
	type jsonTestCase struct {
		name string
		nbt  *TagName
		want *string
	}

	cases := []jsonTestCase{}

	for _, v := range TagNameValidTestCases {
		c := jsonTestCase{
			name: v.name,
			nbt:  v.nbt,
			want: v.json,
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, *c.want, c.nbt.Json())
		})
	}
}

func TestPayload_Json(t *testing.T) {
	type jsonTestCase struct {
		name string
		nbt  Payload
		want map[string]*string
	}

	cases := []jsonTestCase{}

	for _, v := range PayloadValidTestCases {
		c := jsonTestCase{
			name: v.name,
			nbt:  v.nbt,
			want: map[string]*string{
				"default": v.json,
				"compact": v.compactJson,
				"pretty":  v.prettyJson,
			},
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, *c.want["default"], c.nbt.Json(" ", "", 0))
			assert.Equal(t, *c.want["compact"], c.nbt.Json("", "", 0))
			assert.Equal(t, *c.want["pretty"], c.nbt.Json(" ", "    ", 0))
		})
	}
}

func TestTag_Json(t *testing.T) {
	type jsonTestCase struct {
		name string
		nbt  Tag
		want map[string]*string
	}

	cases := []jsonTestCase{}

	for _, v := range TagValidTestCases {
		c := jsonTestCase{
			name: v.name,
			nbt:  v.nbt,
			want: map[string]*string{
				"default": v.json,
				"compact": v.compactJson,
				"pretty":  v.prettyJson,
			},
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, *c.want["default"], c.nbt.Json(" ", "", 0))
			assert.Equal(t, *c.want["compact"], c.nbt.Json("", "", 0))
			assert.Equal(t, *c.want["pretty"], c.nbt.Json(" ", "    ", 0))
		})
	}
}
