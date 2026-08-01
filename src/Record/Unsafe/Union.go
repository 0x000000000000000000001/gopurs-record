package Record_Unsafe_Union

import "gopurs/output/gopurs_runtime"

type r1 interface{}
type r2 interface{}

func _UnsafeUnionFn(r1, r2 interface{}) interface{} {
	rec1 := r1.(gopurs_runtime.Value)
	rec2 := r2.(gopurs_runtime.Value)
	m1 := gopurs_runtime.RecordToMap(rec1)
	m2 := gopurs_runtime.RecordToMap(rec2)
	
	m3 := make(map[string]gopurs_runtime.Value)
	
	for k, v := range m2 {
		m3[k] = v
	}
	for k, v := range m1 {
		m3[k] = v
	}
	
	return gopurs_runtime.Record(m3)
}
