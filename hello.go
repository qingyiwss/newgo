package main

import "fmt"

type DianShan struct {
	key    bool
	shanye string
}

func (dianshan *DianShan) open(key bool) string {
	if key {
		dianshan.shanye = "转起来了"

	} else {
		dianshan.shanye = "不转了"
	}
	return dianshan.shanye
}

func main() {

	ds := DianShan{false, ""}
	s := ds.open(true)
	fmt.Println(s)

}
