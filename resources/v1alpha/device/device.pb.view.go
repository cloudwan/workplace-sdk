// Code generated by protoc-gen-goten-resource
// Resource: Device
// DO NOT EDIT!!!

package device

import (
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	agent "github.com/cloudwan/workplace-sdk/resources/v1alpha/agent"
	area "github.com/cloudwan/workplace-sdk/resources/v1alpha/area"
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha/building"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha/common"
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha/floor"
	site "github.com/cloudwan/workplace-sdk/resources/v1alpha/site"
	zone "github.com/cloudwan/workplace-sdk/resources/v1alpha/zone"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// ensure the imports are used
var (
	_ = fieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &timestamp.Timestamp{}
	_ = &agent.Agent{}
	_ = &area.Area{}
	_ = &workplace_common.BACNetEntity{}
	_ = &building.Building{}
	_ = &floor.Floor{}
	_ = &site.Site{}
	_ = &zone.Zone{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *Device_FieldMask) *Device_FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_NAME:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name")
		break
	default:
		return extraMask
	}
	if extraMask != nil {
		protoFieldMask.Paths = append(protoFieldMask.Paths, extraMask.ToProtoFieldMask().Paths...)
	}
	res := &Device_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
