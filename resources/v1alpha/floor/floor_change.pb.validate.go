// Code generated by protoc-gen-goten-validate
// File: workplace/proto/v1alpha/floor_change.proto
// DO NOT EDIT!!!

package floor

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
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha/building"
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
	_ = &building.Building{}
)

func (obj *FloorChange) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.ChangeType.(type) {
	case *FloorChange_Added_:
		if subobj, ok := interface{}(opt.Added).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("FloorChange", "added", opt.Added, "nested object validation failed", err)
			}
		}
	case *FloorChange_Modified_:
		if subobj, ok := interface{}(opt.Modified).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("FloorChange", "modified", opt.Modified, "nested object validation failed", err)
			}
		}
	case *FloorChange_Current_:
		if subobj, ok := interface{}(opt.Current).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("FloorChange", "current", opt.Current, "nested object validation failed", err)
			}
		}
	case *FloorChange_Removed_:
		if subobj, ok := interface{}(opt.Removed).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("FloorChange", "removed", opt.Removed, "nested object validation failed", err)
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
func (obj *FloorChange_Added) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Floor).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Added", "floor", obj.Floor, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *FloorChange_Modified) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Floor).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Modified", "floor", obj.Floor, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *FloorChange_Current) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Floor).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Current", "floor", obj.Floor, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *FloorChange_Removed) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
