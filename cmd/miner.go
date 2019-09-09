// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"block-chain/config"
	"block-chain/libs"
)

func RunMiner() {
	libs.StartServer(config.NodeID, config.Gaddress)
}
