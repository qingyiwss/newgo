package main

import "fmt"

type DianShan struct {
	key    bool
	shanye string
}

func (dianshan *DianShan) open(key bool) string {
	if key {
		dianshan.shanye = "转起来了weisongshan"

	} else {
		dianshan.shanye = "不转了50e7f55bd7fb897bfa741acf8d8d65f72d479289"
	}
	return dianshan.shanye
}

func main() {

	ds := DianShan{false, ""}
	s := ds.open(true)
	fmt.Println(s)

}
