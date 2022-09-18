package queryerr

type PatternMatchError struct {
	ReceivedType  string
	ReceivedValue string
	ExpectedType  string
}
