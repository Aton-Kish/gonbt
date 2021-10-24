package gonbt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	type parseTestCase struct {
		name string
		snbt *string
		want Tag
	}

	cases := []parseTestCase{}

	for _, v := range NBTValidTestCases {
		c := parseTestCase{
			name: v.name,
			snbt: v.compactSnbt,
			want: v.nbt,
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			stbm := NewSnbtTokenBitmaps(*c.snbt)
			stbm.SetTokenBitmaps()
			stbm.SetMaskBitmaps()

			tag := new(Tag)
			err := Parse(&stbm, tag)

			assert.NoError(t, err)
			assert.Equal(t, Stringify(&c.want), Stringify(tag))
		})
	}
}

func TestPayload_Parse(t *testing.T) {
	type parseTestCase struct {
		name    string
		snbt    *string
		want    Payload
		wantErr error
	}

	cases := []parseTestCase{}

	for _, v := range PayloadValidTestCases {
		c := parseTestCase{
			name:    v.name,
			snbt:    v.compactSnbt,
			want:    v.nbt,
			wantErr: nil,
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			stbm := NewSnbtTokenBitmaps(*c.snbt)
			stbm.SetTokenBitmaps()
			stbm.SetMaskBitmaps()
			stbm.NextToken(``, `" `)

			payload, err := NewPayloadFromSnbt(&stbm)
			assert.NoError(t, err)

			err = payload.Parse(&stbm)

			if c.wantErr == nil {
				assert.NoError(t, err)
				if assert.ObjectsAreEqual(c.want, payload) {
					assert.Equal(t, c.want, payload)
				} else {
					assert.Equal(t, c.want.Stringify(" ", "", 0), payload.Stringify(" ", "", 0))
				}
			} else {
				assert.EqualError(t, err, c.wantErr.Error())
			}
		})
	}
}

func TestTag_Parse(t *testing.T) {
	type parseTestCase struct {
		name    string
		snbt    *string
		want    Tag
		wantErr error
	}

	cases := []parseTestCase{}

	for _, v := range TagValidTestCases {
		c := parseTestCase{
			name:    v.name,
			snbt:    v.compactSnbt,
			want:    v.nbt,
			wantErr: nil,
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			stbm := NewSnbtTokenBitmaps(*c.snbt)
			stbm.SetTokenBitmaps()
			stbm.SetMaskBitmaps()

			var tag Tag
			var err error
			if c.want.TypeId() == TagEnd {
				tag, err = NewTag(TagEnd)
			} else {
				tag, err = NewTagFromSnbt(&stbm)
			}
			assert.NoError(t, err)

			err = tag.Parse(&stbm)

			if c.wantErr == nil {
				assert.NoError(t, err)
				if assert.ObjectsAreEqual(c.want, tag) {
					assert.Equal(t, c.want, tag)
				} else {
					assert.Equal(t, c.want.Stringify(" ", "", 0), tag.Stringify(" ", "", 0))
				}
			} else {
				assert.EqualError(t, err, c.wantErr.Error())
			}
		})
	}
}
