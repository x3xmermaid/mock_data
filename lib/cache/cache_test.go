package cache_test

import (
	ncache "ketitik/netmonk/mock-app-data/lib/cache"
	"testing"
)

type TestStruct struct {
	Data string
}

func TestCache(t *testing.T) {
	duration := 30
	netmonkCache := ncache.NewNMCache(duration)

	t.Run("test race", func(t *testing.T) {
		test := TestStruct{
			Data: "test",
		}
		go netmonkCache.Put("test", test)
		go netmonkCache.Put("test", test)
		netmonkCache.GetValue("test")
		go netmonkCache.Put("test", test)
		go netmonkCache.Put("test", test)
		netmonkCache.GetValue("test")
		netmonkCache.Put("test", test)
		_, err := netmonkCache.GetValue("test")
		if err != nil {
			t.Errorf("Get cache should not be error but have %v", err)
		}
	})

	t.Run("Cache OK", func(t *testing.T) {
		test := TestStruct{
			Data: "test",
		}

		err := netmonkCache.Put("test", test)
		if err != nil {
			t.Errorf("Put cache should not be error")
		}

		_, err = netmonkCache.GetValue("test")
		if err != nil {
			t.Errorf("Get cache should not be error")
		}

		isAvailable := netmonkCache.IsAvailable("test")
		if !isAvailable {
			t.Errorf("Cache should be available")
		}
	})

	t.Run("Put NOK", func(t *testing.T) {
		test := make(chan int)

		err := netmonkCache.Put("test", test)
		if err == nil {
			t.Errorf("Put cache should be error")
		}
	})

	t.Run("Get NOK", func(t *testing.T) {
		_, err := netmonkCache.GetValue("wrong")
		if err == nil {
			t.Errorf("Get cache should be error")
		}

		isAvailable := netmonkCache.IsAvailable("wrong")
		if isAvailable {
			t.Errorf("Cache should not be available")
		}
	})
}
