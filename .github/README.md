# ark
![logo](https://esolangs.org/w/images/c/c9/Ark_logo.png)
# The esoteric programming language that is *somewhat* useful.
#### All details are also on the [esolangs.org wiki page](https://esolangs.org/wiki/ark).
### Usage

 - `:` Terminates the program. | **Required**

- `!`  Outputs the next Alphanumeric character in the code. | **Output**

- `^`  Outputs the stack. | **Output**

- `>`  Adds the next numeric character to the stack. | **Stack Manipulation**

- `<`  Asks for Input and adds it to the stack. (the only way to add multiple digits to the same row) | **Input**/**Stack Manipulation**

- `+`  Add the top two rows in the stack and add the result to the stack. | **Math**/**Stack Manipulation**

- `-`  Subtract the second row by the top row in the stack and add the result to the stack. | **Math**/**Stack Manipulation**

- `*`  Multiply the top two rows of the stack and add the result to the stack. | **Math**/**Stack Manipulation**

- `)`  Subtract the top row by the second row in the stack and add the result to the stack. | **Math**/**Stack Manipulation**

- `/`  Divide the top row by the second row in the stack and add the result to the stack. | **Math**/**Stack Manipulation**

- `(`  Divide the second row by the top row in the stack and add the result to the stack. | **Math**/**Stack Manipulation**

- `#`  Clear the stack. | **Stack Manipulation**

- `$`  Add a random number in between the top row and the second row in the stack. | **Stack Manipulation**

- `%`  Add a random number in between the second row and the top row in the stack. | **Stack Manipulation**
- `{` Used to initialize a function. (like python's: "def") | **Functions**
- `}` Used to end the initialization of a function. | **Functions**
- `~` Used to call a function. | **Functions**
### Examples:
#### A Hello World program:
`!H !e !l !l !o !  !W !o !r !l !d:`
##### Each character is individually read then printed
#### Prints `Hello World`
#### Mathematics:
`>1>1+^:`
##### Adds 1 to the stack twice, then adds the sum of the top two digits in the stack to the top of the stack, resulting in 1, 1, 2.
#### Prints: `112`
#### Infinite loop:
`{~}~:`
##### Similar to python's `s = "exec(s)"; exec(s)` the function's code is to run itself, then the function is called, then the function is called, resulting in an infinite loop¹.
###### ¹there may be a recursion limit depending on the IDE and/or compiler
