package models

type DataRequest struct {
	Certificate     []byte
	SignedChallenge []byte
}
