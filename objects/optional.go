package objects

// Optional 可选结构
type Optional struct {
	val interface{}
}

// OptionalOf 构造Optional
func OptionalOf(val interface{}) *Optional {
	return &Optional{val: val}
}

// OptionalOfNil 设置nil Optinal
func OptionalOfNil() *Optional {
	return OptionalOf(nil)
}

// Get 获取Optinal值
func (optinal *Optional) Of() interface{} {
	return optinal.val
}

// IsPresent Optional值是否nil
func (optinal *Optional) IsPresent() bool {
	return optinal.val != nil
}

// IfPresent  不支持泛型 暂不实现
func (optinal *Optional) IfPresent(action *Consumer) {

}
