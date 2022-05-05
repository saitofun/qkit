package enum

type IntStringerEnum interface {
	Typename() string
	Int() string
	String() string
	Label() string
	ConstValues() []IntStringerEnum
}
