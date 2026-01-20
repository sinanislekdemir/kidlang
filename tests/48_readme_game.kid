// README example: Game Logic
box guess = 7
box secret = 7

if box guess = box secret then
goto winner
end
print Sorry, try again!
goto end

winner:
print You won! The number was box secret

end:
