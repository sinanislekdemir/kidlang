// Random Story Generator
// Create funny stories with random elements!

print ================================
print   STORY GENERATOR
print ================================
print Let's create a funny story!
print

ask Enter a person's name:
box name = answer

ask Enter an adjective (like funny, big, red):
box adjective = answer

ask Enter a noun (like dog, car, pizza):
box noun = answer

ask Enter a verb (like run, jump, dance):
box verb = answer

ask Enter a place (like park, school, moon):
box place = answer

ask Enter a number:
box number = answer

print
print ================================
print      YOUR STORY
print ================================
print

// Generate random story variation
box variation = random % 3

if box variation = 0 then goto story1
if box variation = 1 then goto story2
goto story3

story1:
print Once upon a time, box name went to box place
print to find a box adjective box noun
print They decided to box verb exactly box number times!
print Everyone laughed and lived happily ever after.
goto end

story2:
print In a box adjective kingdom, box name discovered
print a magical box noun in box place
print It could box verb box number times per day!
print And that's how box name became famous!
goto end

story3:
print box name was walking in box place when suddenly,
print a box adjective box noun appeared and started to box verb
print This happened box number times that day!
print What an adventure!
goto end

end:
print
print ================================
print The End! 📖
print
ask Want another story? (yes/no):

if answer = yes then
   print
   goto start
end

start:
print Goodbye!
