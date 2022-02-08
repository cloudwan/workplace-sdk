// Code generated by protoc-gen-goten-validate
// File: workplace/proto/v1alpha2/vendor_connection_custom.proto
// DO NOT EDIT!!!

package vendor_connection_client

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
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha2/common"
	vendor_connection "github.com/cloudwan/workplace-sdk/resources/v1alpha2/vendor_connection"
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
	_ = &workplace_common.BBox{}
	_ = &vendor_connection.PointGrab{}
)

func (obj *PointGrabTelemetryNotifyRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.PcPositions {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("PointGrabTelemetryNotifyRequest", "pcPositions", obj.PcPositions[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
