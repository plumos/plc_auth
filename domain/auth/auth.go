package auth

type UnAsyAuth interface {
	Sign()
	Verify()
	Encode()
	Decode()
}

type AsyAuth interface {
	Encode()
	Decode()
}

type HashAuth interface {
	Hash()
}
