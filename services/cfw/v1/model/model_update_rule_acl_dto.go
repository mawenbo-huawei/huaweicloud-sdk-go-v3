package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateRuleAclDto UpdateRuleAclDto
type UpdateRuleAclDto struct {

	// 地址类型，0表示ipv4，1表示ipv6
	AddressType *UpdateRuleAclDtoAddressType `json:"address_type,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 方向：0表示外到内，1表示内到外,规则type=0（互联网规则）或者type= 2（nat规则）时方向值必填
	Direction *UpdateRuleAclDtoDirection `json:"direction,omitempty"`

	// 规则动作，0表示允许通行（permit），1表示拒绝通行（deny）
	ActionType *UpdateRuleAclDtoActionType `json:"action_type,omitempty"`

	// 规则启用状态，0表示禁用，1表示启用
	Status *int32 `json:"status,omitempty"`

	// 规则应用列表，规则应用类型包括：“HTTP”，\"HTTPS\"，\"TLS1\"，“DNS”，“SSH”，“MYSQL”，“SMTP”，“RDP”，“RDPS”，“VNC”，“POP3”，“IMAP4”，“SMTPS”，“POP3S”，“FTPS”，“ANY”,“BGP”等。
	Applications *[]string `json:"applications,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 长连接时长对应小时
	LongConnectTimeHour *int64 `json:"long_connect_time_hour,omitempty"`

	// 长连接时长对应分钟
	LongConnectTimeMinute *int64 `json:"long_connect_time_minute,omitempty"`

	// 长连接时长秒
	LongConnectTimeSecond *int64 `json:"long_connect_time_second,omitempty"`

	// 长连接时长
	LongConnectTime *int64 `json:"long_connect_time,omitempty"`

	// 是否支持长连接，0表示不支持，1表示支持
	LongConnectEnable *UpdateRuleAclDtoLongConnectEnable `json:"long_connect_enable,omitempty"`

	Source *RuleAddressDto `json:"source,omitempty"`

	Destination *RuleAddressDto `json:"destination,omitempty"`

	Service *RuleServiceDto `json:"service,omitempty"`

	// 规则类型，0：互联网规则，1：vpc规则，2：nat规则
	Type *UpdateRuleAclDtoType `json:"type,omitempty"`

	Tag *TagsVo `json:"tag,omitempty"`
}

func (o UpdateRuleAclDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleAclDto struct{}"
	}

	return strings.Join([]string{"UpdateRuleAclDto", string(data)}, " ")
}

type UpdateRuleAclDtoAddressType struct {
	value int32
}

type UpdateRuleAclDtoAddressTypeEnum struct {
	E_0 UpdateRuleAclDtoAddressType
	E_1 UpdateRuleAclDtoAddressType
}

func GetUpdateRuleAclDtoAddressTypeEnum() UpdateRuleAclDtoAddressTypeEnum {
	return UpdateRuleAclDtoAddressTypeEnum{
		E_0: UpdateRuleAclDtoAddressType{
			value: 0,
		}, E_1: UpdateRuleAclDtoAddressType{
			value: 1,
		},
	}
}

func (c UpdateRuleAclDtoAddressType) Value() int32 {
	return c.value
}

func (c UpdateRuleAclDtoAddressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleAclDtoAddressType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type UpdateRuleAclDtoDirection struct {
	value int32
}

type UpdateRuleAclDtoDirectionEnum struct {
	E_0 UpdateRuleAclDtoDirection
	E_1 UpdateRuleAclDtoDirection
}

func GetUpdateRuleAclDtoDirectionEnum() UpdateRuleAclDtoDirectionEnum {
	return UpdateRuleAclDtoDirectionEnum{
		E_0: UpdateRuleAclDtoDirection{
			value: 0,
		}, E_1: UpdateRuleAclDtoDirection{
			value: 1,
		},
	}
}

func (c UpdateRuleAclDtoDirection) Value() int32 {
	return c.value
}

func (c UpdateRuleAclDtoDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleAclDtoDirection) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type UpdateRuleAclDtoActionType struct {
	value int32
}

type UpdateRuleAclDtoActionTypeEnum struct {
	E_0 UpdateRuleAclDtoActionType
	E_1 UpdateRuleAclDtoActionType
}

func GetUpdateRuleAclDtoActionTypeEnum() UpdateRuleAclDtoActionTypeEnum {
	return UpdateRuleAclDtoActionTypeEnum{
		E_0: UpdateRuleAclDtoActionType{
			value: 0,
		}, E_1: UpdateRuleAclDtoActionType{
			value: 1,
		},
	}
}

func (c UpdateRuleAclDtoActionType) Value() int32 {
	return c.value
}

func (c UpdateRuleAclDtoActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleAclDtoActionType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type UpdateRuleAclDtoLongConnectEnable struct {
	value int32
}

type UpdateRuleAclDtoLongConnectEnableEnum struct {
	E_0 UpdateRuleAclDtoLongConnectEnable
	E_1 UpdateRuleAclDtoLongConnectEnable
}

func GetUpdateRuleAclDtoLongConnectEnableEnum() UpdateRuleAclDtoLongConnectEnableEnum {
	return UpdateRuleAclDtoLongConnectEnableEnum{
		E_0: UpdateRuleAclDtoLongConnectEnable{
			value: 0,
		}, E_1: UpdateRuleAclDtoLongConnectEnable{
			value: 1,
		},
	}
}

func (c UpdateRuleAclDtoLongConnectEnable) Value() int32 {
	return c.value
}

func (c UpdateRuleAclDtoLongConnectEnable) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleAclDtoLongConnectEnable) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type UpdateRuleAclDtoType struct {
	value int32
}

type UpdateRuleAclDtoTypeEnum struct {
	E_0 UpdateRuleAclDtoType
	E_1 UpdateRuleAclDtoType
	E_2 UpdateRuleAclDtoType
}

func GetUpdateRuleAclDtoTypeEnum() UpdateRuleAclDtoTypeEnum {
	return UpdateRuleAclDtoTypeEnum{
		E_0: UpdateRuleAclDtoType{
			value: 0,
		}, E_1: UpdateRuleAclDtoType{
			value: 1,
		}, E_2: UpdateRuleAclDtoType{
			value: 2,
		},
	}
}

func (c UpdateRuleAclDtoType) Value() int32 {
	return c.value
}

func (c UpdateRuleAclDtoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleAclDtoType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
