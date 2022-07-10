package auth

type Sha256 struct {
}

func (r *Sha256) NewAuth() HashAuth {

	return &Sha256{}
}

func (r *Sha256) Hash() {

}
