package _map

import (
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[string]int{"key1": 1, "key2": 2, "key3": 3}
	t.Log(m1)
	t.Logf("len m1=%d", len(m1))

	m2 := map[string]int{}
	m2["key1"] = 1

	m3 := make(map[string]int, 10)
	t.Log(m3)
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[string]int{}
	// 默认值为0
	t.Log(m1["key"])

	m1["key"] = 1
	t.Log(m1["key"])
	if v, ok := m1["key"]; ok {
		t.Logf("Key key's value is %d", v)
	} else {
		t.Log("Key Key is not existing!")
	}
}

func TestTavelMap(t *testing.T) {
	m1 := map[string]int{"key1": 1, "key2": 2, "key3": 3}
	for k, v := range m1 {
		t.Logf("m1[%s] = %d", k, v)
	}
}
