# 🎓 KidLang Algorithm Tutorial

Learn to build real algorithms and solve programming challenges!

## 📚 Table of Contents

1. [What is an Algorithm?](#what-is-an-algorithm)
2. [Searching Algorithms](#searching-algorithms)
3. [Sorting Algorithms](#sorting-algorithms)
4. [Math Algorithms](#math-algorithms)
5. [String Algorithms](#string-algorithms)
6. [Game Algorithms](#game-algorithms)
7. [Challenge Problems](#challenge-problems)

---

## What is an Algorithm?

An **algorithm** is a step-by-step recipe for solving a problem. Just like following a recipe to bake cookies, an algorithm follows steps to solve a problem!

**Example: Making a Sandwich**
1. Get two slices of bread
2. Add peanut butter to one slice
3. Add jelly to the other slice
4. Put the slices together
5. Done!

Let's learn to write algorithms in KidLang!

---

## Searching Algorithms

### Linear Search

Find if a number exists in a list by checking each item one by one.

```kidlang
print === Linear Search ===

// Create our list of numbers
stack numbers
stack numbers[1] = 5
stack numbers[2] = 12
stack numbers[3] = 8
stack numbers[4] = 15
stack numbers[5] = 3

ask What number do you want to find?
box target = answer

// Search through the list
box i = 1
box found = 0
box position = 0

search_loop:
if box i > 5 then
    goto search_done
end

if stack numbers[i] = box target then
    box found = 1
    box position = box i
    goto search_done
end

box i = box i + 1
goto search_loop

search_done:
if box found = 1 then
    print Found box target at position box position!
end

if box found = 0 then
    print box target not found in the list
end
```

**How it works:**
1. Start at the first item
2. Check if it matches what we're looking for
3. If yes, we found it! If no, move to next item
4. Repeat until we check all items

### Find Maximum Number

Find the largest number in a list.

```kidlang
print === Find Maximum ===

stack numbers
stack numbers[1] = 23
stack numbers[2] = 67
stack numbers[3] = 45
stack numbers[4] = 89
stack numbers[5] = 12

// Start with the first number as max
box max = stack numbers[1]
box i = 2

find_max:
if box i > 5 then
    goto done
end

// If current number is bigger, it's the new max
if stack numbers[i] > box max then
    box max = stack numbers[i]
end

box i = box i + 1
goto find_max

done:
print The maximum number is: box max
```

---

## Sorting Algorithms

### Bubble Sort (Simplified)

Sort numbers from smallest to largest.

```kidlang
print === Bubble Sort Demo ===

// Our unsorted list
stack nums
stack nums[1] = 64
stack nums[2] = 34
stack nums[3] = 25
stack nums[4] = 12
stack nums[5] = 22

print Original list:
print stack nums[1]
print stack nums[2]
print stack nums[3]
print stack nums[4]
print stack nums[5]

// Bubble sort - multiple passes
box pass = 1

outer:
if box pass >= 5 then
    goto sort_done
end

box i = 1

inner:
box next = box i + 1
if box next > 5 then
    box pass = box pass + 1
    goto outer
end

// Compare adjacent numbers
if stack nums[i] > stack nums[next] then
    // Swap them
    box temp = stack nums[i]
    stack nums[i] = stack nums[next]
    stack nums[next] = box temp
end

box i = box i + 1
goto inner

sort_done:
print
print Sorted list:
print stack nums[1]
print stack nums[2]
print stack nums[3]
print stack nums[4]
print stack nums[5]
```

**How it works:**
1. Compare pairs of numbers
2. If the first is bigger than the second, swap them
3. Keep doing this until no more swaps are needed
4. The list is now sorted!

---

## Math Algorithms

### Prime Number Checker

Check if a number is prime (only divisible by 1 and itself).

```kidlang
print === Prime Number Checker ===

ask Enter a number:
box num = answer

// Numbers less than 2 are not prime
if box num < 2 then
    print box num is not prime
    goto end
end

// 2 is prime
if box num = 2 then
    print box num is prime!
    goto end
end

// Check if divisible by any number from 2 to num-1
box i = 2
box is_prime = 1

check:
if box i >= box num then
    goto check_done
end

box remainder = box num % box i

if box remainder = 0 then
    box is_prime = 0
    goto check_done
end

box i = box i + 1
goto check

check_done:
if box is_prime = 1 then
    print box num is prime!
end

if box is_prime = 0 then
    print box num is not prime
end

end:
```

### Factorial Calculator

Calculate factorial (e.g., 5! = 5 × 4 × 3 × 2 × 1 = 120)

```kidlang
print === Factorial Calculator ===

ask Enter a number:
box n = answer

if box n < 0 then
    print Factorial is not defined for negative numbers
    goto end
end

box result = 1
box i = box n

calculate:
if box i <= 1 then
    goto done
end

box result = box result * box i
box i = box i - 1
goto calculate

done:
print box n! = box result

end:
```

### Fibonacci Sequence

Generate Fibonacci numbers (each number is the sum of the previous two).

```kidlang
print === Fibonacci Sequence ===

ask How many Fibonacci numbers do you want?
box count = answer

// First two numbers
box a = 0
box b = 1
box i = 1

print 1: box a

if box count = 1 then
    goto done
end

print 2: box b
box i = 2

fib:
if box i >= box count then
    goto done
end

box next = box a + box b
box i = box i + 1
print box i: box next

box a = box b
box b = box next
goto fib

done:
print Fibonacci sequence complete!
```

### Greatest Common Divisor (GCD)

Find the GCD of two numbers using Euclidean algorithm.

```kidlang
print === GCD Calculator ===

ask Enter first number:
box a = answer

ask Enter second number:
box b = answer

// Keep the original values for display
box orig_a = box a
box orig_b = box b

gcd_loop:
if box b = 0 then
    goto gcd_done
end

box temp = box b
box b = box a % box b
box a = box temp
goto gcd_loop

gcd_done:
print GCD of box orig_a and box orig_b is: box a
```

---

## String Algorithms

### Character Counter

Count how many times a character appears in text.

```kidlang
print === Character Counter ===

box text = programming
box search_char = g

print Looking for 'box search_char' in 'box text'

box i = 1
box count = 0
box length = 11

count_loop:
if box i > box length then
    goto count_done
end

box current = box text / box i

if box current = box search_char then
    box count = box count + 1
end

box i = box i + 1
goto count_loop

count_done:
print Found box search_char a total of box count times
```

### Palindrome Checker

Check if a word reads the same forwards and backwards.

```kidlang
print === Palindrome Checker ===

ask Enter a word:
box word = answer

// We'll check characters manually
// For this example, let's check if "radar" is palindrome
box word = radar

box char1 = box word / 1
box char2 = box word / 2
box char3 = box word / 3
box char4 = box word / 4
box char5 = box word / 5

// Compare first with last, second with fourth
box is_palindrome = 1

if box char1 != box char5 then
    box is_palindrome = 0
end

if box char2 != box char4 then
    box is_palindrome = 0
end

if box is_palindrome = 1 then
    print box word is a palindrome!
end

if box is_palindrome = 0 then
    print box word is not a palindrome
end
```

### String Reverser

Reverse a string character by character.

```kidlang
print === String Reverser ===

box text = hello
box length = 5

print Original: box text

box reversed = 
box i = box length

reverse:
if box i < 1 then
    goto done
end

box char = box text / box i
box reversed = box reversed + box char
box i = box i - 1
goto reverse

done:
print Reversed: box reversed
```

---

## Game Algorithms

### Random Number Generator Game

Generate random numbers and track statistics.

```kidlang
print === Random Number Stats ===

box count = 10
box i = 1
box sum = 0
box max = 0
box min = 999

generate:
if box i > box count then
    goto analyze
end

box num = random % 100
print Number box i: box num

// Update sum
box sum = box sum + box num

// Update max
if box num > box max then
    box max = box num
end

// Update min
if box num < box min then
    box min = box num
end

box i = box i + 1
goto generate

analyze:
box average = box sum / box count

print
print Statistics:
print Minimum: box min
print Maximum: box max
print Average: box average
```

### Simple AI Opponent

Create a simple AI that makes decisions based on game state.

```kidlang
print === Rock Paper Scissors AI ===

ask Choose (1=Rock, 2=Paper, 3=Scissors):
box player = answer

// AI makes a choice
box ai_choice = random % 3 + 1

print You chose: box player
print AI chose: box ai_choice

// Determine winner
// 1=Rock, 2=Paper, 3=Scissors
// Rock beats Scissors, Scissors beats Paper, Paper beats Rock

if box player = box ai_choice then
    print It's a tie!
    goto end
end

// Check if player wins
box player_wins = 0

if box player = 1 and box ai_choice = 3 then
    box player_wins = 1
end

if box player = 2 and box ai_choice = 1 then
    box player_wins = 1
end

if box player = 3 and box ai_choice = 2 then
    box player_wins = 1
end

if box player_wins = 1 then
    print You win!
end

if box player_wins = 0 then
    print AI wins!
end

end:
```

### Score Ranking System

Rank players by their scores.

```kidlang
print === Player Rankings ===

// Player scores
stack scores
stack scores[alice] = 450
stack scores[bob] = 380
stack scores[charlie] = 520
stack scores[diana] = 490

print Original Scores:
print Alice: stack scores[alice]
print Bob: stack scores[bob]
print Charlie: stack scores[charlie]
print Diana: stack scores[diana]

// Find highest scorer (simplified)
box max_score = 0
box winner = nobody

if stack scores[alice] > box max_score then
    box max_score = stack scores[alice]
    box winner = Alice
end

if stack scores[bob] > box max_score then
    box max_score = stack scores[bob]
    box winner = Bob
end

if stack scores[charlie] > box max_score then
    box max_score = stack scores[charlie]
    box winner = Charlie
end

if stack scores[diana] > box max_score then
    box max_score = stack scores[diana]
    box winner = Diana
end

print
print Winner: box winner with box max_score points!
```

---

## Challenge Problems

Try to solve these on your own!

### Challenge 1: Sum of Digits
Write a program that takes a number and adds up all its digits.
Example: 123 → 1 + 2 + 3 = 6

### Challenge 2: Number Reverser
Take a number like 1234 and reverse it to 4321.

### Challenge 3: Leap Year Calculator
Check if a year is a leap year:
- Divisible by 4
- But NOT divisible by 100 (unless also divisible by 400)

### Challenge 4: Temperature Statistics
Store 7 daily temperatures and calculate:
- Average temperature
- Hottest day
- Coldest day
- Days above average

### Challenge 5: Simple Encryption
Create a cipher that shifts letters by 3 positions:
- A → D, B → E, C → F, etc.

### Challenge 6: Pattern Printer
Print patterns like:
```
*
**
***
****
*****
```

### Challenge 7: Number Classifier
Ask for a number and tell if it's:
- Positive or negative
- Even or odd
- Prime or composite
- Perfect square or not

### Challenge 8: Mini Database
Create a system to store and search student records:
- Add students with names and grades
- Find a student by name
- Calculate class average
- Find highest and lowest grade

---

## 🎯 Algorithm Design Tips

1. **Break it down**: Split big problems into smaller steps
2. **Use comments**: Explain what each part does
3. **Test with examples**: Try your algorithm with different inputs
4. **Draw it out**: Sketch the flow on paper first
5. **Start simple**: Get it working, then make it better

## 🚀 Next Steps

- Study the example programs in the `examples/` folder
- Try implementing classic algorithms (binary search, merge sort, etc.)
- Create your own games and challenges
- Combine algorithms to solve complex problems

**Keep experimenting and building! 🎉**
