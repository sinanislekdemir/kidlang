// Test complex AND conditions
box a = 5
box b = 10
box c = 15

if box a > 0 and box b > 5 and box c > 10 then
print all three conditions met
end

if box a > 0 and box b > 5 and box c > 20 then
print should not print
end

if box a = 5 and box b = 10 and box c = 15 then
print exact match all three
end
