EN
// Middle Earth Trader - A trading adventure game
// Buy low, sell high, and earn 1000 gold!

print ====================================
print   MIDDLE EARTH TRADER
print ====================================
print
print You are a merchant in Middle Earth
print Goal: Earn 1000 gold in 30 days!
print

box day = 1
box gold = 100
box bag = 50
box used = 0
box potions = 0
box wands = 0
box armor = 0

start:
print
print === DAY box day of 30 ===
print Gold: box gold
print Bag: box used of box bag spaces
print Inventory - Potions: box potions | Wands: box wands | Armor: box armor
print

box p_price = random % 30 + 20
box w_price = random % 50 + 40  
box a_price = random % 80 + 60

print === MARKET PRICES ===
print Potions: box p_price gold
print Wands: box w_price gold
print Armor: box a_price gold (takes 2 spaces)
print
print === ACTIONS ===
print 1. Buy Potions
print 2. Buy Wands
print 3. Buy Armor
print 4. Sell Potions
print 5. Sell Wands
print 6. Sell Armor
print 7. Travel to next city
ask Choose:

box choice = answer

if box choice = 1 then goto buy_pot
if box choice = 2 then goto buy_wand
if box choice = 3 then goto buy_armor
if box choice = 4 then goto sell_pot
if box choice = 5 then goto sell_wand
if box choice = 6 then goto sell_armor
if box choice = 7 then goto travel
print Invalid choice!
sleep 1000
goto start

buy_pot:
ask How many potions?
box amt = answer
box cost = box p_price * box amt
box need = box used + box amt
if box cost > box gold then
   print Not enough gold!
   sleep 1000
   goto start
end
if box need > box bag then
   print Not enough bag space!
   sleep 1000
   goto start
end
box gold = box gold - box cost
box potions = box potions + box amt
box used = box used + box amt
print Bought box amt potions for box cost gold!
sleep 1000
goto start

buy_wand:
ask How many wands?
box amt = answer
box cost = box w_price * box amt
box need = box used + box amt
if box cost > box gold then
   print Not enough gold!
   sleep 1000
   goto start
end
if box need > box bag then
   print Not enough bag space!
   sleep 1000
   goto start
end
box gold = box gold - box cost
box wands = box wands + box amt
box used = box used + box amt
print Bought box amt wands for box cost gold!
sleep 1000
goto start

buy_armor:
ask How many armor sets?
box amt = answer
box cost = box a_price * box amt
box need = box used + box amt * 2
if box cost > box gold then
   print Not enough gold!
   sleep 1000
   goto start
end
if box need > box bag then
   print Not enough bag space!
   sleep 1000
   goto start
end
box gold = box gold - box cost
box armor = box armor + box amt
box used = box used + box amt * 2
print Bought box amt armor for box cost gold!
sleep 1000
goto start

sell_pot:
ask How many potions?
box amt = answer
if box amt > box potions then
   print You don't have that many!
   sleep 1000
   goto start
end
box earn = box p_price * box amt
box gold = box gold + box earn
box potions = box potions - box amt
box used = box used - box amt
print Sold box amt potions for box earn gold!
sleep 1000
goto start

sell_wand:
ask How many wands?
box amt = answer
if box amt > box wands then
   print You don't have that many!
   sleep 1000
   goto start
end
box earn = box w_price * box amt
box gold = box gold + box earn
box wands = box wands - box amt
box used = box used - box amt
print Sold box amt wands for box earn gold!
sleep 1000
goto start

sell_armor:
ask How many armor sets?
box amt = answer
if box amt > box armor then
   print You don't have that many!
   sleep 1000
   goto start
end
box earn = box a_price * box amt
box gold = box gold + box earn
box armor = box armor - box amt
box used = box used - box amt * 2
print Sold box amt armor for box earn gold!
sleep 1000
goto start

travel:
print
print Traveling to next city...
sleep 1000
box day = box day + 1

box event = random % 8

if box event = 0 then
   print A friendly wizard gives you 50 gold! ✨
   box gold = box gold + 50
   sleep 1500
end

if box event = 1 then
   print Bandits attack! Lost 30 gold! ⚔️
   if box gold > 30 then
      box gold = box gold - 30
   end
   if box gold <= 30 then
      box gold = 0
   end
   sleep 1500
end

if box event = 2 then
   print Found a potion on the road! 🧪
   if box used < box bag then
      box potions = box potions + 1
      box used = box used + 1
   end
   sleep 1500
end

if box gold >= 1000 then goto win
if box day > 30 then goto lose

goto start

win:
print
print ====================================
print   CONGRATULATIONS!
print ====================================
print You earned 1000 gold in box day days!
print You are a Master Merchant! 🏆
print
goto end

lose:
print
print ====================================
print   TIME IS UP!
print ====================================
print 30 days have passed...
print Final gold: box gold
print
if box gold >= 500 then
   print Not bad for a merchant!
end
if box gold < 500 then
   print Keep practicing your trading!
end
print
goto end

end:
print Thanks for playing Middle Earth Trader!
print Goodbye! 👋
