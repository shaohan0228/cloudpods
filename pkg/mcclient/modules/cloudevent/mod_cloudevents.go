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

package cloudevent

import (
	"yunion.io/x/onecloud/pkg/mcclient/modulebase"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
)

var (
	Cloudevents    modulebase.ResourceManager
	CloudeventLogs modulebase.ResourceManager
)

func init() {
	Cloudevents = modules.NewCloudeventManager("cloudevent", "cloudevents",
		[]string{"Action", "Service", "Success",
			"Resource_Type", "Cloudprovider_Id", "Manager", "Provider", "Domain", "Domain_Id"},
		[]string{})

	modules.Register(&Cloudevents)

	CloudeventLogs = modules.NewCloudeventManager("event", "events",
		[]string{"id", "ops_time", "obj_id", "obj_type", "obj_name", "user", "user_id", "tenant", "tenant_id", "owner_tenant_id", "action", "notes"},
		[]string{})

}
