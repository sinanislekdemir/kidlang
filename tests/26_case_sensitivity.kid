// Test case sensitivity in string comparisons
box text1 = Hello
box text2 = hello
box text3 = HELLO

if box text1 = box text2 then
print case insensitive match
end

if box text2 = box text3 then
print all case insensitive
end

box mixed = HeLLo
if box mixed = box text1 then
print mixed case matches
end
