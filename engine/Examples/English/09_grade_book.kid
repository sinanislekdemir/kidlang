EN
stack scores
stack names
print === Grade Book ===
box count = 0

menu:
print
print 1 Add student
print 2 Show all students
print 3 Quit
ask Choose:
box choice = answer

if box choice = 1 then
ask Student name:
box name = answer
ask Student score:
box score = answer
box count = box count + 1
stack names[box count] = box name
stack scores[box count] = box score
print Added!
goto menu
end

if box choice = 2 then
print
print === All Students ===
box i = 1
show:
if box i <= box count then
box name = stack names[box i]
box score = stack scores[box i]
print box name : box score
box i = box i + 1
goto show
end
goto menu
end

if box choice = 3 then
print Goodbye!
end
