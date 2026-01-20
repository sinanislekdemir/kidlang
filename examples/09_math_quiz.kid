en
// Math Quiz Game
// Test your math skills!

print ================================
print       MATH QUIZ GAME
print ================================
print "answer" 10 math questions!
print
sleep 1000

box score = 0
box question = 1

quiz:
if box question > 10 then
   goto results
end

// Generate random numbers
box a = random % 20 + 1
box b = random % 20 + 1
box correct = box a + box b

print
print Question box question of 10
print What is box a + box b ?
ask Your answer:

if answer = box correct then
   print ✓ Correct!
   box score = box score + 1
   sleep 500
end

if answer != box correct then
   print ✗ Wrong! The answer was box correct
   sleep 1000
end

box question = box question + 1
goto quiz

results:
print
print ================================
print       QUIZ COMPLETE!
print ================================
print Your score: box score out of 10

box percentage = box score * 10
print Percentage: box percentage%
print

if box score = 10 then
   print PERFECT SCORE! You're a math genius! 🌟
end

if box score >= 7 then
   if box score < 10 then
      print Great job! Keep practicing! 👍
   end
end

if box score >= 5 then
   if box score < 7 then
      print Good effort! Try again to improve! 📚
   end
end

if box score < 5 then
   print Keep studying! Practice makes perfect! 💪
end

print
print Thanks for playing!
