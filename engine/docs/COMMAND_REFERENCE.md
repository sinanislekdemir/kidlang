# 🎓 KidLang Command Reference Guide

Complete reference for all KidLang commands with examples and tips.

## 📚 Table of Contents

1. [Language Selection](#language-selection)
2. [Variables and Data](#variables-and-data)
3. [Input and Output](#input-and-output)
4. [Math Operations](#math-operations)
5. [String Operations](#string-operations)
6. [Control Flow](#control-flow)
7. [File Operations](#file-operations)
8. [System Commands](#system-commands)
9. [Special Values](#special-values)

---

## Language Selection

Set the language at the start of your program.

### Syntax
```kidlang
EN    // English
TR    // Turkish (Türkçe)
FI    // Finnish (Suomi)
DE    // German (Deutsch)
```

### Example
```kidlang
EN
print Hello World!
```

```kidlang
TR
yaz Merhaba Dünya!
```

**Note:** Language selection must be on the first line if used.

---

## Variables and Data

### Box (Simple Variable)

Store a single value: number, text, or calculation result.

**Syntax:**
```kidlang
box variable_name = value
```

**Examples:**
```kidlang
// Numbers
box age = 15
box price = 9.99
box temperature = -5

// Text (quotes optional for simple words)
box name = Alice
box greeting = "Hello, World!"

// Calculations
box result = 10 + 5
box total = box price * 2

// Copy from another box
box new_age = box age
```

**Using box values:**
```kidlang
print box age              // Display the value
box doubled = box age * 2  // Use in calculations
ask Your age?
box age = answer           // Store user input
```

### Stack (List/Dictionary)

Store multiple values with keys or indexes.

**Syntax:**
```kidlang
stack stack_name
stack stack_name[key] = value
value = stack stack_name[key]
```

**Index styles:**
```kidlang
// Both [] and () work the same
stack items[1] = apple
stack items(2) = banana

// Can use numbers or words as keys
stack data[1] = first
stack data[name] = Alice
stack data[score] = 100
```

**Examples:**
```kidlang
// Numbered list
stack scores
stack scores[1] = 95
stack scores[2] = 87
stack scores[3] = 92
print Test 1: stack scores[1]

// Dictionary-style
stack player
stack player[name] = Hero
stack player[health] = 100
stack player[level] = 5
print Player: stack player[name]
print Health: stack player[health]

// Using variables as keys
box test_num = 2
print Score: stack scores[test_num]

// Mixed keys
stack inventory
stack inventory[1] = sword
stack inventory[2] = shield
stack inventory[gold] = 500
```

### File

Declare a file handle for file operations.

**Syntax:**
```kidlang
file file_handle
```

**Example:**
```kidlang
file mydata
```

---

## Input and Output

### Print

Display text, numbers, or values on the screen.

**Syntax:**
```kidlang
print message
print "message"
print box variable
print text1 text2 text3
```

**Examples:**
```kidlang
// Simple text
print Hello, World!
print "Hello, World!"

// Variables
box name = Alice
print box name
print Hello, box name!

// Multiple items
box age = 15
print Name: box name Age: box age

// Numbers and calculations
print 5 + 3
print The result is: 5 + 3

// Empty line
print
```

**Implicit print:**
```kidlang
// Any line without a command prints automatically
Hello World!          // Same as: print Hello World!
This is KidLang       // Same as: print This is KidLang
```

### Ask

Get input from the user.

**Syntax:**
```kidlang
ask question
ask "question"
```

**Examples:**
```kidlang
// Simple question
ask What is your name?
box name = answer

// With prompt
ask "Enter your age: "
box age = answer

// Using the answer
ask Pick a number:
print You picked: answer

// Multiple inputs
ask Enter first number:
box num1 = answer
ask Enter second number:
box num2 = answer
box sum = box num1 + box num2
print Total: box sum
```

**Important:** The user's response is always stored in the special variable `answer`.

---

## Math Operations

### Basic Operators

| Operator | Meaning | Example | Result |
|----------|---------|---------|--------|
| `+` | Addition | `5 + 3` | `8` |
| `-` | Subtraction | `10 - 4` | `6` |
| `*` | Multiplication | `6 * 7` | `42` |
| `/` | Division | `20 / 4` | `5` |
| `%` | Modulo (remainder) | `10 % 3` | `1` |
| `^` | XOR (numbers) | `5 ^ 3` | `6` |

**Examples:**
```kidlang
box a = 10
box b = 3

box sum = box a + box b          // 13
box diff = box a - box b         // 7
box product = box a * box b      // 30
box quotient = box a / box b     // 3
box remainder = box a % box b    // 1
```

### Math Functions

**Square Root:**
```kidlang
print sqrt 16        // 4
box result = sqrt 25 // 5
```

**Absolute Value:**
```kidlang
print abs -15        // 15
box result = abs -7  // 7
```

**Square:**
```kidlang
print sqr 5          // 25
box result = sqr 8   // 64
```

**Trigonometry:**
```kidlang
print sin 1.57       // ~1 (sin of π/2)
print cos 0          // 1
print tan 0.785      // ~1 (tan of π/4)
```

**Logarithm:**
```kidlang
print log 2.718      // ~1 (natural log of e)
box result = log 10
```

**Inverse Trigonometry:**
```kidlang
print asin 0.5       // ~0.524 (30 degrees in radians)
print acos 0.5       // ~1.047 (60 degrees in radians)
```

---

## String Operations

### Concatenation (+)
Join strings together.

```kidlang
box first = Hello
box last = World
box greeting = box first + box last    // "HelloWorld"
box spaced = box first +  + box last   // "Hello World"
```

### Character Removal (-)
Remove all occurrences of a character.

```kidlang
box word = hello
box result = box word - l              // "heo"

box text = programming
box result = box text - m              // "prograing"
```

### String Repeat (*)
Repeat a string multiple times.

```kidlang
box star = *
box line = box star * 10               // "**********"

box ha = ha
box laugh = box ha * 3                 // "hahaha"
```

### Character Access (/)
Get a character at a specific position (1-indexed).

```kidlang
box word = hello
box first = box word / 1               // "h"
box third = box word / 3               // "l"
box last = box word / 5                // "o"
```

### XOR Cipher (^)
Simple encryption using XOR.

```kidlang
box message = secret
box key = 123
box encrypted = box message ^ box key
```

---

## Control Flow

### If Statement

Execute code based on conditions.

**Simple if:**
```kidlang
if condition then
    // code to run if true
end
```

**Examples:**
```kidlang
// Single line
if box age >= 18 then print You are an adult
end

// Multi-line block
if box score > 90 then
    print Excellent!
    print You got an A!
end

// Multiple conditions
box age = 15
box has_ticket = 1

if box age > 12 then
    print Old enough
end

if box has_ticket = 1 then
    print You can enter
end
```

### Comparison Operators

| Operator | Meaning | Example |
|----------|---------|---------|
| `=` | Equal to | `box a = 5` |
| `!=` | Not equal to | `box a != 10` |
| `>` | Greater than | `box a > 3` |
| `<` | Less than | `box a < 20` |
| `>=` | Greater or equal | `box a >= 5` |
| `<=` | Less or equal | `box a <= 15` |

### Logical Operators

**AND** - Both conditions must be true:
```kidlang
if box age > 18 and box score > 80 then
    print You qualify!
end
```

**OR** - At least one condition must be true:
```kidlang
if box age < 5 or box age > 65 then
    print You get a discount
end
```

**Complex conditions:**
```kidlang
if box a > 10 and box b < 20 and box c = 5 then
    print All conditions met!
end

if box x = 1 or box y = 2 or box z = 3 then
    print At least one is true
end
```

### Labels and Goto

Jump to different parts of your program.

**Syntax:**
```kidlang
label_name:
// code here

goto label_name
```

**Examples:**
```kidlang
// Simple jump
start:
print Hello
goto end

print This won't print
end:
print Goodbye

// Loop with goto
box count = 5
loop:
if box count <= 0 then goto done
print box count
box count = box count - 1
goto loop
done:

// Menu system
menu:
print 1. Play
print 2. Exit
ask Choose:
if answer = 1 then goto play
if answer = 2 then goto exit
goto menu

play:
print Playing game...
goto menu

exit:
print Goodbye!
```

**Conditional goto:**
```kidlang
if box score > 100 then goto winner
if box lives = 0 then goto game_over
```

---

## File Operations

### Open
Open or create a file.

**Syntax:**
```kidlang
open file_handle filename
```

**Example:**
```kidlang
file myfile
open myfile data.txt
```

### Read
Read entire file into a box.

**Syntax:**
```kidlang
read file_handle box variable
```

**Example:**
```kidlang
file myfile
open myfile data.txt
read myfile box content
print box content
close myfile
```

### Readline
Read one line from file.

**Syntax:**
```kidlang
readline file_handle box variable
```

**Example:**
```kidlang
file myfile
open myfile data.txt
readline myfile box line1
readline myfile box line2
print First line: box line1
print Second line: box line2
close myfile
```

### Write
Write text to file.

**Syntax:**
```kidlang
write file_handle text
```

**Example:**
```kidlang
file myfile
open myfile output.txt
write myfile Hello, World!
write myfile This is line 2
close myfile
```

### Seek
Jump to a specific line number.

**Syntax:**
```kidlang
seek file_handle line_number
```

**Example:**
```kidlang
file myfile
open myfile data.txt
seek myfile 5           // Jump to line 5
readline myfile box line
print Line 5: box line
close myfile
```

### Close
Close an open file.

**Syntax:**
```kidlang
close file_handle
```

**Example:**
```kidlang
file myfile
open myfile data.txt
// ... work with file ...
close myfile
```

### Complete File Example

```kidlang
// Write to file
file output
open output notes.txt
write output Task 1: Learn KidLang
write output Task 2: Build a project
write output Task 3: Have fun!
close output

// Read from file
file input
open input notes.txt
read input box all_content
print === File Contents ===
print box all_content
close input

// Read line by line
file input2
open input2 notes.txt
readline input2 box line1
readline input2 box line2
readline input2 box line3
print Line 1: box line1
print Line 2: box line2
print Line 3: box line3
close input2
```

---

## System Commands

### Exec
Execute a shell command.

**Syntax:**
```kidlang
exec command
exec "command with args"
```

**Examples:**
```kidlang
// List files
exec ls

// List with options
exec ls -la

// Create directory
exec mkdir mydir

// Copy file
exec cp file1.txt file2.txt

// Display file
exec cat data.txt

// Clear screen
exec clear
```

**Note:** Be careful with system commands as they can affect your system!

### Sleep
Pause program execution.

**Syntax:**
```kidlang
sleep milliseconds
```

**Examples:**
```kidlang
// Wait 1 second
sleep 1000

// Wait half a second
sleep 500

// Countdown
box i = 3
countdown:
if box i < 0 then goto done
print box i
sleep 1000
box i = box i - 1
goto countdown
done:
print GO!
```

---

## Special Values

### random
Get a random integer.

**Examples:**
```kidlang
// Random number
box num = random
print Random: box num

// Random in range 1-10
box dice = random % 10 + 1
print Dice roll: box dice

// Random in range 1-100
box percent = random % 100 + 1
print Random percent: box percent

// Coin flip
box coin = random % 2
if box coin = 0 then print Heads
if box coin = 1 then print Tails
```

### now
Get current date and time.

**Example:**
```kidlang
print Current time: now

box timestamp = now
print Timestamp: box timestamp
```

### answer
Special variable that stores user input from `ask`.

**Example:**
```kidlang
ask What is your name?
print Hello, answer!

box name = answer
print Saved: box name
```

### Newline (\n)
Insert a line break.

**Examples:**
```kidlang
print Line 1\nLine 2\nLine 3

box text = Hello\nWorld
print box text
```

---

## Tips and Best Practices

### Comments
```kidlang
// This is a comment - use them to explain your code
box age = 15  // Store the user's age
```

### Variable Naming
```kidlang
// Good names
box player_score = 100
box max_health = 50
box is_playing = 1

// Avoid single letters (except for loops)
box i = 1      // OK for loop counter
box x = 10     // Less clear
```

### Organize Your Code
```kidlang
// === SETUP ===
box health = 100
box score = 0

// === MAIN GAME ===
game_loop:
// game code here
goto game_loop

// === END ===
game_over:
print Game Over!
```

### Error Checking
```kidlang
ask Enter a positive number:
box num = answer

if box num <= 0 then
    print Error: Must be positive!
    goto end
end

// Continue with valid number
print You entered: box num

end:
```

---

## Quick Reference Table

| Category | Commands |
|----------|----------|
| **Variables** | `box`, `stack`, `file` |
| **I/O** | `print`, `ask` |
| **Math** | `+`, `-`, `*`, `/`, `%`, `^`, `sqrt`, `abs`, `sqr`, `sin`, `cos`, `tan`, `log` |
| **Control** | `if...then...end`, `goto`, `label:` |
| **Logic** | `and`, `or`, `=`, `!=`, `>`, `<`, `>=`, `<=` |
| **Files** | `open`, `read`, `readline`, `write`, `seek`, `close` |
| **System** | `exec`, `sleep` |
| **Special** | `random`, `now`, `answer`, `\n` |

---

## Language Keywords by Language

| English | Turkish | Finnish | German |
|---------|---------|---------|--------|
| box | kutu | laatikko | kiste |
| stack | liste | lista | liste |
| file | dosya | tiedosto | datei |
| print | yaz | tulosta | schreib |
| ask | sor | kysy | frag |
| if | eger | jos | wenn |
| then | ise | niin | dann |
| end | son | loppu | ende |
| goto | git | mene | geh |
| and | ve | ja | und |
| or | veya | tai | oder |
| open | ac | avaa | oeffne |
| close | kapat | sulje | schliesse |
| read | oku | lue | lies |
| write | yaz | kirjoita | schreib |
| exec | calistir | suorita | fuehreaus |
| sleep | bekle | odota | warte |
| answer | cevap | vastaus | antwort |
| random | rastgele | satunnainen | zufall |
| now | tarih | aika | zeit |

---

**Happy coding! 🎉**
