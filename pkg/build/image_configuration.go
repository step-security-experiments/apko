// Copyright 2022 Chainguard, Inc.
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

package build

import (
	"os"

	"chainguard.dev/apko/pkg/build/types"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func LoadImageConfiguration(imageConfigPath string) (types.ImageConfiguration, error) {
	ic := types.ImageConfiguration{}

	data, err := os.ReadFile(imageConfigPath)
	if err != nil {
		return ic, errors.Wrap(err, "failed to read image configuration file")
	}

	err = yaml.Unmarshal(data, &ic)
	if err != nil {
		return ic, errors.Wrap(err, "failed to parse image configuration")
	}

	return ic, nil
}