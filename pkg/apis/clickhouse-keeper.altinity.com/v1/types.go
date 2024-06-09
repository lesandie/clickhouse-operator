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

package v1

import (
	"sync"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	apiChi "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClickHouseKeeperInstallation defines a ClickHouse Keeper ChkCluster
type ClickHouseKeeperInstallation struct {
	meta.TypeMeta   `json:",inline"                     yaml:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"          yaml:"metadata,omitempty"`

	Spec   ChkSpec    `json:"spec"             yaml:"spec"`
	Status *ChkStatus `json:"status,omitempty" yaml:"status,omitempty"`

	runtime             *ClickHouseKeeperInstallationRuntime `json:"-" yaml:"-"`
	statusCreatorMutex  sync.Mutex                           `json:"-" yaml:"-"`
	runtimeCreatorMutex sync.Mutex                           `json:"-" yaml:"-"`
}

type ClickHouseKeeperInstallationRuntime struct {
}

func newClickHouseKeeperInstallationRuntime() *ClickHouseKeeperInstallationRuntime {
	return &ClickHouseKeeperInstallationRuntime{}
}

func (runtime *ClickHouseKeeperInstallationRuntime) GetAttributes() *apiChi.ComparableAttributes {
	return nil
}

func (runtime *ClickHouseKeeperInstallationRuntime) LockCommonConfig() {
}

func (runtime *ClickHouseKeeperInstallationRuntime) UnlockCommonConfig() {
}

// ChkSpec defines spec section of ClickHouseKeeper resource
type ChkSpec struct {
	Configuration *ChkConfiguration `json:"configuration,omitempty"          yaml:"configuration,omitempty"`
	Templates     *apiChi.Templates `json:"templates,omitempty"              yaml:"templates,omitempty"`
}

// ChkConfiguration defines configuration section of .spec
type ChkConfiguration struct {
	Settings *apiChi.Settings `json:"settings,omitempty"  yaml:"settings,omitempty"`
	Clusters []*ChkCluster    `json:"clusters,omitempty"  yaml:"clusters,omitempty"`
}

// ChkCluster defines item of a clusters section of .configuration
type ChkCluster struct {
	Name   string                `json:"name,omitempty"         yaml:"name,omitempty"`
	Layout *apiChi.ClusterLayout `json:"layout,omitempty"       yaml:"layout,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClickHouseKeeperList defines a list of ClickHouseKeeper resources
type ClickHouseKeeperInstallationList struct {
	meta.TypeMeta `json:",inline"  yaml:",inline"`
	meta.ListMeta `json:"metadata" yaml:"metadata"`
	Items         []ClickHouseKeeperInstallation `json:"items" yaml:"items"`
}
