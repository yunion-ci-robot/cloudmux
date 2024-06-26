// Copyright 2019 Yunion
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

package shell

import (
	"yunion.io/x/cloudmux/pkg/multicloud/ksyun"

	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/shellutils"
)

func init() {
	type TagListOptions struct {
		ResourceType string `choices:"kec-image|kec-instance|rds-instance|security-group|vpc|subnet|eip|slb|bucket|kcs-instance|volume|nat|tunnel|bws|peering" default:"kec-instance"`
		ID           string
	}
	shellutils.R(&TagListOptions{}, "tag-list", "list tags", func(cli *ksyun.SRegion, args *TagListOptions) error {
		ret, err := cli.ListTags(args.ResourceType, args.ID)
		if err != nil {
			return errors.Wrap(err, "GetTags")
		}
		printObject(ret)
		return nil
	})
}
