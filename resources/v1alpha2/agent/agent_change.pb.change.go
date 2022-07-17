// Code generated by protoc-gen-goten-resource
// Resource change: AgentChange
// DO NOT EDIT!!!

package agent

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
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

func (c *AgentChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AgentChange_Added_)
	return ok
}

func (c *AgentChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AgentChange_Modified_)
	return ok
}

func (c *AgentChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AgentChange_Current_)
	return ok
}

func (c *AgentChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AgentChange_Removed_)
	return ok
}

func (c *AgentChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *AgentChange_Added_:
		return cType.Added.ViewIndex
	case *AgentChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *AgentChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *AgentChange_Removed_:
		return cType.Removed.ViewIndex
	case *AgentChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *AgentChange) GetAgent() *Agent {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *AgentChange_Added_:
		return cType.Added.Agent
	case *AgentChange_Modified_:
		return cType.Modified.Agent
	case *AgentChange_Current_:
		return cType.Current.Agent
	case *AgentChange_Removed_:
		return nil
	}
	return nil
}

func (c *AgentChange) GetRawResource() gotenresource.Resource {
	return c.GetAgent()
}

func (c *AgentChange) GetAgentName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *AgentChange_Added_:
		return cType.Added.Agent.GetName()
	case *AgentChange_Modified_:
		return cType.Modified.Name
	case *AgentChange_Current_:
		return cType.Current.Agent.GetName()
	case *AgentChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *AgentChange) GetRawName() gotenresource.Name {
	return c.GetAgentName()
}

func (c *AgentChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &AgentChange_Added_{
		Added: &AgentChange_Added{
			Agent:     snapshot.(*Agent),
			ViewIndex: int32(idx),
		},
	}
}

func (c *AgentChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &AgentChange_Modified_{
		Modified: &AgentChange_Modified{
			Name:              name.(*Name),
			Agent:             snapshot.(*Agent),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *AgentChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &AgentChange_Current_{
		Current: &AgentChange_Current{
			Agent: snapshot.(*Agent),
		},
	}
}

func (c *AgentChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &AgentChange_Removed_{
		Removed: &AgentChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
