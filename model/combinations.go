package model

type Combinations struct {
	XMLName     struct{}     `xml:"combinations"`
	Combination *Combination `xml:"combination"`
}

type Combination struct {
	XMLName struct{} `xml:"combination"`
	Id      int      `xml:"id"`
}
