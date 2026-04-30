stack toys
box first = stackset(toys, 1, robot)
box second = stackset(toys, 2, kite)
print box first
print box second
print stackhas(toys, 1)
print stackget(toys, 2)
box keys = stackkeys(toys)
print join(box keys, :)
print stacklen(toys)
print stackdelete(toys, 1)
print stacklen(toys)
print stackhas(toys, 1)
