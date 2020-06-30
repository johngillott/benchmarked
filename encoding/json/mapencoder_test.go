package json

import (
	"encoding/json"
	"testing"

	"github.com/bet365/jingo"
	jsoniterator "github.com/json-iterator/go"
	segmentio "github.com/segmentio/encoding/json"
	"github.com/wI2L/jettison"
)

// go test -benchmem -run=^$ github.com/johngillott/gobenchmark/encoding/json -bench BenchmarkMarshal -count=5 | tee bench.txt && benchstat bench.txt

func BenchmarkMarshalEncoder_map_string_string_jingo_sorted(b *testing.B) {
	b.ReportAllocs()

	m := mapPayload()

	enc := jingo.NewMapEncoder(map[string]string{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		buf := jingo.NewBufferFromPool()
		enc.Marshal(&m, buf)

		b.SetBytes(int64(len(buf.Bytes)))

		buf.ReturnToPool()
	}
}

func BenchmarkMarshalEncoder_map_string_string_jingo_unsorted(b *testing.B) {
	b.ReportAllocs()

	m := mapPayload()

	cfg := jingo.Config{}
	cfg.SetSortMapKeys(false)
	enc := jingo.NewMapEncoderWithConfig(map[string]string{}, cfg)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		buf := jingo.NewBufferFromPool()
		enc.Marshal(&m, buf)
		b.SetBytes(int64(len(buf.Bytes)))

		buf.ReturnToPool()
	}
}

func BenchmarkMarshalEncoder_map_string_string_segmentio_sorted(b *testing.B) {

	b.ReportAllocs()

	m := mapPayload()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		buf := jingo.NewBufferFromPool()

		enc := segmentio.NewEncoder(buf)

		if err := enc.Encode(&m); err != nil {
			b.Fatal("Encode:", err)
		}
		b.SetBytes(int64(len(buf.Bytes)))

		buf.ReturnToPool()
	}
}

func BenchmarkMarshalEncoder_map_string_string_segmentio_unsorted(b *testing.B) {

	b.ReportAllocs()

	m := mapPayload()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		buf := jingo.NewBufferFromPool()
		enc := segmentio.NewEncoder(buf)
		enc.SetSortMapKeys(false)
		if err := enc.Encode(&m); err != nil {
			b.Fatal("Encode:", err)
		}
		b.SetBytes(int64(len(buf.Bytes)))
		buf.ReturnToPool()
	}
}

func BenchmarkMarshal_map_string_string(b *testing.B) {

	m := mapPayload()

	benchMarshal(b, m)
}

func benchMarshal(b *testing.B, x interface{}) {

	type marshalFunc func(interface{}) ([]byte, error)

	for _, bb := range []struct {
		name string
		fn   marshalFunc
	}{
		{"std", json.Marshal},
		{"jsoniterator_sorted", jsoniterator.ConfigCompatibleWithStandardLibrary.Marshal},
		{"jsoniterator_unsorted", jsoniterator.ConfigFastest.Marshal},
		{"segmentio_sorted", segmentio.Marshal},
		{"jettison_sorted", jettison.Marshal},
		{"jettison_unsorted", func(s interface{}) ([]byte, error) { return jettison.MarshalOpts(s, jettison.UnsortedMap()) }},
	} {
		bb := bb
		b.Run(bb.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				bs, err := bb.fn(x)
				if err != nil {
					b.Error(err)
				}
				b.SetBytes(int64(len(bs)))
			}
		})
	}
}
