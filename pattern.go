package versions

import(
	"strings"
	"os"
	)

type Pattern struct {
	Value string
	Operator *Operator
	Version *Version
}

type Operator struct {
	Type int
}

const(
	EQUAL = iota
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL
)

var LiteralToOperator map[string]int
var OperatorMapInitialized bool


func initializeOperatorMap() {
	if OperatorMapInitialized {
		return
	}

	LiteralToOperator = map[string] int{
		">": GREATER,
		"=": EQUAL,
	}
	
	OperatorMapInitialized = true
}

func NewPattern(value string) *Pattern{
	value = strings.TrimLeft(value, " \r\n")
	value = strings.TrimRight(value, " \r\n")

	p := &Pattern{
	Value: value,
	}

	p.parse()

	return p
}


func (p *Pattern) parse() (err os.Error) {
	tokens := strings.Split(p.Value, " ")

	if len(tokens) > 1 {
		
	} else {
		version, err := NewVersion(p.Value)

		if err != nil {
			return err
		}

		p.Version = version			

	}
	
	return nil
}