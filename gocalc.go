package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container" 
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type HistoryEntry struct {
	Expression string
	Result     string
}

func main() {
	a := app.New()
	w := a.NewWindow("A calc that doesn't phone home.")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter expression or use buttons")
	input.TextStyle = fyne.TextStyle{Monospace: true}

	var history []HistoryEntry
	historyList := widget.NewList(
		func() int { return len(history) },
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(fmt.Sprintf("%s = %s", history[id].Expression, history[id].Result))
		},
	)

	historyList.OnSelected = func(id widget.ListItemID) {
		input.SetText(history[id].Expression)
		historyList.UnselectAll()
	}

	clearHistoryBtn := widget.NewButton("Clear History", func() {
		history = []HistoryEntry{}
		historyList.Refresh()
	})

	var evaluate func()

	input.OnSubmitted = func(s string) {
		evaluate()
	}

	appendInput := func(s string) {
		current := input.Text
		if current == "0" && s != "." {
			input.SetText(s)
		} else {
			input.SetText(current + s)
		}
	}

	evaluate = func() {
		expression := strings.TrimSpace(input.Text)
		result, err := evaluateExpression(expression)
		if err != nil {
			input.SetText("Error: " + err.Error())
		} else {
			resultStr := fmt.Sprintf("%.8g", result)
			input.SetText(resultStr)
			history = append([]HistoryEntry{{Expression: expression, Result: resultStr}}, history...)
			historyList.Refresh()
		}
	}

	buttons := []string{
		"7", "8", "9", "/",
		"4", "5", "6", "*",
		"1", "2", "3", "-",
		"0", ".", "^", "+",
		"(", ")", "C", "=",
	}

	buttonGrid := container.New(layout.NewGridLayout(4))
	for _, label := range buttons {
		btn := widget.NewButton(label, func(l string) func() {
			return func() {
				switch l {
				case "C":
					input.SetText("")
				case "=":
					evaluate()
				default:
					appendInput(l)
				}
			}
		}(label))
		buttonGrid.Add(btn)
	}

	topContent := container.NewVBox(
		input,
		buttonGrid,
		clearHistoryBtn,
	)

	content := container.NewBorder(
		topContent,
		nil,
		nil,
		nil,
		container.NewScroll(historyList),
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(300, 500))
	w.SetFixedSize(true)
	w.ShowAndRun()
}

func evaluateExpression(expression string) (float64, error) {
	tokens := tokenize(expression)
	postfix, err := infixToPostfix(tokens)
	if err != nil {
		return 0, err
	}
	return evaluatePostfix(postfix)
}

func tokenize(expression string) []string {
	var tokens []string
	var current string
	for _, char := range expression {
		switch {
		case char >= '0' && char <= '9' || char == '.':
			current += string(char)
		case char == '+' || char == '-' || char == '*' || char == '/' || char == '^' || char == '(' || char == ')':
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			tokens = append(tokens, string(char))
		case char == ' ':
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
		}
	}
	if current != "" {
		tokens = append(tokens, current)
	}
	return tokens
}

func infixToPostfix(tokens []string) ([]string, error) {
	precedence := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "^": 3}
	var postfix []string
	var stack []string

	for _, token := range tokens {
		switch {
		case token == "(":
			stack = append(stack, token)
		case token == ")":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, fmt.Errorf("mismatched parentheses")
			}
			stack = stack[:len(stack)-1] // Pop the "("
		case precedence[token] > 0:
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		default:
			postfix = append(postfix, token)
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix, nil
}

func evaluatePostfix(tokens []string) (float64, error) {
	var stack []float64

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/", "^":
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression")
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			
			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				result = a / b
			case "^":
				result = math.Pow(a, b)
			}
			stack = append(stack, result)
		default:
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid token: %s", token)
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}
	return stack[0], nil
}
