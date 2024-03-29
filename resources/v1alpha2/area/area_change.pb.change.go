// Code generated by protoc-gen-goten-resource
// Resource change: AreaChange
// DO NOT EDIT!!!

package area

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha2/floor"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &field_mask.FieldMask{}
	_ = &floor.Floor{}
)

func (c *AreaChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AreaChange_Added_)
	return ok
}

func (c *AreaChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AreaChange_Modified_)
	return ok
}

func (c *AreaChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AreaChange_Current_)
	return ok
}

func (c *AreaChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AreaChange_Removed_)
	return ok
}

func (c *AreaChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *AreaChange_Added_:
		return cType.Added.ViewIndex
	case *AreaChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *AreaChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *AreaChange_Removed_:
		return cType.Removed.ViewIndex
	case *AreaChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *AreaChange) GetArea() *Area {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *AreaChange_Added_:
		return cType.Added.Area
	case *AreaChange_Modified_:
		return cType.Modified.Area
	case *AreaChange_Current_:
		return cType.Current.Area
	case *AreaChange_Removed_:
		return nil
	}
	return nil
}

func (c *AreaChange) GetRawResource() gotenresource.Resource {
	return c.GetArea()
}

func (c *AreaChange) GetAreaName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *AreaChange_Added_:
		return cType.Added.Area.GetName()
	case *AreaChange_Modified_:
		return cType.Modified.Name
	case *AreaChange_Current_:
		return cType.Current.Area.GetName()
	case *AreaChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *AreaChange) GetRawName() gotenresource.Name {
	return c.GetAreaName()
}

func (c *AreaChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &AreaChange_Added_{
		Added: &AreaChange_Added{
			Area:      snapshot.(*Area),
			ViewIndex: int32(idx),
		},
	}
}

func (c *AreaChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &AreaChange_Modified_{
		Modified: &AreaChange_Modified{
			Name:              name.(*Name),
			Area:              snapshot.(*Area),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *AreaChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &AreaChange_Current_{
		Current: &AreaChange_Current{
			Area: snapshot.(*Area),
		},
	}
}

func (c *AreaChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &AreaChange_Removed_{
		Removed: &AreaChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
