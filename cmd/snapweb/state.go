/*
 * Copyright (C) 2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"github.com/snapcore/snapd/client"
)

// IsDeviceManaged determines if the device is in the 'managed' state
func IsDeviceManaged() bool {
	client := client.New(nil)

	sysInfo, err := client.SysInfo()
	if err != nil {
		panic(err)
	}

	if sysInfo.OnClassic || sysInfo.Managed {
		return true
	}

	return false
}
