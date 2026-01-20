# 🎮 Welcome to KidLang! Your Coding Adventure Starts Here!

**Hey there, future programmer!** 👋

Are you ready to learn how to code? KidLang is a special programming language made just for kids like you (ages 8-13)! It's super fun and easy to learn. Let's start your coding journey!

---

## 🌟 What is Programming?

Programming is like giving instructions to a computer. Just like you follow a recipe to bake cookies, a computer follows your code to do cool things!

---

## 🎯 Your First Program - Say Hello!

Let's write your very first program! Type this:

```kidlang
print Hello, World!
print I am learning to code!
```

**What happens?** The computer will print (show) your message on the screen! 🎉

> **Fun Fact:** You don't even need to type `print` for simple messages. Just write text and it works! Try: `Hello there!`

---

## 📦 Using Boxes (Variables)

Think of a **box** like a container where you can store things. You can put a number, a word, or anything you want in a box!

### Storing a Number

```kidlang
box age = 10
print I am box age years old
```

### Doing Math with Boxes

```kidlang
box apples = 5
box oranges = 3
box total = box apples + box oranges
print I have box total fruits!
```

**Cool Things You Can Do:**
- **Add:** `box a = 10 + 5` → Result: 15
- **Subtract:** `box b = 20 - 8` → Result: 12
- **Multiply:** `box c = 4 * 3` → Result: 12
- **Divide:** `box d = 15 / 3` → Result: 5

---

## 💬 Talking to Your Program (Getting Input)

Want your program to ask you questions? Use **ask**!

```kidlang
ask What is your name?
print Hello box answer
print Nice to meet you!
```

**What happens?** 
1. The program asks for your name
2. You type your name
3. Your answer gets stored in a special box called `answer`
4. The program says hello to you!

### Try This: Age Calculator

```kidlang
ask How old are you?
box my_age = answer
box next_year = box my_age + 1
print Next year you will be box next_year years old!
```

---

## 🤔 Making Decisions (If Statements)

Sometimes you want your program to make choices. Use **if/then/end**!

```kidlang
box score = 85

if box score > 80 then
print Great job! You got an A!
end
```

### Guess the Number Game

```kidlang
ask Guess a number between 1 and 10:
box guess = answer
box secret = 7

if box guess = box secret then
print You won! The secret was box secret
end

if box guess != box secret then
print Sorry! Try again
end
```

**Comparison Symbols:**
- `=` means "is equal to"
- `!=` means "is NOT equal to"
- `>` means "greater than"
- `<` means "less than"
- `>=` means "greater than or equal to"
- `<=` means "less than or equal to"

---

## 🔄 Repeating Things (Loops with Labels)

Want to do something over and over? Use **labels** and **goto**!

```kidlang
box counter = 1

start:
print Counting: box counter
box counter = box counter + 1

if box counter < 6 then goto start

print Done counting!
```

**What happens?** This counts from 1 to 5!

### Countdown Timer

```kidlang
box time = 10

countdown:
print box time
sleep 1
box time = box time - 1

if box time > 0 then goto countdown

print Blast off!
```

---

## 🎲 Fun with Math Functions

KidLang has special math powers!

```kidlang
// Square root (what number times itself equals this?)
box a = sqrt 16
print box a
// Result: 4 (because 4 × 4 = 16)

// Square (multiply a number by itself)
box b = sqr 5
print box b
// Result: 25 (because 5 × 5 = 25)

// Absolute value (remove the minus sign)
box c = abs -10
print box c
// Result: 10

// Random number between 0 and 1
box d = random
print You got: box d
```

---

## 📝 Working with Words (Strings)

You can do cool things with words too!

### Joining Words Together

```kidlang
box first = Hello
box second = World
box together = box first + box second
print box together
// Result: HelloWorld
```

### Repeating Words

```kidlang
box laugh = Ha * 5
print box laugh
// Result: HaHaHaHaHa
```

### Getting One Letter

```kidlang
box word = Pizza
box letter = box word / 1
print box letter
// Result: P (the first letter!)
```

---

## 📚 Using Lists (Stacks)

A **stack** is like a box that can hold lots of things, each with a number or name!

```kidlang
stack toys
stack toys[1] = Robot
stack toys[2] = Ball
stack toys[3] = Puzzle

print My first toy is: stack toys[1]
print My second toy is: stack toys[2]
print My third toy is: stack toys[3]
```

You can also use words as labels:

```kidlang
stack friend
stack friend[name] = Alex
stack friend[age] = 10
stack friend[hobby] = Soccer

print Name: stack friend[name]
print Age: stack friend[age]
print Hobby: stack friend[hobby]
```

---

## 🎮 Mini Project: Multiplication Quiz

Let's combine what you learned into a fun quiz game!

```kidlang
print === MULTIPLICATION QUIZ ===

ask What is 7 times 8?
box answer1 = answer

if box answer1 = 56 then
print Correct! Great job!
end

if box answer1 != 56 then
print Not quite! The answer is 56
end

ask What is 9 times 6?
box answer2 = answer

if box answer2 = 54 then
print Perfect! You're a math star!
end

if box answer2 != 54 then
print The correct answer is 54
end

print Thanks for playing!
```

---

## 🎨 Mini Project: Story Creator

```kidlang
print Let's create a funny story!

ask What is your favorite animal?
box animal = answer

ask What is your favorite food?
box food = answer

ask What is your favorite color?
box color = answer

print ================
print YOUR STORY:
print ================
print Once upon a time, there was a box color box animal
print This box animal loved to eat box food every day!
print One day, the box animal found a magical box food
print And they lived happily ever after!
print ================
```

---

## 🏆 Challenge Projects for You!

Now that you know the basics, try building these fun projects:

### 1. 🎯 Simple Calculator
Make a program that asks for two numbers and adds them together!

### 2. 🌡️ Temperature Converter
Convert temperatures from Fahrenheit to Celsius!
(Hint: Celsius = (Fahrenheit - 32) × 5 / 9)

### 3. 🎲 Dice Roller
Use `random` to simulate rolling dice!

### 4. 📊 Grade Calculator
Ask for test scores and calculate the average!

### 5. 🎪 Carnival Game
Create a number guessing game with multiple chances!

---

## 💡 Tips for Young Programmers

1. **Don't worry about mistakes!** Everyone makes them. Just fix them and learn!
2. **Experiment!** Try changing numbers and words to see what happens
3. **Start small!** Make simple programs first, then add more features
4. **Have fun!** Programming should be enjoyable, like solving puzzles
5. **Save your work!** Give your programs names like `mygame.kid`

---

## 🎓 What You've Learned!

✅ How to print messages  
✅ How to use boxes (variables) to store things  
✅ How to do math (+, -, ×, ÷)  
✅ How to ask questions and get answers  
✅ How to make decisions with if/then  
✅ How to repeat things with labels and goto  
✅ How to use lists (stacks)  
✅ How to build fun projects!  

---

## 🚀 What's Next?

Want to learn more? Check out these files:

- **TUTORIAL_BEGINNER.md** - More beginner lessons
- **TUTORIAL_ALGORITHMS.md** - Learn to sort and search
- **TUTORIAL_PROJECTS.md** - Build complete projects
- **examples/** folder - See 20+ example programs!

---

## 🎉 You're a Programmer Now!

Congratulations! You've learned how to code in KidLang! Keep practicing, keep creating, and most importantly—**have fun coding!** 🌟

Remember: Every expert programmer started exactly where you are now. You're doing amazing! 💪

---

**Happy Coding! 🎮✨**
