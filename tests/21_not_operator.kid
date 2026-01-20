// Test NOT operator with complex conditions
box a = 5
box b = 10

if box a != 0 and box b != 0 then
print both not zero
end

if box a != 5 then
print should not print
end

if box a != 10 or box b != 5 then
print at least one not equal
end

box text = hello
if box text != goodbye then
print text not equal
end
