package build_in

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func fmtInt() string {
	return fmt.Sprint(100)
}

func strconvInt() string {
	return strconv.FormatInt(100, 10)
}

func fmtJoin() string {
	return fmt.Sprintf("%s:%d", "after_rank", 100)
}

func strconvJoin() string {
	return strings.Join([]string{"after_rank", strconv.FormatInt(100, 10)}, ":")
}

func strconvAdd() string {
	return "after_rank:" + strconv.FormatInt(100, 10)
}

func TestFmtJoin(t *testing.T) {
	fmtJoin()
}

func TestStrconvAdd(t *testing.T) {
	strconvAdd()
}

func BenchmarkFmtJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmtJoin()
	}
}

func BenchmarkStrconvAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconvAdd()
	}
}

func BenchmarkStrconvJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconvJoin()
	}
}

func BenchmarkFmtInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmtInt()
	}
}

func BenchmarkStrconvInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconvInt()
	}
}
