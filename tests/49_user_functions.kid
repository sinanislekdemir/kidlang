function add(box a, box b)
box a + box b
end

function double_and_add_one(box value)
box doubled = box value * 2
return box doubled + 1
end

box total = add(4, 5)
print box total
box next = double_and_add_one(6)
print box next
