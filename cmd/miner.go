// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"block-chain/libs"
	"os"
)

func RunMiner() {
	libs.StartServer(os.Getenv("NODE_ID"), os.Getenv("G_ADDR"))
}
