// Functions help you reuse ideas.
// This program uses one function for words
// and one function for math.

function cheer(box name)
Hooray for box name !
end

function extra_stickers(box current)
return box current + 3
end

print ================================
print     STICKER PARTY
print ================================
print

ask What is your name?
box player = answer

box stickers = 5
box new_total = extra_stickers(box stickers)

print cheer(box player)
print You started with box stickers stickers.
print Now you have box new_total stickers.
print Your name has len(box player) letters.
