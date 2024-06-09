// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
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

package managers

import (
	"github.com/altinity/clickhouse-operator/pkg/interfaces"
	chiVolume "github.com/altinity/clickhouse-operator/pkg/model/chi/volume"
	chkVolume "github.com/altinity/clickhouse-operator/pkg/model/chk/volume"
)

type VolumeManagerType string

const (
	VolumeManagerTypeClickHouse VolumeManagerType = "clickhouse"
	VolumeManagerTypeKeeper     VolumeManagerType = "keeper"
)

func NewVolumeManager(what VolumeManagerType) interfaces.IVolumeManager {
	switch what {
	case VolumeManagerTypeClickHouse:
		return chiVolume.NewVolumeManagerClickHouse()
	case VolumeManagerTypeKeeper:
		return chkVolume.NewVolumeManagerKeeper()
	}
	panic("unknown volume manager type")
}
