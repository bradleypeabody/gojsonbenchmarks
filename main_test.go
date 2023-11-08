package main

import (
	"encoding/json"
	"testing"

	"github.com/bradleypeabody/gojsonbenchmarks/codegen"
	gojson "github.com/goccy/go-json"
)

func BenchmarkEncodingJSONMarshal(b *testing.B) {

	things := RandomThings(1, b.N)

	b.ResetTimer()

	for i := 0; i < len(things); i++ {
		_, err := json.Marshal(things[i])
		if err != nil {
			b.Fatal(err)
		}
	}

}

func BenchmarkEncodingJSONUnmarshal(b *testing.B) {

	things := RandomThings(1, b.N)
	thingBytes := make([][]byte, len(things))
	for i := 0; i < len(thingBytes); i++ {
		by, err := json.Marshal(things[i])
		if err != nil {
			b.Fatal(err)
		}
		thingBytes[i] = by
		things[i] = Thing{}
	}

	b.ResetTimer()

	for i := 0; i < len(thingBytes); i++ {
		err := json.Unmarshal(thingBytes[i], &things[i])
		if err != nil {
			b.Fatal(err)
		}
	}

}

func BenchmarkGoccyJSONMarshal(b *testing.B) {

	things := RandomThings(1, b.N)

	b.ResetTimer()

	for i := 0; i < len(things); i++ {
		_, err := gojson.Marshal(things[i])
		if err != nil {
			b.Fatal(err)
		}
	}

}

func BenchmarkGoccyJSONUnmarshal(b *testing.B) {

	things := RandomThings(1, b.N)
	thingBytes := make([][]byte, len(things))
	for i := 0; i < len(thingBytes); i++ {
		by, err := json.Marshal(things[i])
		if err != nil {
			b.Fatal(err)
		}
		thingBytes[i] = by
		things[i] = Thing{}
	}

	b.ResetTimer()

	for i := 0; i < len(thingBytes); i++ {
		err := gojson.Unmarshal(thingBytes[i], &things[i])
		if err != nil {
			b.Fatal(err)
		}
	}

}

func BenchmarkFFJSONMarshal(b *testing.B) {

	things := codegen.RandomThing2s(1, b.N)

	b.ResetTimer()

	for i := 0; i < len(things); i++ {
		_, err := json.Marshal(things[i])
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFFJSONUnmarshal(b *testing.B) {

	things := codegen.RandomThing2s(1, b.N)
	thingBytes := make([][]byte, len(things))
	for i := 0; i < len(thingBytes); i++ {
		by, err := json.Marshal(things[i])
		if err != nil {
			b.Fatal(err)
		}
		thingBytes[i] = by
		things[i] = codegen.Thing2{}
	}

	b.ResetTimer()

	for i := 0; i < len(thingBytes); i++ {
		err := json.Unmarshal(thingBytes[i], &things[i])
		if err != nil {
			b.Fatal(err)
		}
	}

}
