// Test very long condition chains
box a = 1
box b = 2
box c = 3
box d = 4
box e = 5

if box a = 1 and box b = 2 and box c = 3 and box d = 4 and box e = 5 then
print five conditions all true
end

if box a > 0 or box b > 0 or box c > 0 or box d > 0 or box e > 0 then
print at least one positive
end

if box a < 10 and box b < 10 and box c < 10 and box d < 10 and box e < 10 then
print all less than ten
end
