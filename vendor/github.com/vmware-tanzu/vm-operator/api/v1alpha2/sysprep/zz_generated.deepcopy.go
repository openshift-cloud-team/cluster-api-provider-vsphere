//go:build !ignore_autogenerated

// Copyright (c) 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package sysprep

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainPasswordSecretKeySelector) DeepCopyInto(out *DomainPasswordSecretKeySelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainPasswordSecretKeySelector.
func (in *DomainPasswordSecretKeySelector) DeepCopy() *DomainPasswordSecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(DomainPasswordSecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GUIRunOnce) DeepCopyInto(out *GUIRunOnce) {
	*out = *in
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GUIRunOnce.
func (in *GUIRunOnce) DeepCopy() *GUIRunOnce {
	if in == nil {
		return nil
	}
	out := new(GUIRunOnce)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GUIUnattended) DeepCopyInto(out *GUIUnattended) {
	*out = *in
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(PasswordSecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GUIUnattended.
func (in *GUIUnattended) DeepCopy() *GUIUnattended {
	if in == nil {
		return nil
	}
	out := new(GUIUnattended)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identification) DeepCopyInto(out *Identification) {
	*out = *in
	if in.DomainAdminPassword != nil {
		in, out := &in.DomainAdminPassword, &out.DomainAdminPassword
		*out = new(DomainPasswordSecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identification.
func (in *Identification) DeepCopy() *Identification {
	if in == nil {
		return nil
	}
	out := new(Identification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseFilePrintData) DeepCopyInto(out *LicenseFilePrintData) {
	*out = *in
	if in.AutoUsers != nil {
		in, out := &in.AutoUsers, &out.AutoUsers
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseFilePrintData.
func (in *LicenseFilePrintData) DeepCopy() *LicenseFilePrintData {
	if in == nil {
		return nil
	}
	out := new(LicenseFilePrintData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordSecretKeySelector) DeepCopyInto(out *PasswordSecretKeySelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordSecretKeySelector.
func (in *PasswordSecretKeySelector) DeepCopy() *PasswordSecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(PasswordSecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductIDSecretKeySelector) DeepCopyInto(out *ProductIDSecretKeySelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductIDSecretKeySelector.
func (in *ProductIDSecretKeySelector) DeepCopy() *ProductIDSecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(ProductIDSecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sysprep) DeepCopyInto(out *Sysprep) {
	*out = *in
	in.GUIRunOnce.DeepCopyInto(&out.GUIRunOnce)
	if in.GUIUnattended != nil {
		in, out := &in.GUIUnattended, &out.GUIUnattended
		*out = new(GUIUnattended)
		(*in).DeepCopyInto(*out)
	}
	if in.Identification != nil {
		in, out := &in.Identification, &out.Identification
		*out = new(Identification)
		(*in).DeepCopyInto(*out)
	}
	if in.LicenseFilePrintData != nil {
		in, out := &in.LicenseFilePrintData, &out.LicenseFilePrintData
		*out = new(LicenseFilePrintData)
		(*in).DeepCopyInto(*out)
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(UserData)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sysprep.
func (in *Sysprep) DeepCopy() *Sysprep {
	if in == nil {
		return nil
	}
	out := new(Sysprep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserData) DeepCopyInto(out *UserData) {
	*out = *in
	if in.ProductID != nil {
		in, out := &in.ProductID, &out.ProductID
		*out = new(ProductIDSecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserData.
func (in *UserData) DeepCopy() *UserData {
	if in == nil {
		return nil
	}
	out := new(UserData)
	in.DeepCopyInto(out)
	return out
}
