// Test boolean values in conditions
box a = 5
box b = 10

if box a < box b then
print first comparison true
end

if box a > box b then
print should not print
end

if box a > box b then
print should not print
end

if box a < box b and box b > 5 then
print combined condition
end
