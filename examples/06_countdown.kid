en
// Countdown Timer
// Learn about loops and sleep

ask Enter countdown number:
box count = answer

print
print Starting countdown...
print

loop:
if box count < 0 then
   goto done
end

print box count
sleep 1000
box count = box count - 1
goto loop

done:
print
print BLAST OFF!
