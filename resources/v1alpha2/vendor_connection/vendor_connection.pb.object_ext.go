// Code generated by protoc-gen-goten-object
// File: workplace/proto/v1alpha2/vendor_connection.proto
// DO NOT EDIT!!!

package vendor_connection

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha2/common"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = fieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &iam_project.Project{}
	_ = &workplace_common.BBox{}
)

func (o *PointGrab) GotenObjectExt() {}

func (o *PointGrab) MakeFullFieldMask() *PointGrab_FieldMask {
	return FullPointGrab_FieldMask()
}

func (o *PointGrab) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPointGrab_FieldMask()
}

func (o *PointGrab) MakeDiffFieldMask(other *PointGrab) *PointGrab_FieldMask {
	if o == nil && other == nil {
		return &PointGrab_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPointGrab_FieldMask()
	}

	res := &PointGrab_FieldMask{}
	return res
}

func (o *PointGrab) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*PointGrab))
}

func (o *PointGrab) Clone() *PointGrab {
	if o == nil {
		return nil
	}
	result := &PointGrab{}
	return result
}

func (o *PointGrab) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *PointGrab) Merge(source *PointGrab) {
}

func (o *PointGrab) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*PointGrab))
}

func (o *PointGrab_Telemetry) GotenObjectExt() {}

func (o *PointGrab_Telemetry) MakeFullFieldMask() *PointGrab_Telemetry_FieldMask {
	return FullPointGrab_Telemetry_FieldMask()
}

func (o *PointGrab_Telemetry) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPointGrab_Telemetry_FieldMask()
}

func (o *PointGrab_Telemetry) MakeDiffFieldMask(other *PointGrab_Telemetry) *PointGrab_Telemetry_FieldMask {
	if o == nil && other == nil {
		return &PointGrab_Telemetry_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPointGrab_Telemetry_FieldMask()
	}

	res := &PointGrab_Telemetry_FieldMask{}
	return res
}

func (o *PointGrab_Telemetry) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*PointGrab_Telemetry))
}

func (o *PointGrab_Telemetry) Clone() *PointGrab_Telemetry {
	if o == nil {
		return nil
	}
	result := &PointGrab_Telemetry{}
	return result
}

func (o *PointGrab_Telemetry) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *PointGrab_Telemetry) Merge(source *PointGrab_Telemetry) {
}

func (o *PointGrab_Telemetry) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*PointGrab_Telemetry))
}

func (o *VendorConnection) GotenObjectExt() {}

func (o *VendorConnection) MakeFullFieldMask() *VendorConnection_FieldMask {
	return FullVendorConnection_FieldMask()
}

func (o *VendorConnection) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullVendorConnection_FieldMask()
}

