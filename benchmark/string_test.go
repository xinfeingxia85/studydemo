package benchmark

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var strlen int = 1000
var testRegexp string = `^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]+$`

func BenchmarkConcatString(b *testing.B) {
	var str string
	i := 0
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		str += "x"
		i++
		if i >= strlen {
			i = 0
			str = ""
		}
	}
}

func BenchmarkConcatBuffer(b *testing.B) {
	var buffer bytes.Buffer

	i := 0
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")

		i++
		if i >= strlen {
			i = 0
			buffer = bytes.Buffer{}
		}
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	var builder strings.Builder

	i := 0
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		builder.WriteString("x")
		i++
		if i >= strlen {
			i = 0
			builder = strings.Builder{}
		}
	}
}

func BenchmarkParseInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseInt("7182818284", 10, 64)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkParseFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseFloat("3.1415926535", 64)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkParseBool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseBool("true")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMatchString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := regexp.MatchString(testRegexp, "jsmitch@expamle.com")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMatchStringCompiled(b *testing.B) {
	r, err := regexp.Compile(testRegexp)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.MatchString("jsmitch@expamle.com")
	}
}