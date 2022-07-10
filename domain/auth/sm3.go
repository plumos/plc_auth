package auth

type Sm3 struct {
}

func (r *Sm3) NewAuth() HashAuth {

	return &Sm3{}
}

func (r *Sm3) Hash() {

}