func (o *VendorConnection) MakeDiffFieldMask(other *VendorConnection) *VendorConnection_FieldMask {
	if o == nil && other == nil {
		return &VendorConnection_FieldMask{}
	}
	if o == nil || other == nil {
		return FullVendorConnection_FieldMask()
	}

	res := &VendorConnection_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &VendorConnection_FieldTerminalPath{selector: VendorConnection_FieldPathSelectorName})
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &VendorConnection_FieldTerminalPath{selector: VendorConnection_FieldPathSelectorDisplayName})
	}
	if o.GetVendor() != other.GetVendor() {
		res.Paths = append(res.Paths, &VendorConnection_FieldTerminalPath{selector: VendorConnection_FieldPathSelectorVendor})
	}
	{
		subMask := o.GetSpec().MakeDiffFieldMask(other.GetSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &VendorConnection_FieldTerminalPath{selector: VendorConnection_FieldPathSelectorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &VendorConnection_FieldSubPath{selector: VendorConnection_FieldPathSelectorSpec, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetInfo().MakeDiffFieldMask(other.GetInfo())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &VendorConnection_FieldTerminalPath{selector: VendorConnection_FieldPathSelectorInfo})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &VendorConnection_FieldSubPath{selector: VendorConnection_FieldPathSelectorInfo, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &VendorConnection_FieldTerminalPath{selector: VendorConnection_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &VendorConnection_FieldSubPath{selector: VendorConnection_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	return res
}

func (o *VendorConnection) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*VendorConnection))
}

func (o *VendorConnection) Clone() *VendorConnection {
	if o == nil {
		return nil
	}
	result := &VendorConnection{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &Name{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.DisplayName = o.DisplayName
	result.Vendor = o.Vendor
	result.Spec = o.Spec.Clone()
	result.Info = o.Info.Clone()
	result.Metadata = o.Metadata.Clone()
	return result
}

func (o *VendorConnection) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *VendorConnection) Merge(source *VendorConnection) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &Name{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
	o.DisplayName = source.GetDisplayName()
	o.Vendor = source.GetVendor()
	if source.GetSpec() != nil {
		if o.Spec == nil {
			o.Spec = new(VendorConnection_Spec)
		}
		o.Spec.Merge(source.GetSpec())
	}
	if source.GetInfo() != nil {
		if o.Info == nil {
			o.Info = new(VendorConnection_Info)
		}
		o.Info.Merge(source.GetInfo())
	}
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(ntt_meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
}

func (o *VendorConnection) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*VendorConnection))
}

func (o *VendorConnection_Spec) GotenObjectExt() {}

func (o *VendorConnection_Spec) MakeFullFieldMask() *VendorConnection_Spec_FieldMask {
	return FullVendorConnection_Spec_FieldMask()
}

func (o *VendorConnection_Spec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullVendorConnection_Spec_FieldMask()
}

func (o *VendorConnection_Spec) MakeDiffFieldMask(other *VendorConnection_Spec) *VendorConnection_Spec_FieldMask {
	if o == nil && other == nil {
		return &VendorConnection_Spec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullVendorConnection_Spec_FieldMask()
	}

	res := &VendorConnection_Spec_FieldMask{}
	{
		subMask := o.GetPointGrab().MakeDiffFieldMask(other.GetPointGrab())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &VendorConnectionSpec_FieldTerminalPath{selector: VendorConnectionSpec_FieldPathSelectorPointGrab})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &VendorConnectionSpec_FieldSubPath{selector: VendorConnectionSpec_FieldPathSelectorPointGrab, subPath: subpath})
			}
		}
	}
	return res
}

func (o *VendorConnection_Spec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*VendorConnection_Spec))
}

func (o *VendorConnection_Spec) Clone() *VendorConnection_Spec {
	if o == nil {
		return nil
	}
	result := &VendorConnection_Spec{}
	result.PointGrab = o.PointGrab.Clone()
	return result
}

func (o *VendorConnection_Spec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *VendorConnection_Spec) Merge(source *VendorConnection_Spec) {
	if source.GetPointGrab() != nil {
		if o.PointGrab == nil {
			o.PointGrab = new(VendorConnection_Spec_PointGrabSpec)
		}
		o.PointGrab.Merge(source.GetPointGrab())
	}
}

func (o *VendorConnection_Spec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*VendorConnection_Spec))
}

func (o *VendorConnection_Info) GotenObjectExt() {}

func (o *VendorConnection_Info) MakeFullFieldMask() *VendorConnection_Info_FieldMask {
	return FullVendorConnection_Info_FieldMask()
}

func (o *VendorConnection_Info) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullVendorConnection_Info_FieldMask()
}

func (o *VendorConnection_Info) MakeDiffFieldMask(other *VendorConnection_Info) *VendorConnection_Info_FieldMask {
	if o == nil && other == nil {
		return &VendorConnection_Info_FieldMask{}
	}
	if o == nil || other == nil {
		return FullVendorConnection_Info_FieldMask()
	}

	res := &VendorConnection_Info_FieldMask{}
	{
		subMask := o.GetInfo().MakeDiffFieldMask(other.GetInfo())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &VendorConnectionInfo_FieldTerminalPath{selector: VendorConnectionInfo_FieldPathSelectorInfo})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &VendorConnectionInfo_FieldSubPath{selector: VendorConnectionInfo_FieldPathSelectorInfo, subPath: subpath})
			}
		}
	}
	return res
}

func (o *VendorConnection_Info) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*VendorConnection_Info))
}

func (o *VendorConnection_Info) Clone() *VendorConnection_Info {
	if o == nil {
		return nil
	}
	result := &VendorConnection_Info{}
	result.Info = o.Info.Clone()
	return result
}

