package build_in

import (
	"encoding/json"
	"testing"
	//"fmt"
)

func mapToJson() {
	m := map[string]int64{
		"v_show":    1,
		"v_click":   2,
		"v_convt":   3,
		"n_v_show":  4,
		"n_v_click": 5,
		"n_v_convt": 6,
	}

	json.Marshal(m)
}

func structToJson() {
	extraStruct := struct {
		VivoShow  int64 `json:"v_show"`
		VClick    int64 `json:"v_click"`
		VConvt    int64 `json:"v_convt"`
		NVShow    int64 `json:"n_v_show"`
		NVClick   int64 `json:"nv_click"`
		NVConvert int64 `json:"n_v_convt"`
	}{
		1, 2, 3, 4, 5, 6,
	}

	json.Marshal(extraStruct)
	//fmt.Printf("%#v, err: %v\n", string(dat), err)
}

/*
func TestStructToJson(t *testing.T) {
	structToJson()
}
*/

func BenchmarkMap2Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapToJson()
	}
}

func BenchmarkStruct2Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		structToJson()
	}
}
