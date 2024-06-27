// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NodeExecutionInfo node execution info
// swagger:model NodeExecutionInfo
type NodeExecutionInfo struct {

	// 实验执行的id
	ExecRunID string `json:"exec_run_id,omitempty"`

	// 节点的id
	NodeID string `json:"node_id,omitempty"`

	// 节点的执行状态
	NodeStatus string `json:"node_status,omitempty"`

	// 训练任务的id，通过该id查询日志
	TrainingTaskID string `json:"training_task_id,omitempty"`
}

// Validate validates this node execution info
func (m *NodeExecutionInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeExecutionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeExecutionInfo) UnmarshalBinary(b []byte) error {
	var res NodeExecutionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
