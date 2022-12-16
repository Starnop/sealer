// Copyright © 2021 Alibaba Group Holding Ltd.
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

package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sealerio/sealer/version"
)

var shortPrint bool

func NewVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:     "version",
		Short:   "show sealer and related versions",
		Args:    cobra.NoArgs,
		Example: `sealer version`,
		RunE: func(cmd *cobra.Command, args []string) error {
			marshalled, err := json.Marshal(version.Get())
			if err != nil {
				return err
			}
			if shortPrint {
				fmt.Println(version.Get().String())
			} else {
				fmt.Println(string(marshalled))
			}
			return nil
		},
	}
	versionCmd.Flags().BoolVar(&shortPrint, "short", false, "if true, print sealer's own version number.")
	return versionCmd
}
