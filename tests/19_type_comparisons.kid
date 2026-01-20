// Test string to number comparisons
box num = 42
box str = 42

if box num = box str then
print number equals string
end

box num2 = 3.14
box str2 = 3.14

if box num2 = box str2 then
print float equals string
end

box text = hello
box num3 = 5

if box text = box num3 then
print should not print
end

if box text != box num3 then
print text not equal number
end
