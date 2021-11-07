package helper

func IsNumber(keyword string) bool {
	if keyword >= "0" && keyword <= "9" {
		return true
	}
	return false
}

func IsIdentifier(keyword string) bool {
	if IsKeyword(keyword) {
		return false
	}
	if keyword >= "a" && keyword <= "z" {
		return true
	}
	return false
}

func IsOperand(keyword string) bool {
	if keyword == "+" || keyword == "-" || keyword == "=" {
		return true
	}
	return false
}

func IsKeyword(keyword string) bool {
	if keyword == "let" || keyword == "func" {
		return true
	}
	return false
}

func IsEndExpression(keyword string) bool {
	if keyword == ";" {
		return true
	}
	return false
}
