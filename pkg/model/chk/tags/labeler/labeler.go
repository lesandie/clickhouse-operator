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

package labeler

import (
	api "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse-keeper.altinity.com/v1"
	apiChi "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	"github.com/altinity/clickhouse-operator/pkg/interfaces"
	"github.com/altinity/clickhouse-operator/pkg/model/common/tags/labeler"
	"github.com/altinity/clickhouse-operator/pkg/util"
)

// Labeler is an entity which can label CHI artifacts
type Keeper struct {
	*labeler.Labeler
}

// NewLabelerKeeper creates new labeler with context
func NewLabelerKeeper(cr apiChi.ICustomResource, config labeler.Config) *Keeper {
	return &Keeper{
		Labeler: labeler.NewLabeler(cr, config),
	}
}

func (l *Keeper) Label(what interfaces.LabelType, params ...any) map[string]string {
	switch what {
	case interfaces.LabelConfigMapCommon:
		return l.labelConfigMapCHICommon()

	default:
		return l.Labeler.Label(what, params...)
	}
	panic("unknown label type")
}

func (l *Keeper) Selector(what interfaces.SelectorType, params ...any) map[string]string {
	return l.Labeler.Selector(what, params...)
}

// labelConfigMapCHICommon
func (l *Keeper) labelConfigMapCHICommon() map[string]string {
	return util.MergeStringMapsOverwrite(
		l.GetCRScope(),
		map[string]string{
			labeler.LabelConfigMap: labeler.LabelConfigMapValueCHICommon,
		})
}

func GetPodLabels(chk *api.ClickHouseKeeperInstallation) map[string]string {
	// In case Pod template has labels explicitly specified - use them
	//labels := chk2.getPodTemplateLabels(chk)
	//if labels != nil {
	//	return labels
	//}

	// Either no pod template or labels specified.
	// Construct default labels
	return map[string]string{
		"app": chk.GetName(),
		"uid": string(chk.UID),
	}
}
