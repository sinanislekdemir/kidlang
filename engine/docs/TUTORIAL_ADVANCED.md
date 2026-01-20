# 🎓 KidLang Advanced Tutorial

Learn advanced techniques, file handling, and build complex projects!

## 📚 Table of Contents

1. [Advanced File Operations](#advanced-file-operations)
2. [Building Real Applications](#building-real-applications)
3. [Data Structures](#data-structures)
4. [Advanced Math](#advanced-math)
5. [String Manipulation](#string-manipulation)
6. [Best Practices](#best-practices)

---

## Advanced File Operations

### Reading and Writing Data Files

Create a persistent data system that survives program restarts.

```kidlang
print === Contact Book ===

file contacts
open contacts contacts.txt

menu:
print
print 1. Add contact
print 2. View contacts
print 3. Exit
ask Choose:

if answer = 1 then goto add
if answer = 2 then goto view
if answer = 3 then goto exit

add:
ask Enter name:
box name = answer
ask Enter phone:
box phone = answer

box entry = box name + : + box phone
write contacts box entry
print Contact saved!
goto menu

view:
print === Contacts ===
exec cat contacts.txt
print === End ===
sleep 1000
goto menu

exit:
close contacts
print Goodbye!
```

### Processing Files Line by Line

Read a file and process each line separately.

```kidlang
print === Word Counter ===

file doc
open doc story.txt

box total_lines = 0
box total_words = 0

read_loop:
readline doc box line

// Check if we reached end of file (empty line)
if box line =  then
    goto done
end

box total_lines = box total_lines + 1
// Estimate words (simplified - count spaces + 1)
box total_words = box total_words + 5

goto read_loop

done:
close doc
print Total lines: box total_lines
print Estimated words: box total_words
```

### Creating a Simple Database

Store structured data and search it.

```kidlang
print === Student Database ===

file db
open db students.txt

main_menu:
print
print === MENU ===
print 1. Add student
print 2. List all
print 3. Exit
ask Choose:

if answer = 1 then goto add_student
if answer = 2 then goto list_all
if answer = 3 then goto exit

add_student:
ask Student name:
box name = answer
ask Math grade:
box math = answer
ask Science grade:
box science = answer
ask English grade:
box english = answer

// Create data record: name|math|science|english
box record = box name + | + box math + | + box science + | + box english
write db box record
print Student added!
goto main_menu

list_all:
print
print === All Students ===
exec cat students.txt
print === End ===
sleep 2000
goto main_menu

exit:
close db
print Database closed
```

---

## Building Real Applications

### Complete To-Do List Manager

A full-featured task management system.

```kidlang
print ===================================
print      TASK MANAGER PRO
print ===================================

file tasks
open tasks tasks.txt

stack task_list
box task_count = 0

// Load existing tasks
load:
readline tasks box task_line
if box task_line !=  then
    box task_count = box task_count + 1
    stack task_list[task_count] = box task_line
    goto load
end

close tasks

// Main program
main_menu:
print
print === MENU ===
print 1. Add task
print 2. View tasks
print 3. Mark task complete
print 4. Delete task
print 5. Save and exit
print
print Total tasks: box task_count
ask Choose:

if answer = 1 then goto add_task
if answer = 2 then goto view_tasks
if answer = 3 then goto mark_complete
if answer = 4 then goto delete_task
if answer = 5 then goto save_exit
goto main_menu

add_task:
ask Task description:
box new_task = answer
box task_count = box task_count + 1
stack task_list[task_count] = [ ] box new_task
print Task added!
sleep 500
goto main_menu

view_tasks:
print
print === YOUR TASKS ===
if box task_count = 0 then
    print No tasks yet!
    sleep 1000
    goto main_menu
end

box i = 1
view_loop:
if box i > box task_count then goto view_done
print box i. stack task_list[i]
box i = box i + 1
goto view_loop

view_done:
print === End ===
ask Press Enter...
goto main_menu

mark_complete:
if box task_count = 0 then
    print No tasks to complete!
    sleep 1000
    goto main_menu
end

ask Which task number?
box task_num = answer

if box task_num < 1 then goto main_menu
if box task_num > box task_count then goto main_menu

box old_task = stack task_list[task_num]
box new_task = [X] + box old_task
stack task_list[task_num] = box new_task
print Task marked complete!
sleep 500
goto main_menu

delete_task:
ask Which task to delete?
box del_num = answer

if box del_num < 1 then goto main_menu
if box del_num > box task_count then goto main_menu

// Shift remaining tasks down
box i = box del_num
shift_loop:
box next = box i + 1
if box next > box task_count then goto shift_done
stack task_list[i] = stack task_list[next]
box i = box i + 1
goto shift_loop

shift_done:
box task_count = box task_count - 1
print Task deleted!
sleep 500
goto main_menu

save_exit:
// Save all tasks to file
close tasks
exec rm tasks.txt
open tasks tasks.txt

box i = 1
save_loop:
if box i > box task_count then goto save_done
write tasks stack task_list[i]
box i = box i + 1
goto save_loop

save_done:
close tasks
print Tasks saved!
print Goodbye!
```

### Simple Text Adventure Game

Create an interactive story game.

```kidlang
print ====================================
print     THE MYSTERIOUS CAVE
print ====================================
print

box health = 100
box has_torch = 0
box has_key = 0

entrance:
print You stand at the entrance of a dark cave.
print Your health: box health
print
print 1. Enter the cave
print 2. Look around
print 3. Go home
ask What do you do?

if answer = 1 then goto cave_entrance
if answer = 2 then goto look_entrance
if answer = 3 then goto go_home
goto entrance

look_entrance:
print You see a torch leaning against a rock.
ask Take the torch? (yes/no)
if answer = yes then
    box has_torch = 1
    print You picked up the torch!
end
sleep 1000
goto entrance

cave_entrance:
print You enter the dark cave...
sleep 1000

if box has_torch = 1 then
    print Your torch lights the way!
    goto cave_lit
end

print It's too dark! You stumble and fall.
box health = box health - 20
print Lost 20 health!
sleep 1000
goto entrance

cave_lit:
print You see two passages:
print 1. Left passage (sounds of water)
print 2. Right passage (warm breeze)
print 3. Go back
ask Choose:

if answer = 1 then goto left_passage
if answer = 2 then goto right_passage
if answer = 3 then goto entrance
goto cave_lit

left_passage:
print You follow the sound of water...
sleep 1000
print You find an underground lake!
print Something glitters in the water...
ask Search the water? (yes/no)

if answer = yes then
    print You found a golden key!
    box has_key = 1
    sleep 1000
end
goto cave_lit

right_passage:
print You walk down the warm passage...
sleep 1000
print You find a locked door!

if box has_key = 1 then
    print You use the golden key!
    goto treasure_room
end

print You need a key to open it.
sleep 1000
goto cave_lit

treasure_room:
print
print ====================================
print   YOU FOUND THE TREASURE!
print ====================================
print
print Congratulations! You completed the adventure!
goto end

go_home:
print You decide to go home.
print Maybe another day...
goto end

end:
print
print === GAME OVER ===
print Final health: box health
```

### Quiz Game with Score Tracking

```kidlang
print ================================
print    KIDLANG QUIZ GAME
print ================================

box score = 0
box total = 5

// Question 1
print
print Question 1: What is 7 + 8?
print A) 14
print B) 15
print C) 16
ask Your answer:

if answer = B then
    print Correct!
    box score = box score + 1
end

if answer != B then
    print Wrong! Answer is B) 15
end
sleep 1500

// Question 2
print
print Question 2: What is the capital of France?
print A) London
print B) Paris
print C) Berlin
ask Your answer:

if answer = B then
    print Correct!
    box score = box score + 1
end

if answer != B then
    print Wrong! Answer is B) Paris
end
sleep 1500

// Question 3
print
print Question 3: How many days in a week?
print A) 5
print B) 6
print C) 7
ask Your answer:

if answer = C then
    print Correct!
    box score = box score + 1
end

if answer != C then
    print Wrong! Answer is C) 7
end
sleep 1500

// Question 4
print
print Question 4: What is 12 * 3?
print A) 36
print B) 35
print C) 34
ask Your answer:

if answer = A then
    print Correct!
    box score = box score + 1
end

if answer != A then
    print Wrong! Answer is A) 36
end
sleep 1500

// Question 5
print
print Question 5: What color is the sky?
print A) Green
print B) Blue
print C) Red
ask Your answer:

if answer = B then
    print Correct!
    box score = box score + 1
end

if answer != B then
    print Wrong! Answer is B) Blue
end
sleep 1500

// Results
print
print ================================
print       QUIZ COMPLETE!
print ================================
print You scored: box score out of box total

box percent = box score * 100 / box total

if box percent >= 80 then
    print Grade: A - Excellent!
end

if box percent >= 60 and box percent < 80 then
    print Grade: B - Good job!
end

if box percent >= 40 and box percent < 60 then
    print Grade: C - Keep practicing!
end

if box percent < 40 then
    print Grade: F - Study more!
end
```

---

## Data Structures

### Building a Dictionary

```kidlang
print === English-Spanish Dictionary ===

stack dictionary

// Add translations
stack dictionary[hello] = hola
stack dictionary[goodbye] = adios
stack dictionary[yes] = si
stack dictionary[no] = no
stack dictionary[thank you] = gracias
stack dictionary[friend] = amigo
stack dictionary[cat] = gato
stack dictionary[dog] = perro

lookup:
print
ask Enter English word (or 'quit'):
box word = answer

if box word = quit then goto end

print Spanish: stack dictionary[word]
goto lookup

end:
print Adios!
```

### Implementing a Shopping Cart

```kidlang
print === Shopping Cart ===

stack cart
stack prices
box cart_count = 0
box total = 0

// Available items
stack prices[apple] = 2
stack prices[banana] = 1
stack prices[orange] = 3
stack prices[milk] = 5
stack prices[bread] = 4

shopping:
print
print === SHOP ===
print 1. View items
print 2. Add to cart
print 3. View cart
print 4. Checkout
ask Choose:

if answer = 1 then goto view_items
if answer = 2 then goto add_item
if answer = 3 then goto view_cart
if answer = 4 then goto checkout
goto shopping

view_items:
print
print Available Items:
print Apple - $2
print Banana - $1
print Orange - $3
print Milk - $5
print Bread - $4
sleep 2000
goto shopping

add_item:
ask Item name:
box item = answer
box cart_count = box cart_count + 1
stack cart[cart_count] = box item
print Added to cart!
goto shopping

view_cart:
print
print Your Cart:
if box cart_count = 0 then
    print Cart is empty
    sleep 1000
    goto shopping
end

box i = 1
box total = 0
cart_loop:
if box i > box cart_count then goto cart_done
box item = stack cart[i]
box price = stack prices[item]
box total = box total + box price
print box i. box item - $box price
box i = box i + 1
goto cart_loop

cart_done:
print Total: $box total
sleep 2000
goto shopping

checkout:
print
print === CHECKOUT ===
print Total: $box total
print Thank you for shopping!
```

---

## Advanced Math

### Statistics Calculator

```kidlang
print === Statistics Calculator ===

stack numbers
box count = 5

// Get numbers from user
box i = 1
input_loop:
if box i > box count then goto calculate
ask Enter number box i:
stack numbers[i] = answer
box i = box i + 1
goto input_loop

calculate:
// Calculate sum
box sum = 0
box i = 1
sum_loop:
if box i > box count then goto sum_done
box sum = box sum + stack numbers[i]
box i = box i + 1
goto sum_loop

sum_done:
// Calculate average
box average = box sum / box count

// Find min and max
box min = stack numbers[1]
box max = stack numbers[1]
box i = 2

minmax_loop:
if box i > box count then goto minmax_done

if stack numbers[i] < box min then
    box min = stack numbers[i]
end

if stack numbers[i] > box max then
    box max = stack numbers[i]
end

box i = box i + 1
goto minmax_loop

minmax_done:
print
print === Results ===
print Sum: box sum
print Average: box average
print Minimum: box min
print Maximum: box max
```

### Matrix Operations (2D Array)

```kidlang
print === Matrix Demo ===

// 3x3 matrix using stack
stack matrix

// Row 1
stack matrix[1,1] = 1
stack matrix[1,2] = 2
stack matrix[1,3] = 3

// Row 2
stack matrix[2,1] = 4
stack matrix[2,2] = 5
stack matrix[2,3] = 6

// Row 3
stack matrix[3,1] = 7
stack matrix[3,2] = 8
stack matrix[3,3] = 9

// Display matrix
print Matrix:
print stack matrix[1,1] stack matrix[1,2] stack matrix[1,3]
print stack matrix[2,1] stack matrix[2,2] stack matrix[2,3]
print stack matrix[3,1] stack matrix[3,2] stack matrix[3,3]

// Calculate sum of diagonal
box diagonal_sum = stack matrix[1,1] + stack matrix[2,2] + stack matrix[3,3]
print Diagonal sum: box diagonal_sum
```

---

## String Manipulation

### Text Formatter

```kidlang
print === Text Formatter ===

ask Enter text:
box text = answer

print
print Original: box text
print Length: (use character count)

// Make a line of dashes
box dash = -
box line = box dash * 40
print box line

// Repeat text
box repeated = box text * 3
print Repeated: box repeated

// Extract characters
box char1 = box text / 1
box char2 = box text / 2
box char3 = box text / 3
print First 3 chars: box char1 box char2 box char3
```

### Simple Cipher

```kidlang
print === Secret Message ===

ask Enter message:
box message = answer

ask Enter key number:
box key = answer

// Encrypt using XOR
box encrypted = box message ^ box key
print Encrypted: box encrypted

// Decrypt (XOR again with same key)
box decrypted = box encrypted ^ box key
print Decrypted: box decrypted
```

---

## Best Practices

### Code Organization

```kidlang
// === CONSTANTS ===
box MAX_HEALTH = 100
box GAME_VERSION = 1.5

// === VARIABLES ===
box player_health = box MAX_HEALTH
box player_score = 0

// === MAIN PROGRAM ===
start:
// Your main code here

// === FUNCTIONS (using goto) ===
display_stats:
print Health: box player_health
print Score: box player_score
goto return_from_display

// === END ===
end_game:
print Game Over
```

### Error Handling

```kidlang
// Validate input
ask Enter age:
box age = answer

if box age < 0 then
    print Error: Age cannot be negative
    goto ask_again
end

if box age > 150 then
    print Error: Invalid age
    goto ask_again
end

// Age is valid, continue
print Valid age: box age
```

### Debugging Tips

```kidlang
// Use print statements to debug
box x = 10
print DEBUG: x = box x

box result = box x * 2
print DEBUG: after multiply, result = box result

// Mark sections clearly
print === STARTING CALCULATION ===
// calculation code
print === CALCULATION COMPLETE ===
```

---

## 🎯 Advanced Projects

Try building these:

1. **Personal Finance Tracker** - Track income and expenses with categories
2. **Game High Score System** - Save and load top 10 scores
3. **Password Manager** - Store encrypted passwords (educational only!)
4. **Recipe Book** - Store and search recipes
5. **Gradebook System** - Calculate GPAs and generate reports
6. **Inventory Management** - Track items with quantities and prices
7. **Chat Log Analyzer** - Process text files and generate statistics
8. **Number Base Converter** - Convert between binary, decimal, hex

---

## 🚀 Next Steps

You've mastered advanced KidLang! Now you can:
- Build complex applications
- Work with files and persistent data
- Implement real algorithms
- Create games and interactive programs

Keep experimenting and creating amazing things! 🎉
