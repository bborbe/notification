// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import (
	"fmt"
	"io"
)

type HasCause interface {
	Cause() error
}

type HasData interface {
	Data() map[string]any
}

type DataError interface {
	error
	HasData
	HasCause
}

func AddDataToError(err error, data map[string]any) DataError {
	return &dataError{
		err:  err,
		data: data,
	}
}

type dataError struct {
	err  error
	data map[string]any
}

// Unwrap provides compatibility for Go 1.13 error chains.
func (d *dataError) Unwrap() error { return d.Cause() }

func (d *dataError) Cause() error {
	return d.err
}

func (d *dataError) Error() string {
	return d.err.Error()
}

func (d *dataError) Data() map[string]any {
	return d.data
}

func (d *dataError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", d.Cause())
			return
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(s, d.Error())
	}
}

func DataFromError(err error) map[string]any {
	data := make(map[string]any)
	for err != nil {
		hasData, ok := err.(HasData)
		if ok {
			for k, v := range hasData.Data() {
				data[k] = v
			}
		}
		hasCause, ok := err.(HasCause)
		if !ok {
			break
		}
		err = hasCause.Cause()
	}
	return data
}
