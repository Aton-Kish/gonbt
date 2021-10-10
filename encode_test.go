package gonbt

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeNBT(t *testing.T) {
	type encodeTestCase struct {
		name    string
		nbt     Tag
		want    []byte
		wantErr error
	}

	cases := []encodeTestCase{}

	for _, v := range NBTEncodingValidCases {
		c := encodeTestCase{
			name:    v.name,
			nbt:     v.nbt,
			want:    v.raw,
			wantErr: nil,
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := EncodeNBT(buf, &c.nbt)

			if c.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, c.want, buf.Bytes())
			} else {
				assert.EqualError(t, err, c.wantErr.Error())
			}
		})
	}
}

func TestTagName_Encode(t *testing.T) {
	type encodeTestCase struct {
		name    string
		nbt     *TagName
		want    []byte
		wantErr error
	}

	cases := []encodeTestCase{}

	for _, v := range TagNameEncodingValidCases {
		c := encodeTestCase{
			name:    v.name,
			nbt:     v.nbt,
			want:    v.raw,
			wantErr: nil,
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := c.nbt.Encode(buf)

			if c.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, c.want, buf.Bytes())
			} else {
				assert.EqualError(t, err, c.wantErr.Error())
			}
		})
	}
}

func TestPayload_Encode(t *testing.T) {
	type encodeTestCase struct {
		name    string
		nbt     Payload
		want    []byte
		wantErr error
	}

	cases := []encodeTestCase{}

	for _, v := range PayloadEncodingValidCases {
		c := encodeTestCase{
			name:    v.name,
			nbt:     v.nbt,
			want:    v.raw,
			wantErr: nil,
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := c.nbt.Encode(buf)

			if c.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, c.want, buf.Bytes())
			} else {
				assert.EqualError(t, err, c.wantErr.Error())
			}
		})
	}
}

func TestTag_Encode(t *testing.T) {
	type encodeTestCase struct {
		name    string
		nbt     Tag
		want    []byte
		wantErr error
	}

	cases := []encodeTestCase{}

	for _, v := range TagEncodingValidCases {
		c := encodeTestCase{
			name:    v.name,
			nbt:     v.nbt,
			want:    v.raw,
			wantErr: nil,
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := c.nbt.Encode(buf)

			if c.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, c.want, buf.Bytes()[1:])
			} else {
				assert.EqualError(t, err, c.wantErr.Error())
			}
		})
	}
}
