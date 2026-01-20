// Test comparison operators with mixed types
box int1 = 5
box float1 = 5.0

if box int1 = box float1 then
print int equals float
end

box int2 = 3
box float2 = 3.5

if box int2 < box float2 then
print int less than float
end

if box float2 > box int2 then
print float greater than int
end

box str = 5
box num = 5

if box str = box num then
print string equals int
end
