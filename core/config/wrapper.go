package config

import (
	"time"

	"github.com/knadh/koanf"
)

func UnmarshalWithPath(path string, o interface{}) error {
	return instance.UnmarshalWithConf(path, &o, koanf.UnmarshalConf{Tag: "config"})
}

func Unmarshal(o interface{}) error {
	return instance.UnmarshalWithConf("", &o, koanf.UnmarshalConf{Tag: "config"})
}

func Exists(path string) bool {
	return instance.Exists(path)
}

func Int64(path string) int64 {
	return instance.Int64(path)
}

func Int64s(path string) []int64 {
	return instance.Int64s(path)
}

func Int64Map(path string) map[string]int64 {
	return instance.Int64Map(path)
}

func Int(path string) int {
	return instance.Int(path)
}

func Ints(path string) []int {
	return instance.Ints(path)
}

func IntMap(path string) map[string]int {
	return instance.IntMap(path)
}

func Float64(path string) float64 {
	return instance.Float64(path)
}

func Float64s(path string) []float64 {
	return instance.Float64s(path)
}

func Float64Map(path string) map[string]float64 {
	return instance.Float64Map(path)
}

func Duration(path string) time.Duration {
	return instance.Duration(path)
}

func Time(path, layout string) time.Time {
	return instance.Time(path, layout)
}

func String(path string) string {
	return instance.String(path)
}

func Strings(path string) []string {
	return instance.Strings(path)
}

func StringMap(path string) map[string]string {
	return instance.StringMap(path)
}

func Bytes(path string) []byte {
	return instance.Bytes(path)
}

func Bool(path string) bool {
	return instance.Bool(path)
}

func Bools(path string) []bool {
	return instance.Bools(path)
}

func BoolMap(path string) map[string]bool {
	return instance.BoolMap(path)
}
