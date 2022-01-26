// Code generated by protoc-gen-goten-resource
// Resource change: VendorConnectionChange
// DO NOT EDIT!!!

package vendor_connection

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &iam_project.Project{}
	_ = &field_mask.FieldMask{}
)

func (c *VendorConnectionChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*VendorConnectionChange_Added_)
	return ok
}

func (c *VendorConnectionChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*VendorConnectionChange_Modified_)
	return ok
}

func (c *VendorConnectionChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*VendorConnectionChange_Current_)
	return ok
}

func (c *VendorConnectionChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*VendorConnectionChange_Removed_)
	return ok
}

func (c *VendorConnectionChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *VendorConnectionChange_Added_:
		return cType.Added.ViewIndex
	case *VendorConnectionChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *VendorConnectionChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *VendorConnectionChange_Removed_:
		return cType.Removed.ViewIndex
	case *VendorConnectionChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *VendorConnectionChange) GetVendorConnection() *VendorConnection {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *VendorConnectionChange_Added_:
		return cType.Added.VendorConnection
	case *VendorConnectionChange_Modified_:
		return cType.Modified.VendorConnection
	case *VendorConnectionChange_Current_:
		return cType.Current.VendorConnection
	case *VendorConnectionChange_Removed_:
		return nil
	}
	return nil
}

func (c *VendorConnectionChange) GetResource() gotenresource.Resource {
	return c.GetVendorConnection()
}

func (c *VendorConnectionChange) GetVendorConnectionName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *VendorConnectionChange_Added_:
		return cType.Added.VendorConnection.GetName()
	case *VendorConnectionChange_Modified_:
		return cType.Modified.Name
	case *VendorConnectionChange_Current_:
		return cType.Current.VendorConnection.GetName()
	case *VendorConnectionChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *VendorConnectionChange) GetRawName() gotenresource.Name {
	return c.GetVendorConnectionName()
}

func (c *VendorConnectionChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &VendorConnectionChange_Added_{
		Added: &VendorConnectionChange_Added{
			VendorConnection: snapshot.(*VendorConnection),
			ViewIndex:        int32(idx),
		},
	}
}

func (c *VendorConnectionChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &VendorConnectionChange_Modified_{
		Modified: &VendorConnectionChange_Modified{
			Name:              name.(*Name),
			VendorConnection:  snapshot.(*VendorConnection),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *VendorConnectionChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &VendorConnectionChange_Current_{
		Current: &VendorConnectionChange_Current{
			VendorConnection: snapshot.(*VendorConnection),
		},
	}
}

func (c *VendorConnectionChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &VendorConnectionChange_Removed_{
		Removed: &VendorConnectionChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
