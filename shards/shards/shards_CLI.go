package main
import (
    "fmt"
    "strconv"
    "unicode"
    "strings"
    "math/rand"
    "log"
    "time"
    "os"
)
var (
    ipt = ""
    instances = ""
    version = "1.4.0"
    out = ""
    out1 = ""
)
var runtime_tag = ""
var Token = ""
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
func isInt(sr string) bool {
    for _, c := range sr {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}
func innerexec(s string, name string) {
    instances = strings.ReplaceAll(instances, string("| - " + name + " @ " + version + " | failing"), string("| - " + name + " @ " + version + " | good"))
    instances = strings.ReplaceAll(instances, string("| - " + name + " @ " + version + " | bad"), string("| - " + name + " @ " + version + " | good"))
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
                instances = strings.ReplaceAll(instances, string("| - " + name + " @ " + version + " | failing"), string("| - " + name + " @ " + version + " | bad"))
                instances = strings.ReplaceAll(instances, string("| - " + name + " @ " + version + " | good"), string("| - " + name + " @ " + version + " | bad"))
            }
        } else if string(s[i]) == "/" {
            if isInt(string(stack[0])) && isInt(string(stack[1])) {
                s0, _ = strconv.Atoi(string(stack[0]))
                s1, _ = strconv.Atoi(string(stack[1]))
                stack += string(s0 / s1)
            } else {
                stack += "."
                instances = strings.ReplaceAll(instances, string("| - " + name + " @ " + version + " | failing"), string("| - " + name + " @ " + version + " | bad"))
                instances = strings.ReplaceAll(instances, string("| - " + name + " @ " + version + " | good"), string("| - " + name + " @ " + version + " | bad"))
            }
        } else if string(s[i]) == "*" {
            if isInt(string(stack[0])) && isInt(string(stack[1])) {
                s0, _ = strconv.Atoi(string(stack[0]))
                s1, _ = strconv.Atoi(string(stack[1]))
                stack += string(s0 * s1)
            } else {
                stack += "."
                instances = strings.ReplaceAll(instances, string("| - " + name + " @ " + version + " | failing"), string("| - " + name + " @ " + version + " | bad"))
                instances = strings.ReplaceAll(instances, string("| - " + name + " @ " + version + " | good"), string("| - " + name + " @ " + version + " | bad"))
            }
        } else if string(s[i]) == "#" {
            stack = ""
        } else if string(s[i]) == "~" {
            innerexec(fn, name)
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
                instances = strings.ReplaceAll(instances, string("| - " + name + " @ " + version + " | good"), string("| - " + name + " @ " + version + " | failing"))
                instances = strings.ReplaceAll(instances, string("| - " + name + " @ " + version + " | bad"), string("| - " + name + " @ " + version + " | failing"))
            }
        }
    }
    i += 1
}
func new_instance(name string) {
    fmt.Println("New instance initialized...")
    fmt.Println("\"" + name + "\" finished loading.")
    instances += string("| - " + name + " @ " + version + "\n")
    fmt.Println("Ark 1.9.0 in \"dev/terminal/instance_editor\"\nShards internal Ark code editor\ndev/terminal/instance_editor $~ ")
    fmt.Scanln(&ipt)
    go innerexec(ipt, name)
}
func main() {
    fmt.Print("\033[H\033[2J")
	fmt.Println("Shards 1.4.4\nType \"help\" or \"license\" to get started.")
    for {
        fmt.Print("dev/terminal/usr > ")
        fmt.Scanln(&ipt)
        if ipt == "instances" {
            fmt.Println(instances)
        } else if ipt == "help" {
            fmt.Println("go to the github repo for a full help section\nType help for this screen\nType license for the license\nType version for the current version of Shards\nType instances for a list of current instances running\nType ark for a mini Ark interpreter\nType new to create a new instance\nType rm to remove an instance by name\nType view to see the output of all instances\nType clear to clear the \"terminal\"")
        } else if ipt == "license" {
            fmt.Println("just dont sell itg\nor claim it as your own\nbecause if you do\ni will eat spaghetti\nand be sad at you")
        } else if ipt == "version" {
            fmt.Println(version)
        } else if ipt == "ark" {
            fmt.Print("Ark 2.0.0 in \"dev/terminal/usr/ark\"\nShards internal mini-runner\ndev/terminal/usr/ark $~ ")
            fmt.Scanln(&ipt)
            innerexec(ipt, "internal")
            fmt.Println(out1)
            out1 = ""
            out += out1
        } else if ipt == "new" {
            fmt.Println("Name of new instance: ")
            fmt.Scanln(&ipt)
            new_instance(ipt)
        } else if ipt == "rm" {
            fmt.Println("Name of instance to remove:")
            fmt.Scanln(&ipt)
            instances = strings.ReplaceAll(instances, string("| - " + ipt + " @ " + version + " | good"), "")
            instances = strings.ReplaceAll(instances, string("| - " + ipt + " @ " + version + " | bad"), "")
            instances = strings.ReplaceAll(instances, string("| - " + ipt + " @ " + version + " | failing"), "")
        } else if ipt == "clear" {
            fmt.Print("\033[H\033[2J")
        } else if ipt == "view" {
            fmt.Println(out)
        } else if ipt == "debug" {
            fmt.Println(version + "\n" + out + "\n" + instances + "\n")
        }
    }
}
