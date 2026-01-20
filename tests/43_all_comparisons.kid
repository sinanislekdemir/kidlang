// Test all comparison operators
box a = 10
box b = 10
box c = 5

if box a = box b then
print equal
end

if box a != box c then
print not equal
end

if box a > box c then
print greater
end

if box c < box a then
print less
end

if box a >= box b then
print greater or equal
end

if box c <= box a then
print less or equal
end
