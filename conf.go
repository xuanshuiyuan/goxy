// @Author xuanshuiyuan 2024/9/11 17:07:00
package goxy

type MI8S map[int8]string

type OptionFormatKV struct {
	Key   int8   `json:"key"`
	Value string `json:"value"`
}

//选项排序
type OptionSortList []OptionFormatKV

func (o OptionSortList) Len() int {
	return len(o)
}

func (o OptionSortList) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o OptionSortList) Less(i, j int) bool {
	return o[i].Key < o[j].Key
}
