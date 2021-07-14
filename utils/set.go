package utils

import "github.com/bytedance/sonic"

type StringSet map[string]struct{}

// NewSetFromArray 根据array生成一个set结构的数据集合
func NewSetFromArray(arr []string) StringSet {
	arrLen := len(arr)
	set := make(StringSet, arrLen)
	for _, s := range arr {
		set[s] = struct{}{}
	}
	return set
}

// UnmarshalJSON 将array结构的Json数据反序列化为set结构
func (s StringSet) UnmarshalJSON(data []byte) (err error) {
	jsonNode, err := sonic.Get(data)
	if err != nil {
		return
	}

	arr := jsonNode.Array()
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		s[arr[i].(string)] = struct{}{}
	}
	return
}

// Intersection 与传入字符串数组是否存在交集
func (s StringSet) Intersection(arr []string) bool {
	for _, str := range arr {
		if _, ok := s[str]; ok {
			return true
		}
	}
	return false
}
