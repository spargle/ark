![logo](https://github.com/spargle/ark/blob/main/.github/ark%20logo.png)
[![CodeFactor](https://www.codefactor.io/repository/github/spargle/ark/badge/main)](https://www.codefactor.io/repository/github/spargle/ark/overview/main)
# The programming language that is *somewhat* useful.
## install
##### Ark interpreter
```
curl -O https://xnxrijmguoktvvhvbpth.supabase.in/storage/v1/object/public/builds/ark.zip
```
##### Shards
```
curl -O https://xnxrijmguoktvvhvbpth.supabase.in/storage/v1/object/public/builds/shards_1.4.0.zip
```
### Using the interpreter:
#### It will look like: `$~ Ark x.x.x: <souce code>`
### Usage
- `!`  Outputs the next Alphanumeric character in the code. | **Output**
- `^`  Outputs the stack. | **Output**
- `>`  Adds the next numeric character to the stack. | **Stack Manipulation**
- `<`  Asks for Input and adds it to the stack. | **Input**/**Stack Manipulation**
- `+`  Add the top two rows in the stack and add the result to the stack. | **Math**/**Stack Manipulation**
- `-`  Subtract the second row by the top row in the stack and add the result to the stack. | **Math**/**Stack Manipulation**
- `*`  Multiply the top two rows of the stack and add the result to the stack. | **Math**/**Stack Manipulation**
- `/`  Divide the top row by the second row in the stack and add the result to the stack. | **Math**/**Stack Manipulation**
- `#`  Clear the stack. | **Stack Manipulation**
- `{` Used to initialize a function. (like python's: "def") | **Functions**
- `}` Used to end the initialization of a function. | **Functions** 
- `[` Used to initialize a multiple-character addition. | **Stack Manipulation**
- `]` Used to end the initialization of a multiple-character addition. | **Stack Manipulation**
- `~` Used to call a function. | **Functions**
- `` ` `` Used to execute a function alongside the main program | **Functions**
- `@` Stops the program if a runtime stability chech fails. | **Runtime Stability**
### Examples:
#### A Hello World program:
`!H!e!l!l!o! !W!o!r!l!d`
##### Each character is individually read then printed
#### Prints `Hello World`
#### Mathematics:
`@>1>1+^:`
##### Adds 1 to the stack twice, then adds the sum of the top two digits in the stack to the top of the stack, resulting in 1, 1, 2.
#### Prints: `112`
#### Infinite loop:
`{~}~`
##### Similar to python's `s = "exec(s)"; exec(s)` the function's code is to run itself, then the function is called, then the function is called, resulting in an infinite loop.
