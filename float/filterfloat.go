package float

import (
	"math"
	"reflect"
)

// FilterFloat Filtering Float
func FilterFloat(t interface{}, prec int) {
	if reflect.TypeOf(t).Elem().Kind() == reflect.Slice {
		s := reflect.ValueOf(t).Elem()

		for x := 0; x < s.Len(); x++ {
			v := reflect.Indirect(s.Index(x))
			assignFloat(v, prec)
		}
	} else {
		var v reflect.Value
		if reflect.TypeOf(t).Elem().Kind() == reflect.Struct {
			v = reflect.ValueOf(t).Elem()
		} else if reflect.TypeOf(t).Elem().Kind() == reflect.Ptr {
			v = reflect.ValueOf(t).Elem().Elem()
		}
		assignFloat(v, prec)
	}
}

func assignFloat(v reflect.Value, prec int) {
	for i, n := 0, v.NumField(); i < n; i++ {
		f := v.Field(i)
		st, n := checkType(f.Interface(), prec)
		if st {
			v.Field(i).SetFloat(n)
		}
	}
}

func checkType(t interface{}, prec int) (st bool, n float64) {
	switch t.(type) {
	case float64:
		st = true
		n = toFixedRoundDigits(t.(float64), prec)
	default:
		st = false
	}

	return
}

func toFixedRoundDigits(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(int64(num*output+0.5)) / output
}
