package auth

type Des struct {
}

func (r *Des) NewAsyAuth() AsyAuth {

	return &Des{}
}

func (r *Des) Encode() {

}
func (r *Des) Decode() {

}
