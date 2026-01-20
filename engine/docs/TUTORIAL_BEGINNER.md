# 🎓 KidLang Beginner Tutorial

Welcome to KidLang! This tutorial will teach you programming step-by-step using fun examples.

## 📚 Table of Contents

1. [Your First Program](#your-first-program)
2. [Using Boxes (Variables)](#using-boxes-variables)
3. [Talking to Your Program](#talking-to-your-program)
4. [Making Decisions](#making-decisions)
5. [Repeating Actions (Loops)](#repeating-actions-loops)
6. [Working with Lists (Stacks)](#working-with-lists-stacks)

---

## Your First Program

Let's start with the simplest program - saying hello!

```kidlang
// This is a comment - the computer ignores it
// Comments help explain what your code does

print Hello, World!
print Welcome to programming!
print This is fun!
```

**What you learned:**
- Use `//` to write comments that help you remember what your code does
- Use `print` to display messages on the screen
- You don't need quotes around simple text (but you can use them!)

**Try it yourself:**
```kidlang
print Hello!
print My name is [your name]
print I am learning to code!
```

---

## Using Boxes (Variables)

Think of a box as a container that holds a value. You can put numbers or words inside!

### Simple Math with Boxes

```kidlang
// Create boxes to store numbers
box age = 10
box next_year = box age + 1

print I am box age years old
print Next year I will be box next_year
```

**What you learned:**
- `box name = value` creates a box and puts a value in it
- Use `box name` to get the value from a box
- You can do math: `+` (add), `-` (subtract), `*` (multiply), `/` (divide)

### Fun Math Challenge

```kidlang
print === My Calculator ===

box a = 15
box b = 7

box sum = box a + box b
box difference = box a - box b
box product = box a * box b

print box a + box b = box sum
print box a - box b = box difference
print box a × box b = box product
```

**Try it yourself:** Create a calculator that works with different numbers!

---

## Talking to Your Program

Programs are more fun when they can talk to you!

### Getting User Input

```kidlang
// Ask the user a question
ask What is your name?

// The answer is stored in a special box called "answer"
box name = answer

print Hello, box name!
print Nice to meet you!
```

**What you learned:**
- `ask` shows a question and waits for the user to type something
- The answer goes into a special box called `answer`
- Copy `answer` to your own box to save it

### Age Calculator

```kidlang
print === Age Calculator ===

ask What is your name?
box name = answer

ask How old are you?
box age = answer

box age_next = box age + 1
box age_5years = box age + 5

print Hello box name!
print You are box age years old
print Next year you will be box age_next
print In 5 years you will be box age_5years
```

**Try it yourself:** Make a program that asks for two numbers and shows their sum!

---

## Making Decisions

Programs can make choices using `if` statements!

### Simple Choices

```kidlang
ask How old are you?
box age = answer

if box age >= 13 then
    print You are a teenager!
end

if box age < 13 then
    print You are a kid!
end

print Thanks for telling me!
```

**What you learned:**
- `if condition then` checks if something is true
- Put the code to run between `then` and `end`
- Comparison operators: `=` (equal), `>` (greater), `<` (less), `>=` (greater or equal), `<=` (less or equal), `!=` (not equal)

### Number Guessing Game

```kidlang
print === Guess the Number! ===

box secret = 7

ask Guess a number between 1 and 10:
box guess = answer

if box guess = box secret then
    print 🎉 You won! That's correct!
end

if box guess > box secret then
    print Too high! Try again.
end

if box guess < box secret then
    print Too low! Try again.
end
```

### Using AND and OR

You can combine conditions!

```kidlang
ask What is your age?
box age = answer

ask Do you have a ticket? (yes/no)
box ticket = answer

// Both conditions must be true
if box age >= 18 and box ticket = yes then
    print You can enter the concert!
end

// At least one condition must be true
if box age < 5 or box age > 65 then
    print You get a discount!
end
```

**Try it yourself:** Make a game that checks if a password is correct!

---

## Repeating Actions (Loops)

Sometimes you want to do something multiple times. Use labels and `goto`!

### Countdown Timer

```kidlang
print === Countdown ===

box count = 5

loop:
// Check if we're done
if box count < 0 then
    goto done
end

print box count
sleep 1000
box count = box count - 1
goto loop

done:
print BLAST OFF! 🚀
```

**What you learned:**
- `label:` creates a marker in your code
- `goto label` jumps to that marker
- `sleep 1000` pauses for 1000 milliseconds (1 second)
- You can create loops by jumping back to a label

### Multiplication Table

```kidlang
print === Multiplication Table ===

ask Which table do you want? (1-10):
box table = answer

box i = 1

loop:
if box i > 10 then
    goto done
end

box result = box table * box i
print box table × box i = box result

box i = box i + 1
goto loop

done:
print All done!
```

### Counting Game

```kidlang
print Let's count together!

box num = 1

count:
if box num > 5 then
    goto stop
end

print Number: box num
box num = box num + 1
goto count

stop:
print We counted to 5!
```

**Try it yourself:** Make a program that prints your name 10 times!

---

## Working with Lists (Stacks)

Stacks let you store many values in one place!

### Simple Stack

```kidlang
print === My Favorite Things ===

// Create a stack (list)
stack favorites

// Add items to the stack
stack favorites[1] = pizza
stack favorites[2] = soccer
stack favorites[3] = coding

// Display items
print My favorite food is stack favorites[1]
print My favorite sport is stack favorites[2]
print My favorite hobby is stack favorites[3]
```

**What you learned:**
- `stack name` creates a list
- Use `[1]`, `[2]`, `[3]` to store different values
- You can also use `(1)` instead of `[1]` - both work!

### Inventory System

```kidlang
print === Game Inventory ===

stack inventory

// Add items with quantities
stack inventory[swords] = 2
stack inventory[potions] = 5
stack inventory[gold] = 100

print Your Inventory:
print Swords: stack inventory[swords]
print Potions: stack inventory[potions]
print Gold: stack inventory[gold]

// Use an item
stack inventory[potions] = stack inventory[potions] - 1
print Used a potion! Potions left: stack inventory[potions]
```

### Score Tracker

```kidlang
print === Test Scores ===

stack scores

// Record scores
stack scores[math] = 95
stack scores[science] = 87
stack scores[english] = 92

print Test Results:
print Math: stack scores[math]
print Science: stack scores[science]
print English: stack scores[english]

// Calculate average
box total = stack scores[math] + stack scores[science] + stack scores[english]
box average = box total / 3
print Average Score: box average
```

**Try it yourself:** Make a shopping list using a stack!

---

## 🎯 Practice Projects

Now that you've learned the basics, try these projects:

### Project 1: Temperature Converter
Create a program that converts Celsius to Fahrenheit.
Formula: F = (C × 9/5) + 32

### Project 2: Quiz Game
Make a quiz with 3 questions and keep track of the score.

### Project 3: Story Generator
Ask the user for words (noun, verb, adjective) and create a funny story.

### Project 4: Simple Calculator
Ask for two numbers and an operation (+, -, *, /) and show the result.

---

## 🚀 What's Next?

You've learned the fundamentals! Check out:
- **TUTORIAL_INTERMEDIATE.md** - Learn about files, advanced math, and more!
- **TUTORIAL_ALGORITHMS.md** - Learn to build real algorithms!
- **examples/** folder - See 20+ complete programs!

**Keep coding and have fun! 🎉**
