// Test complex OR conditions
box a = 5
box b = 10
box c = 15

if box a = 0 or box b = 0 or box c = 15 then
print at least one matches
end

if box a = 0 or box b = 0 or box c = 0 then
print should not print
end

if box a = 5 or box b = 99 or box c = 99 then
print first condition true
end
