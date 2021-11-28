// ark_interpreter
package main
import "fmt"
import "unicode"
import "strconv"
import "strings"
func isInt(sr string) bool {
    for _, c := range sr {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}
func innerexec(s string) {
    var r = 0
    var stack []string
    var i = 0
    var ipt = ""
    var str0 = 0
    var str1 = 0
    var fn = ""
    var mul = ""
    for i != len(s) {
        if string(s[i]) == "!" {
            i += 1
            fmt.Println(string(s[i]))
        } else if string(s[i]) == ">" {
            i += 1
            stack[r] = string(s[i])
            r += 1
        } else if string(s[i]) == "^" {
            fmt.Println(stack)
        } else if string(s[i]) == "<" {
            fmt.Println("$~ input: ")
            fmt.Scanln(&ipt)
            stack[r] = string(ipt)
            r += 1
        } else if string(s[i]) == "+" {
            if isInt(stack[0]) && isInt(stack[1]) {
                str0, _ = strconv.Atoi(string(stack[0]))
                str1, _ = strconv.Atoi(string(stack[1]))
                stack[r] = string(str0 + str1)
                r += 1
            } else {
                stack[r] = string(str0) + string(str1)
                r += 1
            }
        } else if string(s[i]) == "-" {
            if isInt(stack[0]) && isInt(stack[1]) {
                str0, _ = strconv.Atoi(string(stack[0]))
                str1, _ = strconv.Atoi(string(stack[1]))
                stack[r] = string(str0 - str1)
                r += 1
            } else {
                stack[r] = "nil"
                r += 1
            }
        } else if string(s[i]) == "/" {
            if isInt(stack[0]) && isInt(stack[1]) {
                str0, _ = strconv.Atoi(string(stack[0]))
                str1, _ = strconv.Atoi(string(stack[1]))
                stack[r] = string(str0 / str1)
                r += 1
            } else {
                stack[r] = "nil"
                r += 1
            }
        } else if string(s[i]) == "*" {
            if isInt(stack[0]) && isInt(stack[1]) {
                str0, _ = strconv.Atoi(string(stack[0]))
                str1, _ = strconv.Atoi(string(stack[1]))
                stack[r] = string(str0 * str1)
                r += 1
            } else {
                stack[r] = "nil"
                r += 1
            }
        } else if string(s[i]) == "#" {
            var rm = 0
            for rm != 1000 {
                stack[rm] = ""
                rm += 1
            }
        } else if string(s[i]) == "~" {
            innerexec(fn)
        } else if string(s[i]) == "[" {
            for string(s[i]) != "]" {
                i += 1
                mul += string(s[i])
            }
            mul = strings.ReplaceAll(mul, "[]", "")
            stack[r] = mul
            r += 1
        } else if string(s[i]) == "{" {
            for string(s[i]) != "}" {
                i += 1
                fn += string(s[i])
            }
            fn = strings.ReplaceAll(fn, "{}", "")
        }
        i += 1
    }
}
func main() {
    var see = ""
    fmt.Println("$~ Ark 1.8.0")
    fmt.Scanln(&see)
    innerexec(see)
}
