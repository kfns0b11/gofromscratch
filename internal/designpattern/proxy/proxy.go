// The proxy class holds the object of the proxied class and implements
// the same interface as the object of the proxied class
package proxy

import "fmt"

type Seller interface {
	sell(name string)
}

type Station struct {
	stock int
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("%s purchase one ticket; remain: %d\n", name, station.stock)
	} else {
		fmt.Println("no ticket")
	}
}

type StationProxy struct {
	station *Station
}

func (proxy *StationProxy) sell(name string) {
	if proxy.station.stock > 0 {
		proxy.station.stock--
		fmt.Printf("proxy: %s purchase one ticket; remain: %d\n", name, proxy.station.stock)
	} else {
		fmt.Println("proxy: no ticket")
	}
}