func (o *VendorConnection_Info) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *VendorConnection_Info) Merge(source *VendorConnection_Info) {
	if source.GetInfo() != nil {
		if o.Info == nil {
			o.Info = new(VendorConnection_Info_PointGrabInfo)
		}
		o.Info.Merge(source.GetInfo())
	}
}

func (o *VendorConnection_Info) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*VendorConnection_Info))
}

func (o *VendorConnection_Spec_PointGrabSpec) GotenObjectExt() {}

func (o *VendorConnection_Spec_PointGrabSpec) MakeFullFieldMask() *VendorConnection_Spec_PointGrabSpec_FieldMask {
	return FullVendorConnection_Spec_PointGrabSpec_FieldMask()
}

func (o *VendorConnection_Spec_PointGrabSpec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullVendorConnection_Spec_PointGrabSpec_FieldMask()
}

func (o *VendorConnection_Spec_PointGrabSpec) MakeDiffFieldMask(other *VendorConnection_Spec_PointGrabSpec) *VendorConnection_Spec_PointGrabSpec_FieldMask {
	if o == nil && other == nil {
		return &VendorConnection_Spec_PointGrabSpec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullVendorConnection_Spec_PointGrabSpec_FieldMask()
	}

	res := &VendorConnection_Spec_PointGrabSpec_FieldMask{}
	if o.GetEndpoint() != other.GetEndpoint() {
		res.Paths = append(res.Paths, &VendorConnectionSpecPointGrabSpec_FieldTerminalPath{selector: VendorConnectionSpecPointGrabSpec_FieldPathSelectorEndpoint})
	}
	{
		subMask := o.GetAuth().MakeDiffFieldMask(other.GetAuth())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &VendorConnectionSpecPointGrabSpec_FieldTerminalPath{selector: VendorConnectionSpecPointGrabSpec_FieldPathSelectorAuth})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &VendorConnectionSpecPointGrabSpec_FieldSubPath{selector: VendorConnectionSpecPointGrabSpec_FieldPathSelectorAuth, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetCallbackSpec().MakeDiffFieldMask(other.GetCallbackSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &VendorConnectionSpecPointGrabSpec_FieldTerminalPath{selector: VendorConnectionSpecPointGrabSpec_FieldPathSelectorCallbackSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &VendorConnectionSpecPointGrabSpec_FieldSubPath{selector: VendorConnectionSpecPointGrabSpec_FieldPathSelectorCallbackSpec, subPath: subpath})
			}
		}
	}
	return res
}

func (o *VendorConnection_Spec_PointGrabSpec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*VendorConnection_Spec_PointGrabSpec))
}

func (o *VendorConnection_Spec_PointGrabSpec) Clone() *VendorConnection_Spec_PointGrabSpec {
	if o == nil {
		return nil
	}
	result := &VendorConnection_Spec_PointGrabSpec{}
	result.Endpoint = o.Endpoint
	result.Auth = o.Auth.Clone()
	result.CallbackSpec = o.CallbackSpec.Clone()
	return result
}

func (o *VendorConnection_Spec_PointGrabSpec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *VendorConnection_Spec_PointGrabSpec) Merge(source *VendorConnection_Spec_PointGrabSpec) {
	o.Endpoint = source.GetEndpoint()
	if source.GetAuth() != nil {
		if o.Auth == nil {
			o.Auth = new(VendorConnection_Spec_PointGrabSpec_Auth)
		}
		o.Auth.Merge(source.GetAuth())
	}
	if source.GetCallbackSpec() != nil {
		if o.CallbackSpec == nil {
			o.CallbackSpec = new(VendorConnection_Spec_PointGrabSpec_CallbackSpec)
		}
		o.CallbackSpec.Merge(source.GetCallbackSpec())
	}
}

func (o *VendorConnection_Spec_PointGrabSpec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*VendorConnection_Spec_PointGrabSpec))
}

func (o *VendorConnection_Spec_PointGrabSpec_Auth) GotenObjectExt() {}

func (o *VendorConnection_Spec_PointGrabSpec_Auth) MakeFullFieldMask() *VendorConnection_Spec_PointGrabSpec_Auth_FieldMask {
	return FullVendorConnection_Spec_PointGrabSpec_Auth_FieldMask()
}

