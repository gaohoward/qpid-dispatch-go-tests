// Copyright 2019 The Interconnectedcloud Authors
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

package framework

import (
	v1alpha1 "github.com/interconnectedcloud/qdr-operator/pkg/apis/interconnectedcloud/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InterconnectCustomizer represents a function that allows for
// customizing an Interconnect resource before it is created.
type InterconnectCustomizer func(interconnect *v1alpha1.Interconnect)

// CreateInterconnect creates an interconnect resource
func (c *ContextData) CreateInterconnect(namespace string, size int32, fn ...InterconnectCustomizer) (*v1alpha1.Interconnect, error) {

	obj := &v1alpha1.Interconnect{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Interconnect",
			APIVersion: "interconnectedcloud.github.io/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.UniqueName,
			Namespace: namespace,
		},
		Spec: v1alpha1.InterconnectSpec{
			DeploymentPlan: v1alpha1.DeploymentPlanType{
				Size:      size,
				Image:     TestContext.QdrImage,
				Role:      "interior",
				Placement: "Any",
			},
		},
	}

	// Customize the interconnect resource before creation
	for _, f := range fn {
		f(obj)
	}
	// create the interconnect resource
	return c.Clients.QdrClient.InterconnectedcloudV1alpha1().Interconnects(c.Namespace).Create(obj)
}

func (c *ContextData) DeleteInterconnect(interconnect *v1alpha1.Interconnect) error {
	return c.Clients.QdrClient.InterconnectedcloudV1alpha1().Interconnects(c.Namespace).Delete(interconnect.Name, &metav1.DeleteOptions{})
}

func (c *ContextData) GetInterconnect(name string) (*v1alpha1.Interconnect, error) {
	return c.Clients.QdrClient.InterconnectedcloudV1alpha1().Interconnects(c.Namespace).Get(name, metav1.GetOptions{})
}

func (c *ContextData) UpdateInterconnect(interconnect *v1alpha1.Interconnect) (*v1alpha1.Interconnect, error) {
	return c.Clients.QdrClient.InterconnectedcloudV1alpha1().Interconnects(c.Namespace).Update(interconnect)
}
