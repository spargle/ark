// ark_interpreter c) 2022
package main
import(
    "unicode"
    "fmt"
    "strconv"
    "strings"
    "math/rand"
    "log"
    "time"
    "os"
)
var runtime_tag = ""
var Token = ""
func isInt(sr string) bool {
    for _, c := range sr {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}
func new() string{
    Token = Tag(); time.Sleep(1 * time.Second); log.Println("runtime key\n" + Token); tk1 := Tag(); time.Sleep(1 * time.Second); tk2 := Tag(); time.Sleep(1 * time.Second); tk3 := Tag(); time.Sleep(1 * time.Second); tk4 := Tag(); jko := [5]string{tk1, tk2, Token, tk3, tk4}; res := ""; e := 0
    for e != 5 {n := rand.Intn(5); res += jko[n]; res += "\n"; rm(jko, n); e += 1}
    return res
}
func Tag() string {
    now := time.Now(); unixNano := now.UnixNano(); umillisec := unixNano; r1:=[]rune("abc1def2ghi3klm4nop5qrs6tuv8wxy9z0"); r:=[]rune(""); stt:=""; r2:=""; i:=0
    for i != 36 {now = time.Now(); unixNano = now.UnixNano(); umillisec = unixNano; rand.Seed(umillisec); r1 := r1[rand.Intn(len(r1))]; r2 += string(r1); i += 1}
    r = []rune(r2); stt = ""; i = 0
    for i != 41 {now = time.Now(); unixNano = now.UnixNano(); umillisec = unixNano; rand.Seed(umillisec); m:=r[rand.Intn(len(r))]; stt += string(m); i += 1}
    return string(stt)
}
func rm(slice [5]string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}
func innerexec(s string) {
    var stack = ""
    var i = 0
    var it = ""
    var s0 = 0
    var s1 = 0
    var fn = ""
    var m = ""
    for i != len(s) {
        if string(s[i]) == "!" {
            i += 1
            log.Println(string(s[i]))
        } else if string(s[i]) == ">" {
            i += 1
            stack += string(s[i])
        } else if string(s[i]) == "^" {
            log.Println(stack)
        } else if string(s[i]) == "<" {
            log.Println("$~ input: ")
            fmt.Scanln(&it)
            stack += string(it)
        } else if string(s[i]) == "+" {
            if isInt(string(stack[0])) && isInt(string(stack[1])) {
                s0, _ = strconv.Atoi(string(stack[0]))
                s1, _ = strconv.Atoi(string(stack[1]))
                stack += string(s0 + s1)
            } else {
                stack += string(s0) + string(s1)
            }
        } else if string(s[i]) == "-" {
            if isInt(string(stack[0])) && isInt(string(stack[1])) {
                s0, _ = strconv.Atoi(string(stack[0]))
                s1, _ = strconv.Atoi(string(stack[1]))
                stack += string(s0 - s1)
            } else {
                stack += "."
            }
        } else if string(s[i]) == "/" {
            if isInt(string(stack[0])) && isInt(string(stack[1])) {
                s0, _ = strconv.Atoi(string(stack[0]))
                s1, _ = strconv.Atoi(string(stack[1]))
                stack += string(s0 / s1)
            } else {
                stack += "."
            }
        } else if string(s[i]) == "*" {
            if isInt(string(stack[0])) && isInt(string(stack[1])) {
                s0, _ = strconv.Atoi(string(stack[0]))
                s1, _ = strconv.Atoi(string(stack[1]))
                stack += string(s0 * s1)
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
                m += string(s[i])
            }
            m = strings.ReplaceAll(m, "[]", "")
            stack += m
        } else if string(s[i]) == "{" {
            for string(s[i]) != "}" {
                i += 1
                fn += string(s[i])
            }
        } else if string(s[i]) == "@" {
            if !strings.Contains(runtime_tag, Token) {
                os.Exit(0)
            }
        } else if string(s[i]) == "%" {
            go innerexec(fn)
        } else if string(s[i]) == "" {

        }
        i += 1
    }
}
func main() {
    fmt.Print("\033[H\033[2J")
    var see = ""
    log.Println("\n$~ Ark 2.0.0\n$~ Spargle\n$~ c) 2022\n$~ all systems go\n$~ generating runtime tag...")
    runtime_tag = new()
    log.Println("$~ done.")
    for {
        log.Print("ark 2.0.0 > ")
        fmt.Scanln(&see)
        if see == "clear" {
            log.Print("\033[H\033[2J")
        }
        if see != "clear" {
                if strings.Contains(runtime_tag, Token) {
                    fmt.Print("")
                    innerexec(see)
                } else if !strings.Contains(runtime_tag, Token) {
                    log.Print("runtime stabilization failed. Please try again.")
                }
        }
    }
}
