package models

type Pupil struct {
}

func (_ *Pupil) HasACar(name string) bool {
	return true
}
