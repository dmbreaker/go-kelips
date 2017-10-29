package kelips

import (
	"fmt"
	"testing"
)

func Test_InmemTuples(t *testing.T) {
	ft := NewInmemTuples()

	for i := 0; i < 10; i++ {
		ft.Add([]byte(fmt.Sprintf("filetuple%d", i)), NewHost("127.0.0.1", 1234+uint16(i)))
	}

	for i := 0; i < 5; i++ {
		ft.Add([]byte(fmt.Sprintf("filetuple%d", i)), NewHost("127.0.0.1", 21234))
	}
	ft.Add([]byte("filetuple0"), NewHost("127.0.0.1", 21234))

	for i := 0; i < 10; i++ {
		hosts := ft.Get([]byte(fmt.Sprintf("filetuple%d", i)))
		if hosts == nil || len(hosts) == 0 {
			t.Errorf("not found filetuple%d", i)
		}
	}

	for i := 0; i < 5; i++ {
		hosts := ft.Get([]byte(fmt.Sprintf("filetuple%d", i)))
		if len(hosts) != 2 {
			t.Fatal("should have 2 hosts")
		}
	}

	for i := 0; i < 5; i++ {
		ft.Del([]byte(fmt.Sprintf("filetuple%d", i)), NewHost("127.0.0.1", 21234))
	}

	for i := 0; i < 10; i++ {
		hosts := ft.Get([]byte(fmt.Sprintf("filetuple%d", i)))
		if len(hosts) != 1 {
			t.Errorf("should have=1 got=%d", len(hosts))
		}
	}

	h := NewHost("127.0.0.1", 11234)
	for i := 0; i < 5; i++ {
		ft.Add([]byte(fmt.Sprintf("filetuple%d", i)), h)
	}

	if !ft.ExpireHost(h) {
		t.Fatal("should have removed")
	}

	for i := 0; i < 10; i++ {
		hosts := ft.Get([]byte(fmt.Sprintf("filetuple%d", i)))
		if len(hosts) != 1 {
			t.Errorf("should have=1 got=%d", len(hosts))
		}
	}
}
