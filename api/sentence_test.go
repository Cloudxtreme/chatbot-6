package main

import "testing"

func RandomSentence() *Sentence {
	var sentence *Sentence = new(Sentence)
	sentence.Sentence = "test"
	return sentence
}

func TestSaveSentence(t *testing.T) {
	sentence := RandomSentence()
	var result bool = SaveSentence(sentence)
	if !result {
		t.Error("can't save sentence")
	}
}

func benchmarkSaveSentence(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		sentence := RandomSentence()
		SaveSentence(sentence)
	}
}

func BenchmarkSaveSentence1(b *testing.B)  { benchmarkSaveSentence(1, b) }
func BenchmarkSaveSentence2(b *testing.B)  { benchmarkSaveSentence(2, b) }
func BenchmarkSaveSentence3(b *testing.B)  { benchmarkSaveSentence(3, b) }
func BenchmarkSaveSentence10(b *testing.B) { benchmarkSaveSentence(10, b) }
func BenchmarkSaveSentence20(b *testing.B) { benchmarkSaveSentence(20, b) }
func BenchmarkSaveSentence40(b *testing.B) { benchmarkSaveSentence(40, b) }
