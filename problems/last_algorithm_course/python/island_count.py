def island_count(map):
    counter = 0
    for y in range(len(map)):
        for x in range(len(map[0])):
              if map[y][x] == 1:
                    count(map, x, y)
                    counter += 1
    return counter

# Base Condition
# 1) Out of Map
# 2) It's a 0
# 3) Visited could be marked as 0
def count(map, x, y):
    if (x < 0 or x >= len(map[0]) or 
        y < 0 or y >= len(map)):
        return
    if map[y][x] == 0:
         return
    
    map[y][x] = 0
    count(map, x, y+1)
    count(map, x+1, y)
    count(map, x, y-1)
    count(map, x-1, y)



map = [[1,1,1,1,1,0,1,0,0],
       [1,1,1,0,0,0,0,1,1],
       [1,0,0,0,0,1,0,1,1]]
print(island_count(map))