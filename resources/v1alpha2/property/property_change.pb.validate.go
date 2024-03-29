// Code generated by protoc-gen-goten-validate
// File: workplace/proto/v1alpha2/property_change.proto
// DO NOT EDIT!!!

package property

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	gotenvalidate "github.com/cloudwan/goten-sdk/runtime/validate"
)

// proto imports
import (
	area "github.com/cloudwan/workplace-sdk/resources/v1alpha2/area"
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha2/building"
	device "github.com/cloudwan/workplace-sdk/resources/v1alpha2/device"
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha2/floor"
	site "github.com/cloudwan/workplace-sdk/resources/v1alpha2/site"
	zone "github.com/cloudwan/workplace-sdk/resources/v1alpha2/zone"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Errorf
	_ = net.ParseIP
	_ = regexp.Match
	_ = strings.Split
	_ = time.Now
	_ = utf8.RuneCountInString
	_ = url.Parse
	_ = durationpb.Duration{}
	_ = timestamppb.Timestamp{}
	_ = gotenvalidate.NewValidationError
)

// make sure we're using proto imports
var (
	_ = &field_mask.FieldMask{}
	_ = &area.Area{}
	_ = &building.Building{}
	_ = &device.Device{}
	_ = &floor.Floor{}
	_ = &site.Site{}
	_ = &zone.Zone{}
)

func (obj *PropertyChange) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.ChangeType.(type) {
	case *PropertyChange_Added_:
		if subobj, ok := interface{}(opt.Added).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("PropertyChange", "added", opt.Added, "nested object validation failed", err)
			}
		}
	case *PropertyChange_Modified_:
		if subobj, ok := interface{}(opt.Modified).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("PropertyChange", "modified", opt.Modified, "nested object validation failed", err)
			}
		}
	case *PropertyChange_Current_:
		if subobj, ok := interface{}(opt.Current).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("PropertyChange", "current", opt.Current, "nested object validation failed", err)
			}
		}
	case *PropertyChange_Removed_:
		if subobj, ok := interface{}(opt.Removed).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("PropertyChange", "removed", opt.Removed, "nested object validation failed", err)
			}
		}
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *PropertyChange_Added) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Property).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Added", "property", obj.Property, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *PropertyChange_Modified) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Property).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Modified", "property", obj.Property, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *PropertyChange_Current) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Property).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Current", "property", obj.Property, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *PropertyChange_Removed) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
