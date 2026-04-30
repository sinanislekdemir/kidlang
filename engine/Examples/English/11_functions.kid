EN
// Functions are little helpers you can reuse.

function cheer(box name)
Hooray for box name !
end

function extra_stars(box current)
return box current + 2
end

ask What is your name?
box player = answer

box stars = 4
box total = extra_stars(box stars)

print cheer(box player)
print You now have box total stars.
print Your name has len(box player) letters.
