EN
// Number Guessing Game
// Try to guess the secret number!

print ================================
print    NUMBER GUESSING GAME
print ================================
print

box secret = random % 50 + 1
box tries = 0
box max_tries = 7

print I'm thinking of a number between 1 and 50
print You have box max_tries tries to guess it!
print

game:
box tries = box tries + 1

print Try box tries of box max_tries
ask Enter your guess:

if answer = box secret then
   goto win
end

if answer > box secret then
   print Too high! Try a smaller number.
end

if answer < box secret then
   print Too low! Try a bigger number.
end

if box tries >= box max_tries then
   goto lose
end

print
goto game

win:
print
print ================================
print    CONGRATULATIONS!
print ================================
print You guessed it! The number was box secret
print It took you box tries tries!
goto ending

lose:
print
print ================================
print    GAME OVER
print ================================
print Sorry, you ran out of tries!
print The secret number was box secret
goto ending

ending:
print
print Thanks for playing!
