//Working with Multiple Values

print Multiple Values Demo
print

// Create a stack for scores
stack scores

// Add some scores
stack scores[1] = 85
stack scores[2] = 92
stack scores(3) = 78
stack scores[4] = 95

print Test Scores:
print Test 1: stack scores[1]
print Test 2: stack scores[2]
print Test 3: stack scores[3]
print Test 4: stack scores[4]
print

// Use variables as indices
box test_number = 2
print Your score on test box test_number was: stack scores[test_number]
print

// Use strings as keys
stack inventory
stack inventory[apples] = 5
stack inventory[oranges] = 3
stack inventory(bananas) = 7

print Fruit Inventory:
print Apples: stack inventory[apples]
print Oranges: stack inventory[oranges]
print Bananas: stack inventory[bananas]
print

// Mixed types as indices
stack player_data
stack player_data[1] = Alice
stack player_data[2] = Bob
stack player_data[name] = Player One
stack player_data[score] = 1000
stack player_data[level] = 5

print Player Info:
print Player 1: stack player_data[1]
print Player 2: stack player_data[2]
print Current Player: stack player_data[name]
print Score: stack player_data[score]
print Level: stack player_data[level]
print

print Demo Complete!
