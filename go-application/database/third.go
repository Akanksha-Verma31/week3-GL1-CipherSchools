package database

import "log"

var key int

func GetKey() {
	key = 9
}

type val struct {
	i int
}

func (v *val) GetVal() int {
	v.i = 10
	return v.i
}
func get() {
	vv := val{}
	vv.GetVal()
	log.Println(vv)
	ww := &val{}
	ww.GetVal()
	log.Println(ww.i)
}
