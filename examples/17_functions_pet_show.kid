// Functions can call other functions.
// This program makes a tiny pet show.

function pet_sound(box pet)
if box pet = cat then
   return meow
end

if box pet = dog then
   return woof
end

if box pet = duck then
   return quack
end

beep
end

function show_pet(box pet)
The box pet says pet_sound(box pet) !
end

print ================================
print       PET SHOW
print ================================
print

ask Pick a pet: cat, dog, or duck
box pet = answer

print show_pet(box pet)
print Loud pet name: upper(box pet)
