en
// Simple Adventure Game
// Make choices and explore!

print ================================
print    TREASURE HUNT ADVENTURE
print ================================
print

box health = 100
box gold = 0

start:
print You wake up in a mysterious forest.
print Your health: box health
print Your gold: box gold
print
print What do you want to do?
print 1. Explore the forest
print 2. Check your bag
print 3. Rest
print 4. Quit game
ask Choose:

if answer = 1 then goto explore
if answer = 2 then goto check_bag
if answer = 3 then goto rest
if answer = 4 then goto quit

print Invalid choice!
goto start

explore:
print
print You venture deeper into the forest...
sleep 1000

box event = random % 3

if box event = 0 then goto find_gold
if box event = 1 then goto find_monster
goto find_nothing

find_gold:
box found = random % 20 + 10
box gold = box gold + box found
print You found box found gold coins! ✨
print Total gold: box gold
sleep 2000
goto start

find_monster:
print A monster appears! 👹
box damage = random % 30 + 10
box health = box health - box damage
print It attacks you for box damage damage!
print Health remaining: box health

if box health <= 0 then
   goto game_over
end

sleep 2000
goto start

find_nothing:
print You found nothing interesting here.
sleep 1000
goto start

check_bag:
print
print === YOUR STATUS ===
print Health: box health
print Gold: box gold
print
if box gold >= 100 then
   print You have enough gold to win!
   print Congratulations! 🏆
   goto quit
end
print You need 100 gold to win.
print Keep exploring!
sleep 2000
goto start

rest:
print
print You rest for a while...
sleep 1500
box heal = 20
box health = box health + box heal
if box health > 100 then
   box health = 100
end
print You recovered box heal health!
print Health: box health
sleep 1500
goto start

game_over:
print
print ================================
print       GAME OVER
print ================================
print You were defeated!
print Final gold: box gold
print
goto quit

quit:
print
print Thanks for playing!
print Goodbye! 👋
