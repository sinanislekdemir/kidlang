// Test nested conditionals
box a = 10
box b = 20

if box a > 5 then
print outer condition true
if box b > 15 then
print inner condition true
end
end

if box a < 5 then
print should not print
if box b > 15 then
print should not print either
end
end
