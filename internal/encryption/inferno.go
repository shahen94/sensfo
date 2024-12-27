package encryption

// ─────────────────────────────────────────────────────────────────────────────

type InfernoEncryption struct{}

type InfernoEncryptionResult struct{}

// ─────────────────────────────────────────────────────────────────────────────

func (i *InfernoEncryptionResult) Bias() string {
	return ""
}

func (i *InfernoEncryptionResult) Content() string {
	return ""
}

func (i *InfernoEncryption) ComputeBias(data string) string {
	return ""
}

func (i *InfernoEncryption) Encrypt(data string) (Result, error) {
	return &InfernoEncryptionResult{}, nil
}

func (i *InfernoEncryption) Decrypt(data string, bias string) (string, error) {
	return "", nil
}

// ─────────────────────────────────────────────────────────────────────────────

func NewInfernoEncryption() Encryption {
	return &InfernoEncryption{}
}
