package splitter

import (
	"errors"
	"math"
	"reflect"
)

type Clause struct {
	Key    string
	Method string
	Val    interface{}
}

func Where[T any](data []T, clauses []Clause) ([]T, error) {
	var retVal []T

	for c := 0; c < len(clauses); c++ {
		retVal = []T{}
		for i := 0; i < len(data); i++ {
			r := reflect.ValueOf(data[i])
			f := reflect.Indirect(r).FieldByName(clauses[c].Key)
			if clauses[c].Method == "==" {
				if f.Interface() == clauses[c].Val {
					retVal = append(retVal, data[i])
				}
			} else if clauses[c].Method == ">" {
				valFloat, err := getFloat(clauses[c].Val)
				fFloat, err := getFloat(f.Interface())
				if err == nil {
					if fFloat > valFloat {
						retVal = append(retVal, data[i])
					}
				} else {
					return nil, err
				}
			} else if clauses[c].Method == ">=" {
				valFloat, err := getFloat(clauses[c].Val)
				fFloat, err := getFloat(f.Interface())
				if err == nil {
					if fFloat >= valFloat {
						retVal = append(retVal, data[i])
					}
				} else {
					return nil, err
				}
			} else if clauses[c].Method == "<" {
				valFloat, err := getFloat(clauses[c].Val)
				fFloat, err := getFloat(f.Interface())
				if err == nil {
					if fFloat < valFloat {
						retVal = append(retVal, data[i])
					}
				} else {
					return nil, err
				}
			} else if clauses[c].Method == "<=" {
				valFloat, err := getFloat(clauses[c].Val)
				fFloat, err := getFloat(f.Interface())
				if err == nil {
					if fFloat <= valFloat {
						retVal = append(retVal, data[i])
					}
				} else {
					return nil, err
				}
			} else if clauses[c].Method == "!=" {
				if f.Interface() != clauses[c].Val {
					retVal = append(retVal, data[i])
				}
			}
		}
		data = retVal
	}

	return retVal, nil
}

var errUnexpectedType = errors.New("Non-numeric type could not be converted to float")

func getFloat(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	default:
		return math.NaN(), errUnexpectedType
	}
}
