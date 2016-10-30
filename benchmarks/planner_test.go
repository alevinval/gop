package benchmarks

import (
	"testing"

	"github.com/alevinval/gop"
)

func BenchmarkBuildPlan4steps(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		initial := gop.NewStack()
		initial.Push(Flubber(1))
		final := gop.NewStack()
		final.Push(Flubber(12))
		gop.BuildPlan(initial, final)
	}
}

func BenchmarkBuildPlan8steps(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		initial := gop.NewStack()
		initial.Push(Flubber(1))
		final := gop.NewStack()
		final.Push(Flubber(31))
		gop.BuildPlan(initial, final)
	}
}

func BenchmarkBuildPlan16steps(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		initial := gop.NewStack()
		initial.Push(Flubber(1))
		final := gop.NewStack()
		final.Push(Flubber(1024 * 21))
		gop.BuildPlan(initial, final)
	}
}

func BenchmarkBuildPlan32steps(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		initial := gop.NewStack()
		initial.Push(Flubber(1))
		final := gop.NewStack()
		final.Push(Flubber(1235567))
		gop.BuildPlan(initial, final)
	}
}
