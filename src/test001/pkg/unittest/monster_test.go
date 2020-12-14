package unittest

import "testing"

//数据存序列化存入文件，在反序列化去除对比
func TestStore(t *testing.T) {
	path := "C:\\Users\\yw\\Desktop\\monster.txt"
	m := &Monster{
		Name:  "你好",
		Age:   32,
		Skill: "咏春",
	}
	mon := m
	content, errStore := m.Store(path)
	if errStore != nil {
		t.Fatalf("存储失败：%v", errStore)
	}
	t.Logf("存储成功,信息为：%v", content)

	m.Restore(path)
	//mon.Name="李四"
	if mon.Name != m.Name {
		t.Fatalf("反序列化失败，原先值为：%v,现在值为：%v", m, mon)
	}
	t.Logf("反序列化成功")
}
