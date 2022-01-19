package funcService

type IFuncService interface {
	BinaryToDecimal(inp string) int
	DecimalToBinary(inp string) int
	Polyndrome(inp string) string
}