func (o *VendorConnection_Spec_PointGrabSpec_Auth) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullVendorConnection_Spec_PointGrabSpec_Auth_FieldMask()
}

func (o *VendorConnection_Spec_PointGrabSpec_Auth) MakeDiffFieldMask(other *VendorConnection_Spec_PointGrabSpec_Auth) *VendorConnection_Spec_PointGrabSpec_Auth_FieldMask {
	if o == nil && other == nil {
		return &VendorConnection_Spec_PointGrabSpec_Auth_FieldMask{}
	}
	if o == nil || other == nil {
		return FullVendorConnection_Spec_PointGrabSpec_Auth_FieldMask()
	}

	res := &VendorConnection_Spec_PointGrabSpec_Auth_FieldMask{}
	if o.GetApplication() != other.GetApplication() {
		res.Paths = append(res.Paths, &VendorConnectionSpecPointGrabSpecAuth_FieldTerminalPath{selector: VendorConnectionSpecPointGrabSpecAuth_FieldPathSelectorApplication})
	}
	if o.GetSecret() != other.GetSecret() {
		res.Paths = append(res.Paths, &VendorConnectionSpecPointGrabSpecAuth_FieldTerminalPath{selector: VendorConnectionSpecPointGrabSpecAuth_FieldPathSelectorSecret})
	}
	if o.GetTokenEndpoint() != other.GetTokenEndpoint() {
		res.Paths = append(res.Paths, &VendorConnectionSpecPointGrabSpecAuth_FieldTerminalPath{selector: VendorConnectionSpecPointGrabSpecAuth_FieldPathSelectorTokenEndpoint})
	}
	return res
}

func (o *VendorConnection_Spec_PointGrabSpec_Auth) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*VendorConnection_Spec_PointGrabSpec_Auth))
}

func (o *VendorConnection_Spec_PointGrabSpec_Auth) Clone() *VendorConnection_Spec_PointGrabSpec_Auth {
	if o == nil {
		return nil
	}
	result := &VendorConnection_Spec_PointGrabSpec_Auth{}
	result.Application = o.Application
	result.Secret = o.Secret
	result.TokenEndpoint = o.TokenEndpoint
	return result
}

func (o *VendorConnection_Spec_PointGrabSpec_Auth) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *VendorConnection_Spec_PointGrabSpec_Auth) Merge(source *VendorConnection_Spec_PointGrabSpec_Auth) {
	o.Application = source.GetApplication()
	o.Secret = source.GetSecret()
	o.TokenEndpoint = source.GetTokenEndpoint()
}

