# ark
![logo](https://esolangs.org/w/images/c/c9/Ark_logo.png)
[![Download ark](https://sourceforge.net/sflogo.php?type=16&group_id=3394501)](https://sourceforge.net/p/ark-pl/)
# The esoteric programming language that is *somewhat* useful.
#### All details are also on the [esolangs.org wiki page](https://esolangs.org/wiki/ark).
### Usage

 - `:`  Terminates the program. | **Required**

- `!`  **Output**s the next Alphanumeric character in the code. | **Output**

- `^`  **Output**s the stack. | **Output**

- `>`  Adds the next numeric character to the stack. | **Stack Manipulation**

- `<`  Asks for **Input** and adds it to the stack. (the only way to add multiple digits to the same row) | **Input**/**Stack Manipulation**

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
### examples:
#### A Hello World program:
`!H !e !l !l !o !  !W !o !r !l !d:`
#### prints `Hello World`
#### Mathematics:
`>1>1+^:`
#### prints: `112`

#### No other examples have been created.
