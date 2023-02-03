package main

import "testing"

func BenchmarkAllDirectoriesMain(b *testing.B) {
	b.Log("Profiling AllFilesMain")

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ExecuteApp("enron_mail_20110402.tgz", false)
	}

	b.StopTimer()
}

func BenchmarkInboxFilesMain(b *testing.B) {

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ExecuteApp("enron_mail_20110402.tgz", true)
	}

	b.StopTimer()

}