func (o *VendorConnection_Spec_PointGrabSpec_Auth) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*VendorConnection_Spec_PointGrabSpec_Auth))
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec) GotenObjectExt() {}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec) MakeFullFieldMask() *VendorConnection_Spec_PointGrabSpec_CallbackSpec_FieldMask {
	return FullVendorConnection_Spec_PointGrabSpec_CallbackSpec_FieldMask()
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullVendorConnection_Spec_PointGrabSpec_CallbackSpec_FieldMask()
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec) MakeDiffFieldMask(other *VendorConnection_Spec_PointGrabSpec_CallbackSpec) *VendorConnection_Spec_PointGrabSpec_CallbackSpec_FieldMask {
	if o == nil && other == nil {
		return &VendorConnection_Spec_PointGrabSpec_CallbackSpec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullVendorConnection_Spec_PointGrabSpec_CallbackSpec_FieldMask()
	}

	res := &VendorConnection_Spec_PointGrabSpec_CallbackSpec_FieldMask{}
	{
		subMask := o.GetAuthToken().MakeDiffFieldMask(other.GetAuthToken())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &VendorConnectionSpecPointGrabSpecCallbackSpec_FieldTerminalPath{selector: VendorConnectionSpecPointGrabSpecCallbackSpec_FieldPathSelectorAuthToken})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &VendorConnectionSpecPointGrabSpecCallbackSpec_FieldSubPath{selector: VendorConnectionSpecPointGrabSpecCallbackSpec_FieldPathSelectorAuthToken, subPath: subpath})
			}
		}
	}
	return res
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*VendorConnection_Spec_PointGrabSpec_CallbackSpec))
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec) Clone() *VendorConnection_Spec_PointGrabSpec_CallbackSpec {
	if o == nil {
		return nil
	}
	result := &VendorConnection_Spec_PointGrabSpec_CallbackSpec{}
	result.AuthToken = o.AuthToken.Clone()
	return result
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec) Merge(source *VendorConnection_Spec_PointGrabSpec_CallbackSpec) {
	if source.GetAuthToken() != nil {
		if o.AuthToken == nil {
			o.AuthToken = new(VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken)
		}
		o.AuthToken.Merge(source.GetAuthToken())
	}
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*VendorConnection_Spec_PointGrabSpec_CallbackSpec))
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken) GotenObjectExt() {}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken) MakeFullFieldMask() *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken_FieldMask {
	return FullVendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken_FieldMask()
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullVendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken_FieldMask()
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken) MakeDiffFieldMask(other *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken) *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken_FieldMask {
	if o == nil && other == nil {
		return &VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken_FieldMask{}
	}
	if o == nil || other == nil {
		return FullVendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken_FieldMask()
	}

	res := &VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken_FieldMask{}
	if o.GetTokenHeader() != other.GetTokenHeader() {
		res.Paths = append(res.Paths, &VendorConnectionSpecPointGrabSpecCallbackSpecAuthToken_FieldTerminalPath{selector: VendorConnectionSpecPointGrabSpecCallbackSpecAuthToken_FieldPathSelectorTokenHeader})
	}
	if o.GetTokenValue() != other.GetTokenValue() {
		res.Paths = append(res.Paths, &VendorConnectionSpecPointGrabSpecCallbackSpecAuthToken_FieldTerminalPath{selector: VendorConnectionSpecPointGrabSpecCallbackSpecAuthToken_FieldPathSelectorTokenValue})
	}
	return res
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken))
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken) Clone() *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken {
	if o == nil {
		return nil
	}
	result := &VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken{}
	result.TokenHeader = o.TokenHeader
	result.TokenValue = o.TokenValue
	return result
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken) Merge(source *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken) {
	o.TokenHeader = source.GetTokenHeader()
	o.TokenValue = source.GetTokenValue()
}

func (o *VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*VendorConnection_Spec_PointGrabSpec_CallbackSpec_AuthToken))
}

func (o *VendorConnection_Info_PointGrabInfo) GotenObjectExt() {}

func (o *VendorConnection_Info_PointGrabInfo) MakeFullFieldMask() *VendorConnection_Info_PointGrabInfo_FieldMask {
	return FullVendorConnection_Info_PointGrabInfo_FieldMask()
}

func (o *VendorConnection_Info_PointGrabInfo) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullVendorConnection_Info_PointGrabInfo_FieldMask()
}

func (o *VendorConnection_Info_PointGrabInfo) MakeDiffFieldMask(other *VendorConnection_Info_PointGrabInfo) *VendorConnection_Info_PointGrabInfo_FieldMask {
	if o == nil && other == nil {
		return &VendorConnection_Info_PointGrabInfo_FieldMask{}
	}
	if o == nil || other == nil {
		return FullVendorConnection_Info_PointGrabInfo_FieldMask()
	}

	res := &VendorConnection_Info_PointGrabInfo_FieldMask{}

	if len(o.GetTelemetrySubscriptions()) == len(other.GetTelemetrySubscriptions()) {
		for i, lValue := range o.GetTelemetrySubscriptions() {
			rValue := other.GetTelemetrySubscriptions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &VendorConnectionInfoPointGrabInfo_FieldTerminalPath{selector: VendorConnectionInfoPointGrabInfo_FieldPathSelectorTelemetrySubscriptions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &VendorConnectionInfoPointGrabInfo_FieldTerminalPath{selector: VendorConnectionInfoPointGrabInfo_FieldPathSelectorTelemetrySubscriptions})
	}
	return res
}

func (o *VendorConnection_Info_PointGrabInfo) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*VendorConnection_Info_PointGrabInfo))
}

