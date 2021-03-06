// Copyright 2016 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"github.com/coreos/ignition/v2/config/shared/errors"
	"github.com/coreos/ignition/v2/config/validate/report"
)

func (r Raid) Key() string {
	return r.Name
}

func (r Raid) IgnoreDuplicates() map[string]struct{} {
	return map[string]struct{}{
		"Options": {},
	}
}

func (n Raid) ValidateLevel() (r report.Report) {
	switch n.Level {
	case "linear", "raid0", "0", "stripe":
		if n.Spares != nil && *n.Spares != 0 {
			r.AddOnError(errors.ErrSparesUnsupportedForLevel)
		}
	case "raid1", "1", "mirror":
	case "raid4", "4":
	case "raid5", "5":
	case "raid6", "6":
	case "raid10", "10":
	default:
		r.AddOnError(errors.ErrUnrecognizedRaidLevel)
	}
	return r
}

func (n Raid) ValidateDevices() (r report.Report) {
	for _, d := range n.Devices {
		if err := validatePath(string(d)); err != nil {
			r.AddOnError(errors.ErrPathRelative)
		}
	}
	return
}
