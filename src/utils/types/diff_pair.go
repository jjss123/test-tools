package types

type DiffPair struct {
	Expect map[string]interface{} `json:"Expect"`
	Actual map[string]interface{} `json:"Actual"`
}

func NewDiffPair() DiffPair {
	return DiffPair{make(map[string]interface{}), make(map[string]interface{})}
}

func (p *DiffPair) IsExpectExist(str string) bool {
	if p.Expect == nil {
		p.Expect = make(map[string]interface{})
	}
	if _, ok := p.Expect[str]; !ok {
		return false
	} else {
		return true
	}
}

func (p *DiffPair) AddExpect(str string, val interface{}) {
	if p.Expect == nil {
		p.Expect = make(map[string]interface{})
	}
	p.Expect[str] = val
}

func (p *DiffPair) DelExpect(str string) {
	if p.Expect == nil {
		p.Expect = make(map[string]interface{})
	}
	delete(p.Expect, str)
}

func (p *DiffPair) AddActual(str string, val interface{}) {
	if p.Actual == nil {
		p.Actual = make(map[string]interface{})
	}
	p.Actual[str] = val
}

func (p *DiffPair) DelActual(str string) {
	if p.Actual == nil {
		p.Actual = make(map[string]interface{})
	}
	delete(p.Actual, str)
}

func Diff(first, second map[string]interface{}) (add /*first - second*/, del /*second - first*/ map[string]interface{}) {
	a := make(map[string]interface{})
	d := make(map[string]interface{})
	for key, value := range first {
		if _, ok := second[key]; !ok {
			a[key] = value
		}
	}

	for key, value := range second {
		if _, ok := first[key]; !ok {
			d[key] = value
		}
	}
	return a, d
}

func HasDiff(first, second map[string]interface{}) bool {
	add, del := Diff(first, second)
	return len(add) != 0 || len(del) != 0
}

func (p *DiffPair) Diff() (add map[string]interface{}, del map[string]interface{}) {
	if p == nil {
		return add, del
	}
	return Diff(p.Expect, p.Actual)
}

func (p *DiffPair) HasDiff() bool {
	if p == nil {
		return false
	}
	return HasDiff(p.Expect, p.Actual)
}
