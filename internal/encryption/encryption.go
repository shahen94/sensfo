package encryption

type Encryption interface {
	Encrypt(data string) (Result, error)
	Decrypt(data string, bias string) (string, error)
	ComputeBias(data string) string
}

type Result interface {
	Bias() string
	Content() string
}
