// Simple Calculator
// Perform basic math operations

menu:
print
print ================================
print        CALCULATOR
print ================================
print 1. Addition
print 2. Subtraction
print 3. Multiplication
print 4. Division
print 5. Square
print 6. Square Root
print 7. Exit
print ================================
ask Choose operation (1-7):

if answer = 1 then goto add
if answer = 2 then goto subtract
if answer = 3 then goto multiply
if answer = 4 then goto divide
if answer = 5 then goto square
if answer = 6 then goto squareroot
if answer = 7 then goto exit

print Invalid choice! Please enter 1-7
goto menu

add:
ask Enter first number:
box a = answer
ask Enter second number:
box b = answer
box result = box a + box b
print Result: box a + box b = box result
goto menu

subtract:
ask Enter first number:
box a = answer
ask Enter second number:
box b = answer
box result = box a - box b
print Result: box a - box b = box result
goto menu

multiply:
ask Enter first number:
box a = answer
ask Enter second number:
box b = answer
box result = box a * box b
print Result: box a × box b = box result
goto menu

divide:
ask Enter first number:
box a = answer
ask Enter second number:
box b = answer

if answer = 0 then
   print Error: Cannot divide by zero!
   goto menu
end

box result = box a / box b
print Result: box a ÷ box b = box result
goto menu

square:
ask Enter a number:
box n = answer
box result = sqr box n
print Result: box n² = box result
goto menu

squareroot:
ask Enter a number:
box n = answer
box result = sqrt box n
print Result: √box n = box result
goto menu

exit:
print
print Thank you for using Calculator!
print Goodbye!
