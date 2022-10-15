// Copyright (c) 2022 Aton-Kish
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package snbt

import (
	"errors"
	"fmt"
)

var (
	ErrOutOfRange    = errors.New("out of range")
	ErrStopIteration = errors.New("stop iteration")
	ErrUnexpected    = errors.New("unexpected error")
)

type SnbtError struct {
	Op  string
	Err error
}

func (e *SnbtError) Error() string {
	if e == nil {
		return "<nil>"
	}

	var err string
	if e.Err == nil {
		err = "<nil>"
	} else {
		err = e.Err.Error()
	}

	return fmt.Sprintf("snbt %s: %s", e.Op, err)
}

func (e *SnbtError) Unwrap() error {
	return e.Err
}