func (o *VendorConnection_Info_PointGrabInfo) Clone() *VendorConnection_Info_PointGrabInfo {
	if o == nil {
		return nil
	}
	result := &VendorConnection_Info_PointGrabInfo{}
	result.TelemetrySubscriptions = make([]*VendorConnection_Info_PointGrabInfo_TelemetrySubscription, len(o.TelemetrySubscriptions))
	for i, sourceValue := range o.TelemetrySubscriptions {
		result.TelemetrySubscriptions[i] = sourceValue.Clone()
	}
	return result
}

func (o *VendorConnection_Info_PointGrabInfo) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *VendorConnection_Info_PointGrabInfo) Merge(source *VendorConnection_Info_PointGrabInfo) {
	for _, sourceValue := range source.GetTelemetrySubscriptions() {
		exists := false
		for _, currentValue := range o.TelemetrySubscriptions {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *VendorConnection_Info_PointGrabInfo_TelemetrySubscription
			if sourceValue != nil {
				newDstElement = new(VendorConnection_Info_PointGrabInfo_TelemetrySubscription)
				newDstElement.Merge(sourceValue)
			}
			o.TelemetrySubscriptions = append(o.TelemetrySubscriptions, newDstElement)
		}
	}

}

func (o *VendorConnection_Info_PointGrabInfo) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*VendorConnection_Info_PointGrabInfo))
}

func (o *VendorConnection_Info_PointGrabInfo_TelemetrySubscription) GotenObjectExt() {}

func (o *VendorConnection_Info_PointGrabInfo_TelemetrySubscription) MakeFullFieldMask() *VendorConnection_Info_PointGrabInfo_TelemetrySubscription_FieldMask {
	return FullVendorConnection_Info_PointGrabInfo_TelemetrySubscription_FieldMask()
}

func (o *VendorConnection_Info_PointGrabInfo_TelemetrySubscription) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullVendorConnection_Info_PointGrabInfo_TelemetrySubscription_FieldMask()
}

func (o *VendorConnection_Info_PointGrabInfo_TelemetrySubscription) MakeDiffFieldMask(other *VendorConnection_Info_PointGrabInfo_TelemetrySubscription) *VendorConnection_Info_PointGrabInfo_TelemetrySubscription_FieldMask {
	if o == nil && other == nil {
		return &VendorConnection_Info_PointGrabInfo_TelemetrySubscription_FieldMask{}
	}
	if o == nil || other == nil {
		return FullVendorConnection_Info_PointGrabInfo_TelemetrySubscription_FieldMask()
	}

	res := &VendorConnection_Info_PointGrabInfo_TelemetrySubscription_FieldMask{}
	if o.GetSubscriptionId() != other.GetSubscriptionId() {
		res.Paths = append(res.Paths, &VendorConnectionInfoPointGrabInfoTelemetrySubscription_FieldTerminalPath{selector: VendorConnectionInfoPointGrabInfoTelemetrySubscription_FieldPathSelectorSubscriptionId})
	}
	if o.GetNotificationType() != other.GetNotificationType() {
		res.Paths = append(res.Paths, &VendorConnectionInfoPointGrabInfoTelemetrySubscription_FieldTerminalPath{selector: VendorConnectionInfoPointGrabInfoTelemetrySubscription_FieldPathSelectorNotificationType})
	}
	return res
}

func (o *VendorConnection_Info_PointGrabInfo_TelemetrySubscription) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*VendorConnection_Info_PointGrabInfo_TelemetrySubscription))
}

func (o *VendorConnection_Info_PointGrabInfo_TelemetrySubscription) Clone() *VendorConnection_Info_PointGrabInfo_TelemetrySubscription {
	if o == nil {
		return nil
	}
	result := &VendorConnection_Info_PointGrabInfo_TelemetrySubscription{}
	result.SubscriptionId = o.SubscriptionId
	result.NotificationType = o.NotificationType
	return result
}

func (o *VendorConnection_Info_PointGrabInfo_TelemetrySubscription) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *VendorConnection_Info_PointGrabInfo_TelemetrySubscription) Merge(source *VendorConnection_Info_PointGrabInfo_TelemetrySubscription) {
	o.SubscriptionId = source.GetSubscriptionId()
	o.NotificationType = source.GetNotificationType()
}

func (o *VendorConnection_Info_PointGrabInfo_TelemetrySubscription) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*VendorConnection_Info_PointGrabInfo_TelemetrySubscription))
}
