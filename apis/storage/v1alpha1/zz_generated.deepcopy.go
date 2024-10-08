//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRuleInitParameters) DeepCopyInto(out *BackupRuleInitParameters) {
	*out = *in
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.Retention != nil {
		in, out := &in.Retention, &out.Retention
		*out = new(float64)
		**out = **in
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRuleInitParameters.
func (in *BackupRuleInitParameters) DeepCopy() *BackupRuleInitParameters {
	if in == nil {
		return nil
	}
	out := new(BackupRuleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRuleObservation) DeepCopyInto(out *BackupRuleObservation) {
	*out = *in
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.Retention != nil {
		in, out := &in.Retention, &out.Retention
		*out = new(float64)
		**out = **in
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRuleObservation.
func (in *BackupRuleObservation) DeepCopy() *BackupRuleObservation {
	if in == nil {
		return nil
	}
	out := new(BackupRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRuleParameters) DeepCopyInto(out *BackupRuleParameters) {
	*out = *in
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.Retention != nil {
		in, out := &in.Retention, &out.Retention
		*out = new(float64)
		**out = **in
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRuleParameters.
func (in *BackupRuleParameters) DeepCopy() *BackupRuleParameters {
	if in == nil {
		return nil
	}
	out := new(BackupRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneInitParameters) DeepCopyInto(out *CloneInitParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneInitParameters.
func (in *CloneInitParameters) DeepCopy() *CloneInitParameters {
	if in == nil {
		return nil
	}
	out := new(CloneInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneObservation) DeepCopyInto(out *CloneObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneObservation.
func (in *CloneObservation) DeepCopy() *CloneObservation {
	if in == nil {
		return nil
	}
	out := new(CloneObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloneParameters) DeepCopyInto(out *CloneParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloneParameters.
func (in *CloneParameters) DeepCopy() *CloneParameters {
	if in == nil {
		return nil
	}
	out := new(CloneParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportInitParameters) DeepCopyInto(out *ImportInitParameters) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.SourceHash != nil {
		in, out := &in.SourceHash, &out.SourceHash
		*out = new(string)
		**out = **in
	}
	if in.SourceLocation != nil {
		in, out := &in.SourceLocation, &out.SourceLocation
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportInitParameters.
func (in *ImportInitParameters) DeepCopy() *ImportInitParameters {
	if in == nil {
		return nil
	}
	out := new(ImportInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportObservation) DeepCopyInto(out *ImportObservation) {
	*out = *in
	if in.Sha256Sum != nil {
		in, out := &in.Sha256Sum, &out.Sha256Sum
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.SourceHash != nil {
		in, out := &in.SourceHash, &out.SourceHash
		*out = new(string)
		**out = **in
	}
	if in.SourceLocation != nil {
		in, out := &in.SourceLocation, &out.SourceLocation
		*out = new(string)
		**out = **in
	}
	if in.WrittenBytes != nil {
		in, out := &in.WrittenBytes, &out.WrittenBytes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportObservation.
func (in *ImportObservation) DeepCopy() *ImportObservation {
	if in == nil {
		return nil
	}
	out := new(ImportObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportParameters) DeepCopyInto(out *ImportParameters) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.SourceHash != nil {
		in, out := &in.SourceHash, &out.SourceHash
		*out = new(string)
		**out = **in
	}
	if in.SourceLocation != nil {
		in, out := &in.SourceLocation, &out.SourceLocation
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportParameters.
func (in *ImportParameters) DeepCopy() *ImportParameters {
	if in == nil {
		return nil
	}
	out := new(ImportParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Storage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageInitParameters) DeepCopyInto(out *StorageInitParameters) {
	*out = *in
	if in.BackupRule != nil {
		in, out := &in.BackupRule, &out.BackupRule
		*out = make([]BackupRuleInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Clone != nil {
		in, out := &in.Clone, &out.Clone
		*out = make([]CloneInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeleteAutoresizeBackup != nil {
		in, out := &in.DeleteAutoresizeBackup, &out.DeleteAutoresizeBackup
		*out = new(bool)
		**out = **in
	}
	if in.Encrypt != nil {
		in, out := &in.Encrypt, &out.Encrypt
		*out = new(bool)
		**out = **in
	}
	if in.FilesystemAutoresize != nil {
		in, out := &in.FilesystemAutoresize, &out.FilesystemAutoresize
		*out = new(bool)
		**out = **in
	}
	if in.Import != nil {
		in, out := &in.Import, &out.Import
		*out = make([]ImportInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageInitParameters.
func (in *StorageInitParameters) DeepCopy() *StorageInitParameters {
	if in == nil {
		return nil
	}
	out := new(StorageInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageList) DeepCopyInto(out *StorageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Storage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageList.
func (in *StorageList) DeepCopy() *StorageList {
	if in == nil {
		return nil
	}
	out := new(StorageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageObservation) DeepCopyInto(out *StorageObservation) {
	*out = *in
	if in.BackupRule != nil {
		in, out := &in.BackupRule, &out.BackupRule
		*out = make([]BackupRuleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Clone != nil {
		in, out := &in.Clone, &out.Clone
		*out = make([]CloneObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeleteAutoresizeBackup != nil {
		in, out := &in.DeleteAutoresizeBackup, &out.DeleteAutoresizeBackup
		*out = new(bool)
		**out = **in
	}
	if in.Encrypt != nil {
		in, out := &in.Encrypt, &out.Encrypt
		*out = new(bool)
		**out = **in
	}
	if in.FilesystemAutoresize != nil {
		in, out := &in.FilesystemAutoresize, &out.FilesystemAutoresize
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Import != nil {
		in, out := &in.Import, &out.Import
		*out = make([]ImportObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.SystemLabels != nil {
		in, out := &in.SystemLabels, &out.SystemLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageObservation.
func (in *StorageObservation) DeepCopy() *StorageObservation {
	if in == nil {
		return nil
	}
	out := new(StorageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageParameters) DeepCopyInto(out *StorageParameters) {
	*out = *in
	if in.BackupRule != nil {
		in, out := &in.BackupRule, &out.BackupRule
		*out = make([]BackupRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Clone != nil {
		in, out := &in.Clone, &out.Clone
		*out = make([]CloneParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeleteAutoresizeBackup != nil {
		in, out := &in.DeleteAutoresizeBackup, &out.DeleteAutoresizeBackup
		*out = new(bool)
		**out = **in
	}
	if in.Encrypt != nil {
		in, out := &in.Encrypt, &out.Encrypt
		*out = new(bool)
		**out = **in
	}
	if in.FilesystemAutoresize != nil {
		in, out := &in.FilesystemAutoresize, &out.FilesystemAutoresize
		*out = new(bool)
		**out = **in
	}
	if in.Import != nil {
		in, out := &in.Import, &out.Import
		*out = make([]ImportParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageParameters.
func (in *StorageParameters) DeepCopy() *StorageParameters {
	if in == nil {
		return nil
	}
	out := new(StorageParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageStatus) DeepCopyInto(out *StorageStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageStatus.
func (in *StorageStatus) DeepCopy() *StorageStatus {
	if in == nil {
		return nil
	}
	out := new(StorageStatus)
	in.DeepCopyInto(out)
	return out
}
