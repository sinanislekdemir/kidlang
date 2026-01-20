// Test edge cases with zero and negative numbers
box zero = 0
box neg = -5
box pos = 5

if box zero = 0 then
print zero equals zero
end

if box neg < box zero then
print negative less than zero
end

if box pos > box zero then
print positive greater than zero
end

if box neg < box pos then
print negative less than positive
end

if box neg + box pos = 0 then
print sum is zero
end
