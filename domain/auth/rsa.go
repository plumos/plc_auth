package auth

type Rsa struct {
	RsaType int //  1024 / 2048
}

func (r *Rsa) NewUnAuth(rsaType int) UnAsyAuth {
	if rsaType == 0 {
		return &Rsa{1024}
	}
	return &Rsa{2048}
}

func (r *Rsa) Sign() {

}
func (r *Rsa) Verify() {

}
func (r *Rsa) Encode() {

}
func (r *Rsa) Decode() {

}
