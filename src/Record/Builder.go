package Record_Builder

import "gopurs/output/gopurs_runtime"

func CopyRecord(r gopurs_runtime.Value) gopurs_runtime.Value {
	m := gopurs_runtime.RecordToMap(r)
	m2 := make(map[string]gopurs_runtime.Value, len(m))
	for k, v := range m {
		m2[k] = v
	}
	return gopurs_runtime.Record(m2)
}

func UnsafeInsert(l string, a gopurs_runtime.Value, r gopurs_runtime.Value) gopurs_runtime.Value {
	m := gopurs_runtime.RecordToMap(r)
	m[l] = a
	return gopurs_runtime.Record(m)
}

func UnsafeModify(l string, f gopurs_runtime.Value, r gopurs_runtime.Value) gopurs_runtime.Value {
	m := gopurs_runtime.RecordToMap(r)
	m[l] = gopurs_runtime.Apply(f, m[l])
	return gopurs_runtime.Record(m)
}

func UnsafeDelete(l string, r gopurs_runtime.Value) gopurs_runtime.Value {
	m := gopurs_runtime.RecordToMap(r)
	delete(m, l)
	return gopurs_runtime.Record(m)
}

func UnsafeRename(l1 string, l2 string, r gopurs_runtime.Value) gopurs_runtime.Value {
	m := gopurs_runtime.RecordToMap(r)
	m[l2] = m[l1]
	delete(m, l1)
	return gopurs_runtime.Record(m)
}
