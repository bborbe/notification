// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/notification/cmd/frontend/pkg/index"
)

var _ = Describe("NotificationEntry", func() {
	It("Type returns correct type", func() {
		entry := index.NotificationEntry{}
		Expect(entry.Type()).To(Equal("notification"))
	})
})
