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
    var stack = ""
    var i = 0
    var ipt = ""
    var str0 = 0
    var str1 = 0
    var fn = ""
    var mul = ""
    var fi = ""
    for i != len(s) {
        if string(s[i]) == "!" {
            i += 1
            fmt.Println(string(s[i]))
        } else if string(s[i]) == ">" {
            i += 1
            stack += string(s[i])
        } else if string(s[i]) == "^" {
            fmt.Println(stack)
        } else if string(s[i]) == "<" {
            fmt.Println("$~ input: ")
            fmt.Scanln(&ipt)
            stack += string(ipt)
        } else if string(s[i]) == "+" {
            if isInt(string(stack[0])) && isInt(string(stack[1])) {
                str0, _ = strconv.Atoi(string(stack[0]))
                str1, _ = strconv.Atoi(string(stack[1]))
                stack += string(str0 + str1)
            } else {
                stack += string(str0) + string(str1)
            }
        } else if string(s[i]) == "-" {
            if isInt(string(stack[0])) && isInt(string(stack[1])) {
                str0, _ = strconv.Atoi(string(stack[0]))
                str1, _ = strconv.Atoi(string(stack[1]))
                stack += string(str0 - str1)
            } else {
                stack += "."
            }
        } else if string(s[i]) == "/" {
            if isInt(string(stack[0])) && isInt(string(stack[1])) {
                str0, _ = strconv.Atoi(string(stack[0]))
                str1, _ = strconv.Atoi(string(stack[1]))
                stack += string(str0 / str1)
            } else {
                stack += "."
            }
        } else if string(s[i]) == "*" {
            if isInt(string(stack[0])) && isInt(string(stack[1])) {
                str0, _ = strconv.Atoi(string(stack[0]))
                str1, _ = strconv.Atoi(string(stack[1]))
                stack += string(str0 * str1)
            } else {
                stack += "."
            }
        } else if string(s[i]) == "#" {
            stack = ""
        } else if string(s[i]) == "~" {
            innerexec(fn)
        } else if string(s[i]) == "[" {
            for string(s[i]) != "]" {
                i += 1
                mul += string(s[i])
            }
            mul = strings.ReplaceAll(mul, "[]", "")
            stack += mul
        } else if string(s[i]) == "{" {
            for string(s[i]) != "}" {
                i += 1
                fn += string(s[i])
            }
        }
    }
}
func main() {
    fmt.Print("\033[H\033[2J")
    var see = ""
    fmt.Println("$~ Ark 1.8.9\n$~ Spargle\n$~ c) 2021\n$~all systems go\n")
    for {
        fmt.Print("ark 1.8.9 > ")
        fmt.Scanln(&see)
        if see == "clear" {
            fmt.Print("\033[H\033[2J")
        }
        if see != "clear" {
            innerexec(see)
        }
    }
}
