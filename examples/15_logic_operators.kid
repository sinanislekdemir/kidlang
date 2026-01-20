// Logic Operators Example - AND/OR
print Welcome to Logic Test!
print 

box age = 0
box haslicense = 0

ask What is your age?
box age = answer

ask Do you have a license? (1=yes, 0=no)
box haslicense = answer

// Test AND operator
if box age > 17 and box haslicense = 1 then
    print You can drive a car!
end

// Test OR operator  
if box age < 13 or box age > 65 then
    print You get a discount!
end

// Complex condition
if box age > 15 and box age < 20 then
    print You are a teenager!
end

print 
print Thanks for testing!
