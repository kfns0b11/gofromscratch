package proxy

import "testing"

func TestProxy(t *testing.T) {
	sp := &StationProxy{
		&Station{
			stock: 1,
		},
	}

	sp.sell("lili")
	sp.sell("lucky")
}
