// Test mixed AND/OR conditions
box a = 5
box b = 10
box c = 15

if box a > 0 and box b > 5 or box c = 0 then
print and has precedence
end

if box a = 0 or box b = 10 and box c = 15 then
print or with and
end

if box a = 5 and box b = 10 or box c = 99 then
print mixed true
end
