package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type (
	SimpleExcel struct {
		NumberCell int
		InputCell  map[string]*Cell
		Keys       []string
		Circle     bool
		ListCir    []string
	}
	Cell struct {
		CellName  string
		RawInput  string
		DependsOn []string
		Tokens    []string
		IntValue  int64
		Done      bool
	}
)

// return of this value; if circel, then bool = true
func (c *Cell) cal(cells map[string]*Cell, depends []string) (int64, []string) {
	if !c.Done {
		if len(c.DependsOn) > 0 {
			mapVal := map[string]int64{}
			// get all value of depending cell
			for _, v := range c.DependsOn {
				for i, d := range depends {
					if d == v {
						return 0, depends[i:]
					}
				}
				dCell, ok := cells[v]
				if !ok {
					cells[v] = &Cell{IntValue: 0, Done: true}
					mapVal[v] = 0
					continue
				}
				depends = append(depends, dCell.CellName)
				res, circel := dCell.cal(cells, depends)
				if len(circel) > 0 {
					return 0, circel
				}
				mapVal[v] = res
				depends = depends[:len(depends)]
			}
			// cal current cell
			for i, v := range c.Tokens {
				if nv, ok := mapVal[v]; ok {
					c.Tokens[i] = strconv.Itoa(int(nv))
				}
			}
			c.IntValue, _ = ReversePolishNotation(c.Tokens)
		} else {
			var err error
			c.IntValue, err = strconv.ParseInt(c.RawInput, 10, 64)
			if err != nil {
				c.IntValue, _ = ReversePolishNotation(c.Tokens)
			}
		}
		c.Done = true
	}
	return c.IntValue, []string{}
}

// token for each cell
func (c *Cell) token() {
	c.Tokens = strings.Fields(c.RawInput)
	if len(c.Tokens) > 1 {
		for _, key := range c.Tokens {
			if !checkOperator(key) && !checkInt(key) {
				c.DependsOn = append(c.DependsOn, key)
			}
		}
	}
}

// token all raw input of each cell
func (se *SimpleExcel) Token() {
	for _, cell := range se.InputCell {
		cell.token()
	}
}
func (se *SimpleExcel) cal() {
	for _, cell := range se.InputCell {
		if !cell.Done {
			cir := []string{cell.CellName}
			_, ccir := cell.cal(se.InputCell, cir)
			if len(ccir) > 0 {
				se.Circle = true
				se.ListCir = ccir
				return
			}
		}
	}
}

func (se *SimpleExcel) printOut() {
	if se.Circle {
		sort.Strings(se.ListCir)
		res := strings.Join(se.ListCir, ", ")
		fmt.Printf("Circular dependency detected: %s", res)
	} else {
		sort.Strings(se.Keys)
		for _, keys := range se.Keys {
			fmt.Println(keys)
			fmt.Println(se.InputCell[keys].IntValue)
		}
	}

}
func (se *SimpleExcel) justPrint() {
	sort.Strings(se.Keys)
	for _, keys := range se.Keys {
		fmt.Println(keys)
		fmt.Println(se.InputCell[keys].RawInput)
	}
}
func pop2(stack []int64) ([]int64, int64, int64) {
	var ab []int64
	stack, ab = stack[:len(stack)-2], stack[len(stack)-2:]
	return stack, ab[0], ab[1]
}
func checkOperator(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}

func checkInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	} else {
		return true
	}
}

// cal reverse polist notation
func ReversePolishNotation(tokens []string) (int64, error) {
	var stack []int64
	var value int64
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			var a, b int64
			if len(stack) < 2 {
				return 0, fmt.Errorf("not enough elmenents on the stack to calculate %s", token)
			}
			stack, a, b = pop2(stack)
			switch token {
			case "+":
				value = a + b
			case "-":
				value = a - b
			case "*":
				value = a * b
			case "/":
				value = int64(float64(a) / float64(b))
			}
		default:
			var err error
			value, err = strconv.ParseInt(token, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid token %s", token)
			}
		}
		stack = append(stack, value)
	}
	if len(stack) != 1 {
		return 0, fmt.Errorf("incomplete expression %s", stack)
	}
	return stack[0], nil
}

// read input and return error if input is not valid
func readInput() (*SimpleExcel, error) {
	stdin := os.Stdin
	scanner := bufio.NewScanner(stdin)
	se := &SimpleExcel{InputCell: map[string]*Cell{}, Keys: []string{}}
	count := 0
	for scanner.Scan() {
		in := scanner.Text()
		if count == 0 { // read number of line
			numberCells, err := strconv.Atoi(in)
			if err != nil {
				return se, err
			}
			se.NumberCell = numberCells

		} else { // parse input
			if count%2 == 1 {
				se.Keys = append(se.Keys, in)
			} else {
				keys := se.Keys[len(se.Keys)-1]
				se.InputCell[keys] = &Cell{RawInput: in, CellName: keys}
			}
		}
		count++
		if count == se.NumberCell*2+1 {
			break
		}
	}
	stdin.Close()
	return se, nil
}

func main() {
	sc, _ := readInput()
	sc.Token()
	sc.cal()
	sc.printOut()
}
