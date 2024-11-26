package interfaces

type VerificationService interface {
	GenerateChallenge(certificate []byte) ([]byte, error)
	VerifyChallange(certificate []byte, signedChallenge []byte) (*string, error)
	Init()
}
