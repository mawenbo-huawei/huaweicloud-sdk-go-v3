package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 隐私屏蔽响应体
type PrivacyResponseBody struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 隐私屏蔽规则应用的url

	Url *string `json:"url,omitempty"`
	// 屏蔽字段

	Category *PrivacyResponseBodyCategory `json:"category,omitempty"`
	// 屏蔽字段名

	Index *string `json:"index,omitempty"`
}

func (o PrivacyResponseBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PrivacyResponseBody struct{}"
	}

	return strings.Join([]string{"PrivacyResponseBody", string(data)}, " ")
}

type PrivacyResponseBodyCategory struct {
	value string
}

type PrivacyResponseBodyCategoryEnum struct {
	PARAMS PrivacyResponseBodyCategory
	COOKIE PrivacyResponseBodyCategory
	HEADER PrivacyResponseBodyCategory
	FORM   PrivacyResponseBodyCategory
}

func GetPrivacyResponseBodyCategoryEnum() PrivacyResponseBodyCategoryEnum {
	return PrivacyResponseBodyCategoryEnum{
		PARAMS: PrivacyResponseBodyCategory{
			value: "params",
		},
		COOKIE: PrivacyResponseBodyCategory{
			value: "cookie",
		},
		HEADER: PrivacyResponseBodyCategory{
			value: "header",
		},
		FORM: PrivacyResponseBodyCategory{
			value: "form",
		},
	}
}

func (c PrivacyResponseBodyCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *PrivacyResponseBodyCategory) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}