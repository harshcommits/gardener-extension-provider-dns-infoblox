// Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	componentbaseconfigv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerConfiguration defines the configuration for the GCP provider.
type ControllerConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	// ClientConnection specifies the kubeconfig file and client connection
	// settings for the proxy server to use when communicating with the apiserver.
	// +optional
	ClientConnection *componentbaseconfigv1alpha1.ClientConnectionConfiguration `json:"clientConnection,omitempty"`
}

// ProviderConfigManager contains configurations settings for the providerconfig.
type ProviderConfigManager struct {
	metav1.TypeMeta `json:",inline"`
	// Hosts contains information about the Host ip
	Host *string `json:"host,omitempty"`

	// View Contains Information about the view parameter fo infoblox config used to pass to infoblox dns client
	View *string `json:"view,omitempty"`

	//Port            *int    `json:"port,omitempty"`
	//SSLVerify       *bool   `json:"sslVerify,omitempty"`
	//Version         *string `json:"version,omitempty"`
	//PoolConnections *int    `json:"httpPoolConnections,omitempty"`
	//RequestTimeout  *int    `json:"httpRequestTimeout,omitempty"`
	//CaCert          *string `json:"caCert,omitempty"`
	//MaxResults      int     `json:"maxResults,omitempty"`
	//ProxyURL        *string `json:"proxyUrl,omitempty"`
}
