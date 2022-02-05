# Ark Docs
## Welcome to Ark Docs! This is where you can find [code instructions](https://github.com/spargle/ark/blob/main/.github/docs.md#instructions), [installation instructions](https://github.com/spargle/ark/blob/main/.github/docs.md#installation), and [how-to guides](https://github.com/spargle/ark/blob/main/.github/docs.md#how-to-ark).
# Installation:
## Download this repository:
```
curl -O https://ffcgytupxhgxsidfzvbe.supabase.in/storage/v1/object/public/zip/Ark/ark.zip
```
### Paste or type your code where it says `$~ Ark x.x.x:` and hit enter.
### It will then run your code! (It's a beautiful thing, isn't it.)

# How To Ark:
## Open up any text editor, Atom, Google Docs, Notes, Even the Google Search Bar works.
### The basic Hello, World program.
### `!H!e!l!l!o! !W!o!r!l!d`
### Cool, huh? Each `!` is telling the interpreter that the next character will be printed to STDOUT, or be written on the screen.
### There is another way to do this, though. Using the Stack.
### The stack is a big list of numbers and characters. It is the only "variable" in Ark.
### To push something to the stack, simply do `>a`. This will push "a" to the stack. The character after `>` can be anything that is not a Ark key character like: `!`, `#`, or any other character that is used by the interpreter.
### "But, what if I want to push a number that isn't less than ten?" you may ask, well the answer is: Brackets!!!
### `[1873498]` pushes 1, 8, 7, 3, 4, 9, and 8 to the stack! How cool!
### To show the output of the stack: `^`. This "prints" the stack as a big fat number.
## The full, stack-based Hello, World program:
### ```[Hello, World!]^```
### Ta-Da!
# Instructions
## The `!` instruction:
### The `!` instruction prints the next character (the character after the `!`)
#### Example: `!A` prints "A".

## The `^` instruction:
### The `^` instruction prints the stack.
#### Example: `>A >1 ^` prints "A1"

## The `>` instruction:
### The `>` instruction adds the next character to the stack.
#### Example: `>A >1` adds "A" and 1 to the stack.

## The `<` instruction:
### The `<` instruction allows the user to input text and then push it to the stack.
#### Example: `<` when the user types: `yeet`, pushes "yeet" to the stack.

## The `+, -, *, and /` instruction:
### The `+` instructions adds the top two items in the stack and pushes the sum to the stack.
#### Example: `>1 >2 +` pushes 1 and 2 to the stack, then pushes 3 to the stack.
### The `-` instructions subtracts the second item from the top item in the stack and pushes the difference to the stack.
#### Example: `>1 >2 -` pushes 1 and 2 to the stack, then pushes -1 to the stack.
### The `*` instructions multiplys the top two items in the stack and pushes the product to the stack.
#### Example: `>1 >2 *` pushes 1 and 2 to the stack, then pushes 2 to the stack.
### The `/` instructions divides the second item by the top item in the stack and pushes the quotient to the stack.
#### Example: `>1 >2 /` pushes 1 and 2 to the stack, then pushes 0.5 to the stack.
### Troubleshooting:
#### Make sure the top two items in the stack are integers and not characters.
#### Change the position of the math instruction to change what is pushed to the stack.

## The `#` instruction:
### The `#` instruction clears the stack.
#### Example: `>A >B # ^` prints nothing because the `#` deleted everything on the stack.

## The `{, }, and ~` instructions:
### The `{, }, and ~` instructions are used to define and run a function. Anything inbetween the two curly brackets is ignored and not run until the `~` command is run by the interpreter.
#### Example: `{>1} >5 ~` pushes a 5 then a 1 to the stack (the stack looks like this: `5, 1` because the 1 isnt pushed untill the `{}` function is called.).

## The `[ and ]` instructions:
### The `[ and ]` instructions push whatever is inbetween them to the stack.
#### Example: `[123] > 4` pushes 123 and 4 to the stack.

# This is the end.
