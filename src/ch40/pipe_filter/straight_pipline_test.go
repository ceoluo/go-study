package pipe_filter

import "testing"

func TestStraightPipline(t *testing.T)  {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sumer := NewSumFilter()
	strpip := NewStraightPipline("p1", spliter, converter, sumer )
	ret, err := strpip.Process("1,2,3")
	if err != nil{
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("The expected is 6, but the actul is %d", ret)
	}
}

func BenchmarkStraightPipline(b *testing.B) {
	b.ResetTimer()
	for i := 0;i<b.N;i++{
		spliter := NewSplitFilter(",")
		converter := NewToIntFilter()
		sumer := NewSumFilter()
		strpip := NewStraightPipline("p1", spliter, converter, sumer )
		ret, err := strpip.Process("1,2,3")
		if err != nil{
			b.Fatal(err)
		}
		if ret != 6 {
			b.Fatalf("The expected is 6, but the actul is %d", ret)
		}
	}
	b.StopTimer()
}
