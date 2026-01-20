EN
stack items
box count = 0
print === Shopping List ===
print
menu:
print What do you want to do?
print 1 Add item
print 2 Show list
print 3 Quit
ask Choose:
box choice = answer

if box choice = 1 then
ask Enter item name:
box name = answer
box count = box count + 1
stack items[box count] = box name
print Added!
goto menu
end

if box choice = 2 then
print
print Your shopping list:
box i = 1
show:
if box i <= box count then
print box i . stack items[box i]
box i = box i + 1
goto show
end
print
goto menu
end

if box choice = 3 then
print Goodbye!
end
