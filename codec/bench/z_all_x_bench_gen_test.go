// Copyright (c) 2012-2018 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

//go:build alltests && x && generated

package codec

// see notes in z_all_bench_test.go

import "testing"

func benchmarkCodecXGenGroup(t *testing.B) {
	benchmarkDivider()
	t.Run("Benchmark__Msgpack____Encode", Benchmark__Msgpack____Encode)
	t.Run("Benchmark__Json_______Encode", Benchmark__Json_______Encode)
	t.Run("Benchmark__Std_Json___Encode", Benchmark__Std_Json___Encode)
	t.Run("Benchmark__Gob________Encode", Benchmark__Gob________Encode)
	t.Run("Benchmark__JsonIter___Encode", Benchmark__JsonIter___Encode)
	t.Run("Benchmark__Bson_______Encode", Benchmark__Bson_______Encode)
	t.Run("Benchmark__VMsgpack___Encode", Benchmark__VMsgpack___Encode)
	t.Run("Benchmark__Msgp_______Encode", Benchmark__Msgp_______Encode)
	t.Run("Benchmark__Easyjson___Encode", Benchmark__Easyjson___Encode)
	t.Run("Benchmark__Ffjson_____Encode", Benchmark__Ffjson_____Encode)
	// t.Run("Benchmark__Gcbor______Encode", Benchmark__Gcbor______Encode)
	// t.Run("Benchmark__Xdr________Encode", Benchmark__Xdr________Encode)
	t.Run("Benchmark__Sereal_____Encode", Benchmark__Sereal_____Encode)

	benchmarkDivider()
	t.Run("Benchmark__Msgpack____Decode", Benchmark__Msgpack____Decode)
	t.Run("Benchmark__Json_______Decode", Benchmark__Json_______Decode)
	t.Run("Benchmark__Std_Json___Decode", Benchmark__Std_Json___Decode)
	t.Run("Benchmark__Gob________Decode", Benchmark__Gob________Decode)
	t.Run("Benchmark__JsonIter___Decode", Benchmark__JsonIter___Decode)
	t.Run("Benchmark__Bson_______Decode", Benchmark__Bson_______Decode)
	// t.Run("Benchmark__VMsgpack___Decode", Benchmark__VMsgpack___Decode)
	//t.Run("Benchmark__Msgp_______Decode", Benchmark__Msgp_______Decode) // fails
	t.Run("Benchmark__Easyjson___Decode", Benchmark__Easyjson___Decode)
	t.Run("Benchmark__Ffjson_____Decode", Benchmark__Ffjson_____Decode)
	// t.Run("Benchmark__Gcbor______Decode", Benchmark__Gcbor______Decode)
	// t.Run("Benchmark__Xdr________Decode", Benchmark__Xdr________Decode)
	// t.Run("Benchmark__Sereal_____Decode", Benchmark__Sereal_____Decode)
}

func BenchmarkCodecXGenSuite(t *testing.B) { benchmarkSuite(t, benchmarkCodecXGenGroup) }
