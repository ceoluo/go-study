package singleton_test

import (
	hungrySingleton "github.com/go-study/src/design-pattern/singleton/hungry"
	lazySingleton "github.com/go-study/src/design-pattern/singleton/lazy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHungrySingleton(t *testing.T) {
	instance1 := hungrySingleton.GetSingletonInstance()
	instance2 := hungrySingleton.GetSingletonInstance()
	t.Logf("\ninstance1:%p\ninstance2:%p\n", instance1, instance2)
	assert.Equal(t, instance1, instance1)
}


func BenchmarkGetHungrySingleton(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			instance1 := hungrySingleton.GetSingletonInstance()
			instance2 := hungrySingleton.GetSingletonInstance()
			//b.Logf("\ninstance1:%p\ninstance2:%p\n", instance1, instance2)
			if instance1 != instance2 {
				b.Errorf("test fail")
			}
		}
	})
}


func TestGetLazySingleton(t *testing.T) {
	instance1 := lazySingleton.GetSingletonInstance()
	instance2 := lazySingleton.GetSingletonInstance()
	t.Logf("\ninstance1:%p\ninstance2:%p\n", instance1, instance2)
	assert.Equal(t, instance1, instance1)
}


func BenchmarkGetLazySingleton(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			instance1 := lazySingleton.GetSingletonInstance()
			instance2 := lazySingleton.GetSingletonInstance()
			//b.Logf("\ninstance1:%p\ninstance2:%p\n", instance1, instance2)
			if instance1 != instance2 {
				b.Errorf("test fail")
			}
		}
	})
}

/* go test -benchmem -bench="." -v
create a hungry singeleton instance
=== RUN   TestGetHungrySingleton
    singleton_test.go:14:
        instance1:0xd1eef0
        instance2:0xd1eef0
--- PASS: TestGetHungrySingleton (0.00s)
=== RUN   TestGetLazySingleton
create a lazy singeleton instance
    singleton_test.go:36:
        instance1:0xd1eef0
        instance2:0xd1eef0
--- PASS: TestGetLazySingleton (0.00s)
goos: windows
goarch: amd64
pkg: github.com/go-study/src/design-pattern/singleton
cpu: Intel(R) Core(TM) i7-9700 CPU @ 3.00GHz
BenchmarkGetHungrySingleton
BenchmarkGetHungrySingleton-8           1000000000               0.1909 ns/op          0 B/op          0 allocs/op
BenchmarkGetLazySingleton
BenchmarkGetLazySingleton-8             1000000000               0.4482 ns/op          0 B/op          0 allocs/op
PASS
 */