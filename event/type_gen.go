// Code generated by "enumer -type=Type -trimprefix Type -output type_gen.go"; DO NOT EDIT.

package event

import (
	"fmt"
)

const _TypeName = "InitTestStartTestEndIterStartIterEndExit"

var _TypeIndex = [...]uint8{0, 4, 13, 20, 29, 36, 40}

func (i Type) String() string {
	i -= 1
	if i >= Type(len(_TypeIndex)-1) {
		return fmt.Sprintf("Type(%d)", i+1)
	}
	return _TypeName[_TypeIndex[i]:_TypeIndex[i+1]]
}

var _TypeValues = []Type{1, 2, 3, 4, 5, 6}

var _TypeNameToValueMap = map[string]Type{
	_TypeName[0:4]:   1,
	_TypeName[4:13]:  2,
	_TypeName[13:20]: 3,
	_TypeName[20:29]: 4,
	_TypeName[29:36]: 5,
	_TypeName[36:40]: 6,
}

// TypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TypeString(s string) (Type, error) {
	if val, ok := _TypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Type values", s)
}

// TypeValues returns all values of the enum
func TypeValues() []Type {
	return _TypeValues
}

// IsAType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Type) IsAType() bool {
	for _, v := range _TypeValues {
		if i == v {
			return true
		}
	}
	return false
}