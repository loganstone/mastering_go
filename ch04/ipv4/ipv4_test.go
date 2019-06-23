package ipv4

import (
	"net"
	"testing"
)

func TestFind(t *testing.T) {
	good := "192.168.0.2"
	trial := net.ParseIP(good)
	if trial == nil {
		t.Errorf("good: %s", good)
	}
	if trial.To4() == nil {
		t.Errorf("good: %s", good)
	}
	actual := Find(good)
	if actual != good {
		t.Errorf("good: %s, actual: %s", good, actual)
	}

	bad := "299.168.0.2"
	trial = net.ParseIP(bad)
	if trial != nil {
		t.Errorf("bad: %s", bad)
	}
	actual = Find(bad)
	if actual == bad {
		t.Errorf("bad: %s, actual: %s", bad, actual)
	}

}
