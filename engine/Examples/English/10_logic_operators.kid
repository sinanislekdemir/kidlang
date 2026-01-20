EN
print === Logic Operators Demo ===
print
ask Enter your age:
box age = answer
ask Are you a student? (yes/no):
box student = answer

print
print --- Using AND ---
if box age >= 18 and box student = yes then
print You are an adult student!
end

print
print --- Using OR ---
if box age < 13 or box age > 65 then
print You get a discount!
end

print
print --- Combined ---
if box age >= 18 and box student = no or box age > 65 then
print You might want to consider studying!
end

print
print Done!
