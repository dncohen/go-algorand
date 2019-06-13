// Copyright (C) 2019 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"golang.org/x/sys/unix"

	"github.com/algorand/go-algorand/logging"
)

const mlockallFlags = unix.MCL_CURRENT | unix.MCL_FUTURE | unix.MCL_ONFAULT

func tryMlockall(log logging.Logger) {
	err := unix.Mlockall(mlockallFlags)
	if err != nil {
		log.Infof("failed to mlockall: %s", err)
	} else {
		log.Infof("successfully executed mlockall")
	}
}