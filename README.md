# golang cmd
go version
go env

go run .()

## Testing
-go test (No file name needed (file name will throw error and its not a convention), just go to directory and run the command.)
-go test -bench=.   // Package test korar jonne
-this function will be in testfile
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}