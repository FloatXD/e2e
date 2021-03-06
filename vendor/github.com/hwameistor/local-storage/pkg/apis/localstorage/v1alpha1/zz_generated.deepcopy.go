// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ATASmartHealthAttribute) DeepCopyInto(out *ATASmartHealthAttribute) {
	*out = *in
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = new(ATASmartHealthAttributeFlag)
		**out = **in
	}
	if in.Raw != nil {
		in, out := &in.Raw, &out.Raw
		*out = new(ATASmartHealthAttributeRawData)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ATASmartHealthAttribute.
func (in *ATASmartHealthAttribute) DeepCopy() *ATASmartHealthAttribute {
	if in == nil {
		return nil
	}
	out := new(ATASmartHealthAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ATASmartHealthAttributeFlag) DeepCopyInto(out *ATASmartHealthAttributeFlag) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ATASmartHealthAttributeFlag.
func (in *ATASmartHealthAttributeFlag) DeepCopy() *ATASmartHealthAttributeFlag {
	if in == nil {
		return nil
	}
	out := new(ATASmartHealthAttributeFlag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ATASmartHealthAttributeRawData) DeepCopyInto(out *ATASmartHealthAttributeRawData) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ATASmartHealthAttributeRawData.
func (in *ATASmartHealthAttributeRawData) DeepCopy() *ATASmartHealthAttributeRawData {
	if in == nil {
		return nil
	}
	out := new(ATASmartHealthAttributeRawData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ATASmartHealthDetailsInfo) DeepCopyInto(out *ATASmartHealthDetailsInfo) {
	*out = *in
	if in.AttributesTable != nil {
		in, out := &in.AttributesTable, &out.AttributesTable
		*out = make([]ATASmartHealthAttribute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ATASmartHealthDetailsInfo.
func (in *ATASmartHealthDetailsInfo) DeepCopy() *ATASmartHealthDetailsInfo {
	if in == nil {
		return nil
	}
	out := new(ATASmartHealthDetailsInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessibilityTopology) DeepCopyInto(out *AccessibilityTopology) {
	*out = *in
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessibilityTopology.
func (in *AccessibilityTopology) DeepCopy() *AccessibilityTopology {
	if in == nil {
		return nil
	}
	out := new(AccessibilityTopology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRBDSystemConfig) DeepCopyInto(out *DRBDSystemConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRBDSystemConfig.
func (in *DRBDSystemConfig) DeepCopy() *DRBDSystemConfig {
	if in == nil {
		return nil
	}
	out := new(DRBDSystemConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HAState) DeepCopyInto(out *HAState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HAState.
func (in *HAState) DeepCopy() *HAState {
	if in == nil {
		return nil
	}
	out := new(HAState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDisk) DeepCopyInto(out *LocalDisk) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDisk.
func (in *LocalDisk) DeepCopy() *LocalDisk {
	if in == nil {
		return nil
	}
	out := new(LocalDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalPool) DeepCopyInto(out *LocalPool) {
	*out = *in
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]LocalDisk, len(*in))
		copy(*out, *in)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalPool.
func (in *LocalPool) DeepCopy() *LocalPool {
	if in == nil {
		return nil
	}
	out := new(LocalPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageAlert) DeepCopyInto(out *LocalStorageAlert) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageAlert.
func (in *LocalStorageAlert) DeepCopy() *LocalStorageAlert {
	if in == nil {
		return nil
	}
	out := new(LocalStorageAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalStorageAlert) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageAlertList) DeepCopyInto(out *LocalStorageAlertList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalStorageAlert, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageAlertList.
func (in *LocalStorageAlertList) DeepCopy() *LocalStorageAlertList {
	if in == nil {
		return nil
	}
	out := new(LocalStorageAlertList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalStorageAlertList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageAlertSpec) DeepCopyInto(out *LocalStorageAlertSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageAlertSpec.
func (in *LocalStorageAlertSpec) DeepCopy() *LocalStorageAlertSpec {
	if in == nil {
		return nil
	}
	out := new(LocalStorageAlertSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageAlertStatus) DeepCopyInto(out *LocalStorageAlertStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageAlertStatus.
func (in *LocalStorageAlertStatus) DeepCopy() *LocalStorageAlertStatus {
	if in == nil {
		return nil
	}
	out := new(LocalStorageAlertStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageConfig) DeepCopyInto(out *LocalStorageConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageConfig.
func (in *LocalStorageConfig) DeepCopy() *LocalStorageConfig {
	if in == nil {
		return nil
	}
	out := new(LocalStorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageNode) DeepCopyInto(out *LocalStorageNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageNode.
func (in *LocalStorageNode) DeepCopy() *LocalStorageNode {
	if in == nil {
		return nil
	}
	out := new(LocalStorageNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalStorageNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageNodeList) DeepCopyInto(out *LocalStorageNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalStorageNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageNodeList.
func (in *LocalStorageNodeList) DeepCopy() *LocalStorageNodeList {
	if in == nil {
		return nil
	}
	out := new(LocalStorageNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalStorageNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageNodeSpec) DeepCopyInto(out *LocalStorageNodeSpec) {
	*out = *in
	out.Topo = in.Topo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageNodeSpec.
func (in *LocalStorageNodeSpec) DeepCopy() *LocalStorageNodeSpec {
	if in == nil {
		return nil
	}
	out := new(LocalStorageNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageNodeStatus) DeepCopyInto(out *LocalStorageNodeStatus) {
	*out = *in
	if in.Pools != nil {
		in, out := &in.Pools, &out.Pools
		*out = make(map[string]LocalPool, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageNodeStatus.
func (in *LocalStorageNodeStatus) DeepCopy() *LocalStorageNodeStatus {
	if in == nil {
		return nil
	}
	out := new(LocalStorageNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolume) DeepCopyInto(out *LocalVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolume.
func (in *LocalVolume) DeepCopy() *LocalVolume {
	if in == nil {
		return nil
	}
	out := new(LocalVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeConvert) DeepCopyInto(out *LocalVolumeConvert) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeConvert.
func (in *LocalVolumeConvert) DeepCopy() *LocalVolumeConvert {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeConvert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeConvert) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeConvertList) DeepCopyInto(out *LocalVolumeConvertList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolumeConvert, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeConvertList.
func (in *LocalVolumeConvertList) DeepCopy() *LocalVolumeConvertList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeConvertList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeConvertList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeConvertSpec) DeepCopyInto(out *LocalVolumeConvertSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeConvertSpec.
func (in *LocalVolumeConvertSpec) DeepCopy() *LocalVolumeConvertSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeConvertSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeConvertStatus) DeepCopyInto(out *LocalVolumeConvertStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeConvertStatus.
func (in *LocalVolumeConvertStatus) DeepCopy() *LocalVolumeConvertStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeConvertStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeExpand) DeepCopyInto(out *LocalVolumeExpand) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeExpand.
func (in *LocalVolumeExpand) DeepCopy() *LocalVolumeExpand {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeExpand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeExpand) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeExpandList) DeepCopyInto(out *LocalVolumeExpandList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolumeExpand, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeExpandList.
func (in *LocalVolumeExpandList) DeepCopy() *LocalVolumeExpandList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeExpandList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeExpandList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeExpandSpec) DeepCopyInto(out *LocalVolumeExpandSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeExpandSpec.
func (in *LocalVolumeExpandSpec) DeepCopy() *LocalVolumeExpandSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeExpandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeExpandStatus) DeepCopyInto(out *LocalVolumeExpandStatus) {
	*out = *in
	if in.Subs != nil {
		in, out := &in.Subs, &out.Subs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeExpandStatus.
func (in *LocalVolumeExpandStatus) DeepCopy() *LocalVolumeExpandStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeExpandStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeList) DeepCopyInto(out *LocalVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeList.
func (in *LocalVolumeList) DeepCopy() *LocalVolumeList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeMigrate) DeepCopyInto(out *LocalVolumeMigrate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeMigrate.
func (in *LocalVolumeMigrate) DeepCopy() *LocalVolumeMigrate {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeMigrate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeMigrate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeMigrateList) DeepCopyInto(out *LocalVolumeMigrateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolumeMigrate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeMigrateList.
func (in *LocalVolumeMigrateList) DeepCopy() *LocalVolumeMigrateList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeMigrateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeMigrateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeMigrateSpec) DeepCopyInto(out *LocalVolumeMigrateSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeMigrateSpec.
func (in *LocalVolumeMigrateSpec) DeepCopy() *LocalVolumeMigrateSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeMigrateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeMigrateStatus) DeepCopyInto(out *LocalVolumeMigrateStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeMigrateStatus.
func (in *LocalVolumeMigrateStatus) DeepCopy() *LocalVolumeMigrateStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeMigrateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeReplica) DeepCopyInto(out *LocalVolumeReplica) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeReplica.
func (in *LocalVolumeReplica) DeepCopy() *LocalVolumeReplica {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeReplica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeReplica) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeReplicaList) DeepCopyInto(out *LocalVolumeReplicaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolumeReplica, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeReplicaList.
func (in *LocalVolumeReplicaList) DeepCopy() *LocalVolumeReplicaList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeReplicaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeReplicaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeReplicaSpec) DeepCopyInto(out *LocalVolumeReplicaSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeReplicaSpec.
func (in *LocalVolumeReplicaSpec) DeepCopy() *LocalVolumeReplicaSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeReplicaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeReplicaStatus) DeepCopyInto(out *LocalVolumeReplicaStatus) {
	*out = *in
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HAState != nil {
		in, out := &in.HAState, &out.HAState
		*out = new(HAState)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeReplicaStatus.
func (in *LocalVolumeReplicaStatus) DeepCopy() *LocalVolumeReplicaStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeReplicaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeSpec) DeepCopyInto(out *LocalVolumeSpec) {
	*out = *in
	in.Accessibility.DeepCopyInto(&out.Accessibility)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(VolumeConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeSpec.
func (in *LocalVolumeSpec) DeepCopy() *LocalVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeStatus) DeepCopyInto(out *LocalVolumeStatus) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeStatus.
func (in *LocalVolumeStatus) DeepCopy() *LocalVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NVMeSmartHealthDetailsInfo) DeepCopyInto(out *NVMeSmartHealthDetailsInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NVMeSmartHealthDetailsInfo.
func (in *NVMeSmartHealthDetailsInfo) DeepCopy() *NVMeSmartHealthDetailsInfo {
	if in == nil {
		return nil
	}
	out := new(NVMeSmartHealthDetailsInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfig) DeepCopyInto(out *NodeConfig) {
	*out = *in
	if in.Topology != nil {
		in, out := &in.Topology, &out.Topology
		*out = new(Topology)
		**out = **in
	}
	if in.LocalStorageConfig != nil {
		in, out := &in.LocalStorageConfig, &out.LocalStorageConfig
		*out = new(LocalStorageConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfig.
func (in *NodeConfig) DeepCopy() *NodeConfig {
	if in == nil {
		return nil
	}
	out := new(NodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhyDiskPowerOnTimeStatus) DeepCopyInto(out *PhyDiskPowerOnTimeStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhyDiskPowerOnTimeStatus.
func (in *PhyDiskPowerOnTimeStatus) DeepCopy() *PhyDiskPowerOnTimeStatus {
	if in == nil {
		return nil
	}
	out := new(PhyDiskPowerOnTimeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhyDiskSmartStatus) DeepCopyInto(out *PhyDiskSmartStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhyDiskSmartStatus.
func (in *PhyDiskSmartStatus) DeepCopy() *PhyDiskSmartStatus {
	if in == nil {
		return nil
	}
	out := new(PhyDiskSmartStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhyDiskTemperatureStatus) DeepCopyInto(out *PhyDiskTemperatureStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhyDiskTemperatureStatus.
func (in *PhyDiskTemperatureStatus) DeepCopy() *PhyDiskTemperatureStatus {
	if in == nil {
		return nil
	}
	out := new(PhyDiskTemperatureStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhysicalDisk) DeepCopyInto(out *PhysicalDisk) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhysicalDisk.
func (in *PhysicalDisk) DeepCopy() *PhysicalDisk {
	if in == nil {
		return nil
	}
	out := new(PhysicalDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PhysicalDisk) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhysicalDiskList) DeepCopyInto(out *PhysicalDiskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PhysicalDisk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhysicalDiskList.
func (in *PhysicalDiskList) DeepCopy() *PhysicalDiskList {
	if in == nil {
		return nil
	}
	out := new(PhysicalDiskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PhysicalDiskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhysicalDiskSpec) DeepCopyInto(out *PhysicalDiskSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhysicalDiskSpec.
func (in *PhysicalDiskSpec) DeepCopy() *PhysicalDiskSpec {
	if in == nil {
		return nil
	}
	out := new(PhysicalDiskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhysicalDiskStatus) DeepCopyInto(out *PhysicalDiskStatus) {
	*out = *in
	if in.SmartCheck != nil {
		in, out := &in.SmartCheck, &out.SmartCheck
		*out = new(SmartCheck)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhysicalDiskStatus.
func (in *PhysicalDiskStatus) DeepCopy() *PhysicalDiskStatus {
	if in == nil {
		return nil
	}
	out := new(PhysicalDiskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SCSIErrorCounter) DeepCopyInto(out *SCSIErrorCounter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCSIErrorCounter.
func (in *SCSIErrorCounter) DeepCopy() *SCSIErrorCounter {
	if in == nil {
		return nil
	}
	out := new(SCSIErrorCounter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SCSISmartHealthDetailsInfo) DeepCopyInto(out *SCSISmartHealthDetailsInfo) {
	*out = *in
	if in.Read != nil {
		in, out := &in.Read, &out.Read
		*out = new(SCSIErrorCounter)
		**out = **in
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = new(SCSIErrorCounter)
		**out = **in
	}
	if in.Verify != nil {
		in, out := &in.Verify, &out.Verify
		*out = new(SCSIErrorCounter)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCSISmartHealthDetailsInfo.
func (in *SCSISmartHealthDetailsInfo) DeepCopy() *SCSISmartHealthDetailsInfo {
	if in == nil {
		return nil
	}
	out := new(SCSISmartHealthDetailsInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartCheck) DeepCopyInto(out *SmartCheck) {
	*out = *in
	if in.SmartStatus != nil {
		in, out := &in.SmartStatus, &out.SmartStatus
		*out = new(PhyDiskSmartStatus)
		**out = **in
	}
	if in.Temperature != nil {
		in, out := &in.Temperature, &out.Temperature
		*out = new(PhyDiskTemperatureStatus)
		**out = **in
	}
	if in.PowerOnTime != nil {
		in, out := &in.PowerOnTime, &out.PowerOnTime
		*out = new(PhyDiskPowerOnTimeStatus)
		**out = **in
	}
	if in.NVMeSmartHealthStatus != nil {
		in, out := &in.NVMeSmartHealthStatus, &out.NVMeSmartHealthStatus
		*out = new(NVMeSmartHealthDetailsInfo)
		**out = **in
	}
	if in.ATASmartHealthStatus != nil {
		in, out := &in.ATASmartHealthStatus, &out.ATASmartHealthStatus
		*out = new(ATASmartHealthDetailsInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.SCSISmartHealthStatus != nil {
		in, out := &in.SCSISmartHealthStatus, &out.SCSISmartHealthStatus
		*out = new(SCSISmartHealthDetailsInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.LastTime != nil {
		in, out := &in.LastTime, &out.LastTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartCheck.
func (in *SmartCheck) DeepCopy() *SmartCheck {
	if in == nil {
		return nil
	}
	out := new(SmartCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemConfig) DeepCopyInto(out *SystemConfig) {
	*out = *in
	if in.DRBD != nil {
		in, out := &in.DRBD, &out.DRBD
		*out = new(DRBDSystemConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemConfig.
func (in *SystemConfig) DeepCopy() *SystemConfig {
	if in == nil {
		return nil
	}
	out := new(SystemConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Topology) DeepCopyInto(out *Topology) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Topology.
func (in *Topology) DeepCopy() *Topology {
	if in == nil {
		return nil
	}
	out := new(Topology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeConfig) DeepCopyInto(out *VolumeConfig) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]VolumeReplica, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeConfig.
func (in *VolumeConfig) DeepCopy() *VolumeConfig {
	if in == nil {
		return nil
	}
	out := new(VolumeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeReplica) DeepCopyInto(out *VolumeReplica) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeReplica.
func (in *VolumeReplica) DeepCopy() *VolumeReplica {
	if in == nil {
		return nil
	}
	out := new(VolumeReplica)
	in.DeepCopyInto(out)
	return out
}
