/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudup

import (
	"fmt"
	api "k8s.io/kops/pkg/apis/kops"
	"strings"
	"testing"
)

func buildMinimalNodeInstanceGroup(zones ...string) *api.InstanceGroup {
	g := &api.InstanceGroup{}
	g.Name = "nodes"
	g.Spec.Role = api.InstanceGroupRoleNode
	g.Spec.Zones = zones

	return g
}

func buildMinimalMasterInstanceGroup(zones ...string) *api.InstanceGroup {
	g := &api.InstanceGroup{}
	g.Name = "master"
	g.Spec.Role = api.InstanceGroupRoleMaster
	g.Spec.Zones = zones

	return g
}

func TestPopulateInstanceGroup_Name_Required(t *testing.T) {
	cluster := buildMinimalCluster()
	g := buildMinimalNodeInstanceGroup()
	g.Name = ""

	channel := &api.Channel{}

	expectErrorFromPopulateInstanceGroup(t, cluster, g, channel, "Name")
}

func TestPopulateInstanceGroup_Role_Required(t *testing.T) {
	cluster := buildMinimalCluster()
	g := buildMinimalNodeInstanceGroup()
	g.Spec.Role = ""

	channel := &api.Channel{}

	expectErrorFromPopulateInstanceGroup(t, cluster, g, channel, "Role")
}

func expectErrorFromPopulateInstanceGroup(t *testing.T, cluster *api.Cluster, g *api.InstanceGroup, channel *api.Channel, message string) {
	_, err := PopulateInstanceGroupSpec(cluster, g, channel)
	if err == nil {
		t.Fatalf("Expected error from PopulateInstanceGroup")
	}
	actualMessage := fmt.Sprintf("%v", err)
	if !strings.Contains(actualMessage, message) {
		t.Fatalf("Expected error %q, got %q", message, actualMessage)
	}
}
