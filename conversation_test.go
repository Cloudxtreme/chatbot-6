package main

import "testing"

func RandomConversation() *Conversation {
	var conversation *Conversation = new(Conversation)
	conversation.Question = "hi"
	conversation.Response = "hi, what's up ?"
	return conversation
}

func TestSaveConversation(t *testing.T) {
	conversation := RandomConversation()
	var result bool = SaveConversation(conversation)
	if !result {
		t.Error("can't save conversation")
	}
}

func benchmarkSaveConversation(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		conversation := RandomConversation()
		SaveConversation(conversation)
	}
}

func BenchmarkSaveConversation1(b *testing.B)  { benchmarkSaveConversation(1, b) }
func BenchmarkSaveConversation2(b *testing.B)  { benchmarkSaveConversation(2, b) }
func BenchmarkSaveConversation3(b *testing.B)  { benchmarkSaveConversation(3, b) }
func BenchmarkSaveConversation10(b *testing.B) { benchmarkSaveConversation(10, b) }
func BenchmarkSaveConversation20(b *testing.B) { benchmarkSaveConversation(20, b) }
func BenchmarkSaveConversation40(b *testing.B) { benchmarkSaveConversation(40, b) }
