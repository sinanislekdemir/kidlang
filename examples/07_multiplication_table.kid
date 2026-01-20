// Multiplication Table Generator
// Practice your multiplication!

ask Which multiplication table do you want? (1-12):
box table = answer

if box table < 1 then
   print Please enter a number between 1 and 12
   goto end
end

if box table > 12 then
   print Please enter a number between 1 and 12
   goto end
end

print
print ================================
print   Multiplication Table of box table
print ================================

box i = 1

loop:
if box i > 10 then
   goto done
end

box result = box table * box i
print box table × box i = box result

box i = box i + 1
goto loop

done:
print ================================

end:
print Done!
