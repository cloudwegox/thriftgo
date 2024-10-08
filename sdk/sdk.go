// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sdk

import (
	"github.com/cloudwegox/thriftgo/utils/dir_utils"

	"github.com/cloudwegox/thriftgo/config"

	"github.com/cloudwegox/thriftgo/plugin"
)

func RunThriftgoAsSDK(wd string, SDKPlugins []plugin.SDKPlugin, args ...string) error {

	// this should execute at the first line!
	dir_utils.SetGlobalwd(wd)

	// for sdk mode, every time when function is invoked, config file should be reload
	err := config.LoadConfig()
	if err != nil {
		return err
	}

	return InvokeThriftgo(SDKPlugins, append([]string{"thriftgo"}, args...)...)
}
