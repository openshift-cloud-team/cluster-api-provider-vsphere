//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AviLoadBalancerConfig) DeepCopyInto(out *AviLoadBalancerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AviLoadBalancerConfig.
func (in *AviLoadBalancerConfig) DeepCopy() *AviLoadBalancerConfig {
	if in == nil {
		return nil
	}
	out := new(AviLoadBalancerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AviLoadBalancerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AviLoadBalancerConfigList) DeepCopyInto(out *AviLoadBalancerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AviLoadBalancerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AviLoadBalancerConfigList.
func (in *AviLoadBalancerConfigList) DeepCopy() *AviLoadBalancerConfigList {
	if in == nil {
		return nil
	}
	out := new(AviLoadBalancerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AviLoadBalancerConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AviLoadBalancerConfigSpec) DeepCopyInto(out *AviLoadBalancerConfigSpec) {
	*out = *in
	if in.AdvancedL4 != nil {
		in, out := &in.AdvancedL4, &out.AdvancedL4
		*out = new(bool)
		**out = **in
	}
	out.CredentialSecretRef = in.CredentialSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AviLoadBalancerConfigSpec.
func (in *AviLoadBalancerConfigSpec) DeepCopy() *AviLoadBalancerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AviLoadBalancerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AviLoadBalancerConfigStatus) DeepCopyInto(out *AviLoadBalancerConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AviLoadBalancerConfigStatus.
func (in *AviLoadBalancerConfigStatus) DeepCopy() *AviLoadBalancerConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AviLoadBalancerConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientSecretReference) DeepCopyInto(out *ClientSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientSecretReference.
func (in *ClientSecretReference) DeepCopy() *ClientSecretReference {
	if in == nil {
		return nil
	}
	out := new(ClientSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HAProxyLoadBalancerConfig) DeepCopyInto(out *HAProxyLoadBalancerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HAProxyLoadBalancerConfig.
func (in *HAProxyLoadBalancerConfig) DeepCopy() *HAProxyLoadBalancerConfig {
	if in == nil {
		return nil
	}
	out := new(HAProxyLoadBalancerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HAProxyLoadBalancerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HAProxyLoadBalancerConfigList) DeepCopyInto(out *HAProxyLoadBalancerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HAProxyLoadBalancerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HAProxyLoadBalancerConfigList.
func (in *HAProxyLoadBalancerConfigList) DeepCopy() *HAProxyLoadBalancerConfigList {
	if in == nil {
		return nil
	}
	out := new(HAProxyLoadBalancerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HAProxyLoadBalancerConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HAProxyLoadBalancerConfigSpec) DeepCopyInto(out *HAProxyLoadBalancerConfigSpec) {
	*out = *in
	if in.EndPointURLs != nil {
		in, out := &in.EndPointURLs, &out.EndPointURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.CredentialSecretRef = in.CredentialSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HAProxyLoadBalancerConfigSpec.
func (in *HAProxyLoadBalancerConfigSpec) DeepCopy() *HAProxyLoadBalancerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(HAProxyLoadBalancerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HAProxyLoadBalancerConfigStatus) DeepCopyInto(out *HAProxyLoadBalancerConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HAProxyLoadBalancerConfigStatus.
func (in *HAProxyLoadBalancerConfigStatus) DeepCopy() *HAProxyLoadBalancerConfigStatus {
	if in == nil {
		return nil
	}
	out := new(HAProxyLoadBalancerConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPConfig) DeepCopyInto(out *IPConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPConfig.
func (in *IPConfig) DeepCopy() *IPConfig {
	if in == nil {
		return nil
	}
	out := new(IPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPool) DeepCopyInto(out *IPPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPool.
func (in *IPPool) DeepCopy() *IPPool {
	if in == nil {
		return nil
	}
	out := new(IPPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolCondition) DeepCopyInto(out *IPPoolCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolCondition.
func (in *IPPoolCondition) DeepCopy() *IPPoolCondition {
	if in == nil {
		return nil
	}
	out := new(IPPoolCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolList) DeepCopyInto(out *IPPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolList.
func (in *IPPoolList) DeepCopy() *IPPoolList {
	if in == nil {
		return nil
	}
	out := new(IPPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolReference) DeepCopyInto(out *IPPoolReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolReference.
func (in *IPPoolReference) DeepCopy() *IPPoolReference {
	if in == nil {
		return nil
	}
	out := new(IPPoolReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolSpec) DeepCopyInto(out *IPPoolSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolSpec.
func (in *IPPoolSpec) DeepCopy() *IPPoolSpec {
	if in == nil {
		return nil
	}
	out := new(IPPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolStatus) DeepCopyInto(out *IPPoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]IPPoolCondition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolStatus.
func (in *IPPoolStatus) DeepCopy() *IPPoolStatus {
	if in == nil {
		return nil
	}
	out := new(IPPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerConfig) DeepCopyInto(out *LoadBalancerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerConfig.
func (in *LoadBalancerConfig) DeepCopy() *LoadBalancerConfig {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerConfigCondition) DeepCopyInto(out *LoadBalancerConfigCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerConfigCondition.
func (in *LoadBalancerConfigCondition) DeepCopy() *LoadBalancerConfigCondition {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerConfigCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerConfigList) DeepCopyInto(out *LoadBalancerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerConfigList.
func (in *LoadBalancerConfigList) DeepCopy() *LoadBalancerConfigList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerConfigProviderReference) DeepCopyInto(out *LoadBalancerConfigProviderReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerConfigProviderReference.
func (in *LoadBalancerConfigProviderReference) DeepCopy() *LoadBalancerConfigProviderReference {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerConfigProviderReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerConfigSpec) DeepCopyInto(out *LoadBalancerConfigSpec) {
	*out = *in
	out.ProviderRef = in.ProviderRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerConfigSpec.
func (in *LoadBalancerConfigSpec) DeepCopy() *LoadBalancerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerConfigStatus) DeepCopyInto(out *LoadBalancerConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]LoadBalancerConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerConfigStatus.
func (in *LoadBalancerConfigStatus) DeepCopy() *LoadBalancerConfigStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Network) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterface) DeepCopyInto(out *NetworkInterface) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterface.
func (in *NetworkInterface) DeepCopy() *NetworkInterface {
	if in == nil {
		return nil
	}
	out := new(NetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkInterface) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceCondition) DeepCopyInto(out *NetworkInterfaceCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceCondition.
func (in *NetworkInterfaceCondition) DeepCopy() *NetworkInterfaceCondition {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceList) DeepCopyInto(out *NetworkInterfaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceList.
func (in *NetworkInterfaceList) DeepCopy() *NetworkInterfaceList {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkInterfaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfacePortAllocation) DeepCopyInto(out *NetworkInterfacePortAllocation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfacePortAllocation.
func (in *NetworkInterfacePortAllocation) DeepCopy() *NetworkInterfacePortAllocation {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfacePortAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceProviderReference) DeepCopyInto(out *NetworkInterfaceProviderReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceProviderReference.
func (in *NetworkInterfaceProviderReference) DeepCopy() *NetworkInterfaceProviderReference {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceProviderReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceSpec) DeepCopyInto(out *NetworkInterfaceSpec) {
	*out = *in
	if in.ProviderRef != nil {
		in, out := &in.ProviderRef, &out.ProviderRef
		*out = new(NetworkInterfaceProviderReference)
		**out = **in
	}
	if in.PortAllocation != nil {
		in, out := &in.PortAllocation, &out.PortAllocation
		*out = new(NetworkInterfacePortAllocation)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceSpec.
func (in *NetworkInterfaceSpec) DeepCopy() *NetworkInterfaceSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceStatus) DeepCopyInto(out *NetworkInterfaceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]NetworkInterfaceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IPConfigs != nil {
		in, out := &in.IPConfigs, &out.IPConfigs
		*out = make([]IPConfig, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceStatus.
func (in *NetworkInterfaceStatus) DeepCopy() *NetworkInterfaceStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkList) DeepCopyInto(out *NetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Network, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkList.
func (in *NetworkList) DeepCopy() *NetworkList {
	if in == nil {
		return nil
	}
	out := new(NetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkProviderReference) DeepCopyInto(out *NetworkProviderReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkProviderReference.
func (in *NetworkProviderReference) DeepCopy() *NetworkProviderReference {
	if in == nil {
		return nil
	}
	out := new(NetworkProviderReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DNSSearchDomains != nil {
		in, out := &in.DNSSearchDomains, &out.DNSSearchDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NTP != nil {
		in, out := &in.NTP, &out.NTP
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkStatus) DeepCopyInto(out *NetworkStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkStatus.
func (in *NetworkStatus) DeepCopy() *NetworkStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMXNET3NetworkInterface) DeepCopyInto(out *VMXNET3NetworkInterface) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMXNET3NetworkInterface.
func (in *VMXNET3NetworkInterface) DeepCopy() *VMXNET3NetworkInterface {
	if in == nil {
		return nil
	}
	out := new(VMXNET3NetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMXNET3NetworkInterface) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMXNET3NetworkInterfaceList) DeepCopyInto(out *VMXNET3NetworkInterfaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VMXNET3NetworkInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMXNET3NetworkInterfaceList.
func (in *VMXNET3NetworkInterfaceList) DeepCopy() *VMXNET3NetworkInterfaceList {
	if in == nil {
		return nil
	}
	out := new(VMXNET3NetworkInterfaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMXNET3NetworkInterfaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMXNET3NetworkInterfaceSpec) DeepCopyInto(out *VMXNET3NetworkInterfaceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMXNET3NetworkInterfaceSpec.
func (in *VMXNET3NetworkInterfaceSpec) DeepCopy() *VMXNET3NetworkInterfaceSpec {
	if in == nil {
		return nil
	}
	out := new(VMXNET3NetworkInterfaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMXNET3NetworkInterfaceStatus) DeepCopyInto(out *VMXNET3NetworkInterfaceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMXNET3NetworkInterfaceStatus.
func (in *VMXNET3NetworkInterfaceStatus) DeepCopy() *VMXNET3NetworkInterfaceStatus {
	if in == nil {
		return nil
	}
	out := new(VMXNET3NetworkInterfaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereDistributedNetwork) DeepCopyInto(out *VSphereDistributedNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereDistributedNetwork.
func (in *VSphereDistributedNetwork) DeepCopy() *VSphereDistributedNetwork {
	if in == nil {
		return nil
	}
	out := new(VSphereDistributedNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereDistributedNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereDistributedNetworkCondition) DeepCopyInto(out *VSphereDistributedNetworkCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereDistributedNetworkCondition.
func (in *VSphereDistributedNetworkCondition) DeepCopy() *VSphereDistributedNetworkCondition {
	if in == nil {
		return nil
	}
	out := new(VSphereDistributedNetworkCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereDistributedNetworkList) DeepCopyInto(out *VSphereDistributedNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VSphereDistributedNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereDistributedNetworkList.
func (in *VSphereDistributedNetworkList) DeepCopy() *VSphereDistributedNetworkList {
	if in == nil {
		return nil
	}
	out := new(VSphereDistributedNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereDistributedNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereDistributedNetworkSpec) DeepCopyInto(out *VSphereDistributedNetworkSpec) {
	*out = *in
	if in.IPPools != nil {
		in, out := &in.IPPools, &out.IPPools
		*out = make([]IPPoolReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereDistributedNetworkSpec.
func (in *VSphereDistributedNetworkSpec) DeepCopy() *VSphereDistributedNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(VSphereDistributedNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereDistributedNetworkStatus) DeepCopyInto(out *VSphereDistributedNetworkStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]VSphereDistributedNetworkCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereDistributedNetworkStatus.
func (in *VSphereDistributedNetworkStatus) DeepCopy() *VSphereDistributedNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(VSphereDistributedNetworkStatus)
	in.DeepCopyInto(out)
	return out
}
