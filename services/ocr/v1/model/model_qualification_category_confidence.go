package model

import (
	"encoding/json"

	"strings"
)

type QualificationCategoryConfidence struct {
	// 诚信考核信息（非必有，依赖对应从业资格证板式）。

	Category *float32 `json:"category,omitempty"`
	// 初次领证日期（非必有，依赖对应从业资格证板式）

	InitialIssueDate *float32 `json:"initial_issue_date,omitempty"`
	// 有效起始日期（非必有，依赖对应从业资格证板式）

	IssueDate *float32 `json:"issue_date,omitempty"`
	// 有效期至

	ExpiryDate *float32 `json:"expiry_date,omitempty"`
}

func (o QualificationCategoryConfidence) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QualificationCategoryConfidence struct{}"
	}

	return strings.Join([]string{"QualificationCategoryConfidence", string(data)}, " ")
}