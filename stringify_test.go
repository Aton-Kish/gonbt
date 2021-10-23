package gonbt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringify(t *testing.T) {
	type stringifyTestCase struct {
		name string
		nbt  Tag
		want map[string]*string
	}

	cases := []stringifyTestCase{}

	for _, v := range NBTValidTestCases {
		c := stringifyTestCase{
			name: v.name,
			nbt:  v.nbt,
			want: map[string]*string{
				"default": v.snbt,
				"compact": v.compactSnbt,
				"pretty":  v.prettySnbt,
			},
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, *c.want["default"], Stringify(&c.nbt))
			assert.Equal(t, *c.want["compact"], CompactStringify(&c.nbt))
			assert.Equal(t, *c.want["pretty"], PrettyStringify(&c.nbt, " ", "    "))
		})
	}
}

func TestTagName_Stringify(t *testing.T) {
	type stringifyTestCase struct {
		name string
		nbt  *TagName
		want *string
	}

	cases := []stringifyTestCase{}

	for _, v := range TagNameValidTestCases {
		c := stringifyTestCase{
			name: v.name,
			nbt:  v.nbt,
			want: v.snbt,
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, *c.want, c.nbt.Stringify())
		})
	}
}

func TestPayload_Stringify(t *testing.T) {
	type stringifyTestCase struct {
		name string
		nbt  Payload
		want map[string]*string
	}

	cases := []stringifyTestCase{}

	for _, v := range PayloadValidTestCases {
		c := stringifyTestCase{
			name: v.name,
			nbt:  v.nbt,
			want: map[string]*string{
				"default": v.snbt,
				"compact": v.compactSnbt,
				"pretty":  v.prettySnbt,
			},
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, *c.want["default"], c.nbt.Stringify(" ", "", 0))
			assert.Equal(t, *c.want["compact"], c.nbt.Stringify("", "", 0))
			assert.Equal(t, *c.want["pretty"], c.nbt.Stringify(" ", "    ", 0))
		})
	}
}

func TestTag_Stringify(t *testing.T) {
	type stringifyTestCase struct {
		name string
		nbt  Tag
		want map[string]*string
	}

	cases := []stringifyTestCase{}

	for _, v := range TagValidTestCases {
		c := stringifyTestCase{
			name: v.name,
			nbt:  v.nbt,
			want: map[string]*string{
				"default": v.snbt,
				"compact": v.compactSnbt,
				"pretty":  v.prettySnbt,
			},
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, *c.want["default"], c.nbt.Stringify(" ", "", 0))
			assert.Equal(t, *c.want["compact"], c.nbt.Stringify("", "", 0))
			assert.Equal(t, *c.want["pretty"], c.nbt.Stringify(" ", "    ", 0))
		})
	}
}
