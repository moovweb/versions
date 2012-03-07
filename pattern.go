package versions

import(
	"strings"
	"os"
	)

type Pattern struct {
	Operator *Operator
	Version *Version
}

type Operator struct {
	Type int
}

func NewOperator(tipe int) *Operator {
	return &Operator{
	Type: tipe,
	}
}

const(
	LESS = iota
	LESS_EQUAL
	EQUAL
	PESSIMISTIC
	GREATER_EQUAL
	GREATER
)

var LiteralToOperator map[string]int
var OperatorMapInitialized bool

func initializeOperatorMap() {
	if OperatorMapInitialized {
		return
	}

	LiteralToOperator = map[string] int{
		"<": LESS,
		"<=": LESS_EQUAL,
		"~>": PESSIMISTIC,
		"=": EQUAL,
		">=": GREATER_EQUAL,
		">": GREATER,
	}
	
	OperatorMapInitialized = true
}

func NewPattern(value string) (p *Pattern, err os.Error) {
	value = strings.TrimLeft(value, " \r\n")
	value = strings.TrimRight(value, " \r\n")

	p, err = parse(value)

	return
}


func parse(value string) (p *Pattern, err os.Error) {
	tokens := strings.Split(value, " ")
	rawValue := value

	p = &Pattern{}

	if len(tokens) > 1 {
		p.Operator = NewOperator( LiteralToOperator[tokens[0]] )
		rawValue = tokens[1]
	}

	version, err := NewVersion(rawValue)

	if err != nil {
		return nil, err
	}

	p.Version = version			
	
	return
}

func (p *Pattern) Match(version *Version) bool {
	var result bool

	switch(p.Operator.Type) {
	case LESS:
		result = p.Less(version)
	case LESS_EQUAL:
		result = p.LessEqual(version)
	case EQUAL:
		result = p.Equal(version)
	case PESSIMISTIC:
		result = p.Pessimistic(version)
	case GREATER_EQUAL:
		result = p.GreaterEqual(version)
	case GREATER:
		result = p.Greater(version)
	}

	return result
}
