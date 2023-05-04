def maze_solver(maze, start_x, start_y, end_x, end_y):
   path = []
   walk(maze, start_x, start_y, end_x, end_y, path)
   return path

def walk(maze, current_x, current_y, end_x, end_y, path):
    # Base Cases: 
    # 1) Off the map
    # 2) # is a wall
    # 3) It's the end
    # 4) Visited

    # Off the map
    if current_x < 0 or current_x >= len(maze[0]) or current_y < 0 or current_y >= len(maze):
        return False
    # # is a wall
    if maze[current_y][current_x] == "#":
        return False
    # It's the end
    if current_y == end_y and current_x == end_x:
        path.append((current_x, current_y))
        return True
    # Visited
    if maze[current_y][current_x] == "X":
        return False
    
    # Pre
    maze[current_y][current_x] = "X"
    path.append((current_x, current_y))

    # Recurse
    if (walk(maze, current_x, current_y + 1, end_x, end_y, path) 
        or walk(maze, current_x + 1, current_y, end_x, end_y, path) 
        or walk(maze, current_x, current_y - 1, end_x, end_y, path) 
        or walk(maze, current_x - 1, current_y, end_x, end_y, path)):
        return True

    # Post
    path.pop()

    return False

maze = [["#", "#", "#", "#", "#", "", "#", "#"],
        ["#", "",  "#", "",  "",  "", "#", "#"],
        ["#" ,"",  "",  "",  "#", "#","#",  "#"],
        ["#", "", "#", "#", "#", "#", "#", "#"]]

print(maze_solver(maze, 1, 3, 5, 0))