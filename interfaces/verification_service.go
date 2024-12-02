package interfaces

type VerificationService interface {
	GenerateChallenge(expected []byte) ([]byte, error)
	VerifyChallange(message []byte, signature []byte, expectedAddress string) (*string, error)
	VerifyCertificateChallange(certificate []byte, signedChallenge []byte) (*string, error)
	Init()
}
