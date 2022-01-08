package main
import (
    "fmt"
    "strconv"
    "unicode"
    "strings"
)
var (
    ipt = ""
    instances = ""
    version = "1.0.0"
    out = ""
    out1 = ""
)
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
            fn = strings.ReplaceAll(fn, "{}", "")
        }
        i += 1
    }
}
func new_instance(name string) {
    fmt.Println("New instance initialized...")
    fmt.Println("\"" + name + "\" finished loading.")
    instances += string("| - " + name + " @ " + version + "\n")
    fmt.Println("Ark 1.8.5 in \"dev/terminal/instance_editor\"\nShards internal Ark code editor\ndev/terminal/instance_editor $~ ")
    fmt.Scanln(&ipt)
    go innerexec(ipt)
}
func main() {
    fmt.Print("\033[H\033[2J")
	fmt.Println("Shards 1.0.0\nType \"help\" or \"license\" to get started.")
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
            fmt.Print("Ark 1.8.4 in \"dev/terminal/usr/ark\"\nShards internal mini-runner\ndev/terminal/usr/ark $~ ")
            fmt.Scanln(&ipt)
            innerexec(ipt)
            fmt.Println(out1)
            out1 = ""
        } else if ipt == "new" {
            fmt.Println("Name of new instance: ")
            fmt.Scanln(&ipt)
            new_instance(ipt)
        } else if ipt == "rm" {
            fmt.Println("Name of instance to remove:")
            fmt.Scanln(&ipt)
            instances = strings.ReplaceAll(instances, string("| - " + ipt + " @ " + version + "\n"), "")
        } else if ipt == "clear" {
            fmt.Print("\033[H\033[2J")
        } else if ipt == "view" {
            fmt.Println(out)
        }
    }
}
