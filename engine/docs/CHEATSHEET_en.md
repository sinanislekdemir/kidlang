# Kidlang Cheatsheet (English)

## Language Selection
```
EN
```

## Comments
```
// This is a comment
```

## Variables

### Box (simple variable)
```
box name = 42
box name = 3.14
box name = hello
box name = box other
```

### Stack (dictionary/map)
```
stack toys
stack toys[1] = car        // or toys(1) = car
stack toys[2] = ball       // or toys(2) = ball
toys[1] = truck            // or toys(1) = truck
print toys[1]              // or print toys(1)
```

### File
```
file myfile
```

## Output
```
print Hello World          // Quotes optional
print "Hello World"        // Quotes preserve literal text
print box name             // Prints value of variable 'name'
print "box name"           // Prints literal "box name"
print 1 + 2
Any text without command   // Implicit print
```

## Input
```
ask What is your name?     // Prompts user for input
print answer               // Result stored in 'answer' variable
ask "Enter value:"         // Quotes optional
box x = answer             // Use the answer
```

## Math Operations
```
box result = 5 + 3
box result = 10 - 2
box result = 4 * 3
box result = 10 / 2
box result = 10 % 3
box result = 2 ^ 3  // XOR for numbers, cipher for strings
```

## Math Functions
```
sqrt 16       // Square root
abs -5        // Absolute value
sqr 4         // Square (4*4)
sin 1.57      // Sine
cos 0         // Cosine
tan 0.785     // Tangent
log 2.718     // Natural logarithm
asin 0.5      // Arcsine
acos 0.5      // Arccosine
```

## String Operations
```
box text = hello + world   // Concatenation (quotes optional)
box text = hello - l       // Remove all occurrences of character
box text = abc * 3         // Repeat string
box char = hello / 2       // Get character at index 2
```

## Conditionals
```
if box x = 5 then print x is 5
end

if box x > 10 then
print x is greater than 10
end

if box x < 5 then goto skip
```

### Comparison Operators
```
=    // Equal
!=   // Not equal
>    // Greater than
<    // Less than
>=   // Greater or equal
<=   // Less or equal
```

### Logical Operators
```
if box x > 5 and box y < 10 then print both true
end

if box a = 1 or box b = 2 then print one is true
end
```

## Labels and Jumps
```
start:
print "Hello"
goto start

skip:
print "Jumped here"
```

## File Operations
```
file myfile
open myfile data.txt       // Opens/creates file
read myfile box content    // Read entire file into box
read myfile stack lines    // Read file line-by-line into stack (auto-typed)
readline myfile box line   // Read one line
write myfile some text     // Write text to file
write myfile stack data    // Write stack to file (one item per line)
seek myfile 5              // Seek to line 5
close myfile               // Close file
```

### Stack-File Integration
When reading a file into a stack, each line becomes an item with auto-typing:
- Pure integers like "42" → integer
- Floats like "3.14" → float
- Text like "hello" → string

```
file input
open input data.txt
read input stack lines     // Lines[1], lines[2], etc.
print stack lines[1]
close input

stack output
stack output[1] = value1
stack output[2] = value2
file out
open out result.txt
write out stack output     // Writes each value as a line
close out
```

## Special Values
```
random  // Random integer
now     // Current date/time
\n      // Newline character
```

## System Commands
```
exec ls -la          // Execute shell command (quotes optional)
sleep 1000           // Sleep for 1000 milliseconds
```

## Execution Flow
- Programs execute line by line from top to bottom
- `goto` jumps to a label
- `if...then...end` creates conditional blocks
- `end` closes an `if` block
