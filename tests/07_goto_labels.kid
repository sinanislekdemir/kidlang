// Test goto and labels
box counter = 0
start:
box counter = box counter + 1
print box counter
if box counter < 3 then goto start
print done
