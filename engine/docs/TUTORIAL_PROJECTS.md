# 🎓 KidLang Project Tutorial

Learn by building! This tutorial walks you through complete projects step-by-step.

## 📚 Projects in This Tutorial

1. [Number Guessing Game](#project-1-number-guessing-game)
2. [Calculator Program](#project-2-calculator-program)
3. [Story Generator](#project-3-story-generator)
4. [Contact Manager](#project-4-contact-manager)
5. [Math Quiz Game](#project-5-math-quiz-game)
6. [Simple RPG Battle](#project-6-simple-rpg-battle)
7. [Countdown Timer](#project-7-countdown-timer)
8. [Temperature Converter](#project-8-temperature-converter)

---

## Project 1: Number Guessing Game

**Goal:** Create a game where the computer picks a random number and you try to guess it.

### What You'll Learn
- Using random numbers
- Loops with labels and goto
- Conditional logic
- User input validation

### Step-by-Step Build

**Step 1: Basic structure**
```kidlang
print === Number Guessing Game ===
box secret = random % 10 + 1
print I'm thinking of a number between 1 and 10
```

**Step 2: Get user guess**
```kidlang
ask Enter your guess:
box guess = answer
```

**Step 3: Check if correct**
```kidlang
if box guess = box secret then
    print Correct! You won!
end

if box guess != box secret then
    print Wrong! Try again.
end
```

**Step 4: Add loop to play multiple times**
```kidlang
game_loop:
ask Enter your guess:
box guess = answer

if box guess = box secret then
    print Correct!
    goto end_game
end

print Wrong! Try again.
goto game_loop

end_game:
print Thanks for playing!
```

**Step 5: Add hints (too high/low)**
```kidlang
if box guess > box secret then
    print Too high!
end

if box guess < box secret then
    print Too low!
end
```

**Step 6: Add try counter**
```kidlang
box tries = 0
box max_tries = 5

game_loop:
box tries = box tries + 1

if box tries > box max_tries then
    print Game Over! The number was box secret
    goto end_game
end

print Try box tries of box max_tries
// rest of game logic...
```

### Complete Program

```kidlang
print ================================
print   NUMBER GUESSING GAME
print ================================

box secret = random % 50 + 1
box tries = 0
box max_tries = 7

print I'm thinking of a number between 1 and 50
print You have box max_tries tries!
print

game_loop:
box tries = box tries + 1
print Try box tries of box max_tries

ask Enter your guess:
box guess = answer

if box guess = box secret then
    print
    print *** WINNER! ***
    print You guessed it in box tries tries!
    goto end_game
end

if box guess > box secret then
    print Too high! Try lower.
end

if box guess < box secret then
    print Too low! Try higher.
end

if box tries >= box max_tries then
    print
    print Game Over!
    print The number was box secret
    goto end_game
end

print
goto game_loop

end_game:
print Thanks for playing!
```

### Enhancements to Try
- Add difficulty levels (easy: 1-10, hard: 1-100)
- Track best score (fewest guesses)
- Add hints every 3 wrong guesses
- Let player choose number range

---

## Project 2: Calculator Program

**Goal:** Build a calculator that performs basic math operations.

### What You'll Learn
- Menu systems
- Mathematical operations
- Input validation
- Program flow control

### Step-by-Step Build

**Step 1: Create menu**
```kidlang
print === Calculator ===
print 1. Add
print 2. Subtract
print 3. Multiply
print 4. Divide
ask Choose operation:
```

**Step 2: Get numbers**
```kidlang
ask Enter first number:
box num1 = answer

ask Enter second number:
box num2 = answer
```

**Step 3: Perform operation**
```kidlang
if answer = 1 then
    box result = box num1 + box num2
    print Result: box result
end
```

### Complete Program

```kidlang
print ================================
print      KIDLANG CALCULATOR
print ================================

calculator:
print
print === MENU ===
print 1. Add
print 2. Subtract
print 3. Multiply
print 4. Divide
print 5. Square Root
print 6. Exit
print
ask Choose operation:
box choice = answer

if box choice = 6 then goto exit

if box choice = 5 then goto sqrt_op

// Operations needing two numbers
ask Enter first number:
box num1 = answer

ask Enter second number:
box num2 = answer

if box choice = 1 then
    box result = box num1 + box num2
    print box num1 + box num2 = box result
end

if box choice = 2 then
    box result = box num1 - box num2
    print box num1 - box num2 = box result
end

if box choice = 3 then
    box result = box num1 * box num2
    print box num1 × box num2 = box result
end

if box choice = 4 then
    if box num2 = 0 then
        print Error: Cannot divide by zero!
        goto calculator
    end
    box result = box num1 / box num2
    print box num1 ÷ box num2 = box result
end

goto calculator

sqrt_op:
ask Enter number:
box num = answer
print √box num = sqrt box num
goto calculator

exit:
print Thanks for calculating!
```

### Enhancements to Try
- Add modulo (remainder) operation
- Add power/exponent function
- Keep history of calculations
- Add memory (store/recall value)

---

## Project 3: Story Generator

**Goal:** Create random funny stories by asking for words.

### What You'll Learn
- String concatenation
- Creative programming
- User interaction
- Text formatting

### Complete Program

```kidlang
print ================================
print    STORY GENERATOR
print ================================
print Fill in the blanks to create a funny story!
print

ask Enter a person's name:
box name = answer

ask Enter an animal:
box animal = answer

ask Enter a color:
box color = answer

ask Enter a place:
box place = answer

ask Enter a food:
box food = answer

ask Enter a number:
box number = answer

ask Enter an adjective:
box adjective = answer

ask Enter a verb (ending in -ing):
box verb = answer

print
print ================================
print      YOUR STORY
print ================================
print
print Once upon a time, box name went to box place.
print There, they met a box adjective box color box animal
print who was box verb with box number pieces of box food!
print
print box name said, "What a box adjective day!"
print The box animal agreed and they became best friends.
print
print THE END
print ================================

ask Generate another story? (yes/no):
if answer = yes then goto start
```

### Enhancements to Try
- Add multiple story templates
- Save stories to a file
- Create longer stories with more blanks
- Add story categories (adventure, comedy, scary)

---

## Project 4: Contact Manager

**Goal:** Build a program to save and manage contacts.

### What You'll Learn
- File operations (open, write, read, close)
- Data persistence
- CRUD operations (Create, Read, Update, Delete)
- Menu-driven programs

### Complete Program

```kidlang
print ================================
print     CONTACT MANAGER
print ================================

file contacts
open contacts contacts.txt

box contact_count = 0

menu:
print
print === MENU ===
print 1. Add contact
print 2. View all contacts
print 3. Search contact
print 4. Exit
print
print Total contacts: box contact_count
ask Choose:

if answer = 1 then goto add
if answer = 2 then goto view
if answer = 3 then goto search
if answer = 4 then goto exit
goto menu

add:
print
ask Contact name:
box name = answer

ask Phone number:
box phone = answer

ask Email:
box email = answer

// Save as: name|phone|email
box record = box name + | + box phone + | + box email
write contacts box record
box contact_count = box contact_count + 1

print Contact saved!
sleep 1000
goto menu

view:
print
print === ALL CONTACTS ===
if box contact_count = 0 then
    print No contacts yet!
    sleep 1000
    goto menu
end

print
exec cat contacts.txt
print
print === End of list ===
ask Press Enter to continue...
goto menu

search:
print
ask Enter name to search:
box search_name = answer

print Searching for: box search_name
print
exec grep box search_name contacts.txt
print
ask Press Enter to continue...
goto menu

exit:
close contacts
print Contacts saved!
print Goodbye!
```

### Enhancements to Try
- Add delete contact feature
- Add edit contact feature
- Sort contacts alphabetically
- Add birthday field
- Export contacts to CSV

---

## Project 5: Math Quiz Game

**Goal:** Create an interactive math practice game.

### What You'll Learn
- Random number generation
- Score tracking
- Timed challenges (with sleep)
- Educational game design

### Complete Program

```kidlang
print ================================
print      MATH QUIZ GAME
print ================================

box score = 0
box questions = 5
box current = 1

quiz:
if box current > box questions then goto results

print
print Question box current of box questions

// Generate random numbers
box a = random % 20 + 1
box b = random % 20 + 1

// Random operation
box op = random % 4

if box op = 0 then goto add_q
if box op = 1 then goto sub_q
if box op = 2 then goto mul_q
if box op = 3 then goto div_q

add_q:
box correct = box a + box b
print What is box a + box b?
goto get_answer

sub_q:
// Make sure a > b for positive result
if box a < box b then
    box temp = box a
    box a = box b
    box b = box temp
end
box correct = box a - box b
print What is box a - box b?
goto get_answer

mul_q:
box correct = box a * box b
print What is box a × box b?
goto get_answer

div_q:
// Make sure it divides evenly
box a = box a * box b
box correct = box a / box b
print What is box a ÷ box b?
goto get_answer

get_answer:
ask Your answer:
box ans = answer

if box ans = box correct then
    print ✓ Correct!
    box score = box score + 1
end

if box ans != box correct then
    print ✗ Wrong! Answer was box correct
end

sleep 1500
box current = box current + 1
goto quiz

results:
print
print ================================
print       QUIZ COMPLETE!
print ================================
print Your score: box score out of box questions

box percent = box score * 100 / box questions

if box percent = 100 then
    print PERFECT SCORE! 🌟
end

if box percent >= 80 and box percent < 100 then
    print Excellent! 🎉
end

if box percent >= 60 and box percent < 80 then
    print Good job! 👍
end

if box percent < 60 then
    print Keep practicing! 📚
end

print
ask Play again? (yes/no):
if answer = yes then
    box score = 0
    box current = 1
    goto quiz
end

print Thanks for playing!
```

### Enhancements to Try
- Add timer for each question
- Track best score
- Add difficulty levels
- Add more operation types
- Save high scores to file

---

## Project 6: Simple RPG Battle

**Goal:** Create a turn-based battle game.

### What You'll Learn
- Game state management
- Turn-based logic
- Random events
- Health/stat tracking

### Complete Program

```kidlang
print ================================
print      RPG BATTLE GAME
print ================================

// Player stats
box player_health = 100
box player_attack = 15

// Enemy stats
box enemy_health = 80
box enemy_attack = 10

print You encounter a wild monster!
print

battle:
// Check win/lose conditions
if box player_health <= 0 then goto player_loses
if box enemy_health <= 0 then goto player_wins

// Display status
print === STATUS ===
print Your Health: box player_health
print Enemy Health: box enemy_health
print

// Player turn
print Your turn!
print 1. Attack
print 2. Defend
print 3. Use Potion
ask Choose action:

if answer = 1 then goto player_attack
if answer = 2 then goto player_defend
if answer = 3 then goto use_potion
goto battle

player_attack:
box damage = box player_attack + random % 10
box enemy_health = box enemy_health - box damage
print You attack for box damage damage!
sleep 1000
goto enemy_turn

player_defend:
print You brace for impact!
box defend_bonus = 5
sleep 1000
goto enemy_turn

use_potion:
box heal = 30
box player_health = box player_health + box heal
print You heal for box heal HP!
sleep 1000
goto enemy_turn

enemy_turn:
if box enemy_health <= 0 then goto player_wins

print
print Enemy's turn!
sleep 1000

box damage = box enemy_attack + random % 5

// Apply defend bonus
if box defend_bonus > 0 then
    box damage = box damage - box defend_bonus
    box defend_bonus = 0
end

if box damage < 0 then
    box damage = 0
end

box player_health = box player_health - box damage
print Enemy attacks for box damage damage!
sleep 1500
goto battle

player_wins:
print
print ================================
print      VICTORY!
print ================================
print You defeated the monster!
print Remaining health: box player_health
goto end

player_loses:
print
print ================================
print      DEFEAT!
print ================================
print You were defeated...
print Better luck next time!
goto end

end:
print
ask Play again? (yes/no):
if answer = yes then
    box player_health = 100
    box enemy_health = 80
    box defend_bonus = 0
    goto battle
end

print Thanks for playing!
```

### Enhancements to Try
- Add multiple enemy types
- Add experience and leveling up
- Add inventory system
- Add special moves/magic
- Add boss battles

---

## Project 7: Countdown Timer

**Goal:** Create a customizable countdown timer.

### What You'll Learn
- Time delays with sleep
- Loop control
- User customization
- Event triggers

### Complete Program

```kidlang
print ================================
print     COUNTDOWN TIMER
print ================================

ask Enter countdown time (seconds):
box seconds = answer

if box seconds <= 0 then
    print Invalid time!
    goto end
end

print
print Starting countdown from box seconds seconds...
print

countdown:
if box seconds <= 0 then goto timer_done

print box seconds...
sleep 1000
box seconds = box seconds - 1
goto countdown

timer_done:
print
print ================================
print      TIME'S UP!
print ================================
print
print 🔔 DING DING DING! 🔔

end:
```

### Enhancements to Try
- Add pause/resume functionality
- Add sound alerts (using exec)
- Create interval timer (work/break cycles)
- Save timer presets
- Add countdown for specific event/date

---

## Project 8: Temperature Converter

**Goal:** Convert between Celsius, Fahrenheit, and Kelvin.

### What You'll Learn
- Formula implementation
- Decimal math
- Menu systems
- Unit conversion

### Complete Program

```kidlang
print ================================
print   TEMPERATURE CONVERTER
print ================================

menu:
print
print === MENU ===
print 1. Celsius to Fahrenheit
print 2. Fahrenheit to Celsius
print 3. Celsius to Kelvin
print 4. Kelvin to Celsius
print 5. Exit
ask Choose conversion:

if answer = 1 then goto c_to_f
if answer = 2 then goto f_to_c
if answer = 3 then goto c_to_k
if answer = 4 then goto k_to_c
if answer = 5 then goto exit
goto menu

c_to_f:
ask Enter temperature in Celsius:
box c = answer
// F = C * 9/5 + 32
box f = box c * 9 / 5 + 32
print box c°C = box f°F
sleep 2000
goto menu

f_to_c:
ask Enter temperature in Fahrenheit:
box f = answer
// C = (F - 32) * 5/9
box temp = box f - 32
box c = box temp * 5 / 9
print box f°F = box c°C
sleep 2000
goto menu

c_to_k:
ask Enter temperature in Celsius:
box c = answer
// K = C + 273.15
box k = box c + 273
print box c°C = box k K
sleep 2000
goto menu

k_to_c:
ask Enter temperature in Kelvin:
box k = answer
// C = K - 273.15
box c = box k - 273
print box k K = box c°C
sleep 2000
goto menu

exit:
print Thanks for using the converter!
```

### Enhancements to Try
- Add Rankine scale
- Add temperature descriptions (hot/cold/freezing)
- Save conversion history
- Add batch conversion mode
- Add weather-related advice based on temperature

---

## 🎯 General Programming Tips

### Planning Your Project

1. **Define the goal** - What should it do?
2. **List features** - What capabilities does it need?
3. **Break it down** - What are the steps?
4. **Start simple** - Build the core first
5. **Test often** - Run it frequently
6. **Add features** - Improve gradually

### Debugging Your Code

```kidlang
// Add debug prints
print DEBUG: Starting calculation
box result = box a + box b
print DEBUG: Result is box result

// Test with known values
box test = 5 + 3
print Should be 8: box test

// Mark sections
print === SECTION 1 START ===
// your code
print === SECTION 1 END ===
```

### Making Programs User-Friendly

```kidlang
// Clear messages
print Welcome to My Program!
print Please enter your name below:

// Validate input
ask Enter age:
box age = answer
if box age < 0 then
    print Error: Age must be positive
    goto ask_again
end

// Provide feedback
print Processing...
// do work
print Done!
```

---

## 🚀 Challenge: Build Your Own Project!

Now that you've learned from these projects, create your own! Here are ideas:

1. **Hangman Game** - Word guessing game
2. **Virtual Pet** - Care for a digital pet
3. **Weather Logger** - Track daily weather
4. **Budget Tracker** - Manage money
5. **Flashcard App** - Study helper
6. **Dice Roller** - For board games
7. **Password Generator** - Create secure passwords
8. **Unit Converter** - Convert various units
9. **Grocery List** - Shopping helper
10. **Journal App** - Daily diary

**Remember:**
- Start with a simple version
- Test it thoroughly
- Add features one at a time
- Ask for feedback
- Have fun coding!

---

**You're now ready to build amazing projects! 🎉**
