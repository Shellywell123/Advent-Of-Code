# macros
RIGHT = 0
DOWN  = 1
LEFT  = 2
UP    = 3
VOID  = ' '
SPACE = '.'
WALL  = '#'

def ParseData(file):    
    """
    function to read and parse in data from a txt file
    """
    with open(file) as f:
        data = f.read().split('\n')
    
    map = []
    map_width = 0
    for y in range(0,len(data)):
        row = data[y]
        if row.strip() == '':
            instructions_raw = data[y+1]
            break        
        map.append(list(row))
        if len(list(row)) > map_width:
            map_width = len(row)+1

    instructions = instructions_raw.replace('L',',L,').replace('R',',R,').split(',')
    for i in range(0,len(instructions)-1):
        if instructions[i] not in ['L','R']:
            instructions[i] = int(instructions[i])

    # fill empty spaces
    for m in range(0,len(map)):
        new_row = map[m]
        while len(new_row) < map_width:
            new_row.append(' ')
        map[m] = new_row

    return map,instructions

def Part1(filename, printing = False):
    """
    Solutions to Part 1
    """
    map,instructions = ParseData(filename)
    map_copy = map[:]

    # find starting coord
    for x in range(0,len(map[0])):
        if map[0][x] == SPACE:
            start_coord = [x,0]
            break
    
    # intisialise stuff    
    facing = RIGHT #start facing right
    coord = start_coord
    next_coord = start_coord

    for instruction in instructions:

        if printing:
            print('instruction! = ',instruction)

        # turning
        if type(instruction) == str:
            if ((facing == UP    and instruction == 'L') or (facing == DOWN  and instruction == 'R')):
                facing = LEFT
                continue
            if ((facing == LEFT  and instruction == 'L') or (facing == RIGHT and instruction == 'R')):
                facing = DOWN
                continue
            if ((facing == DOWN  and instruction == 'L') or (facing == UP    and instruction == 'R')):
                facing = RIGHT
                continue
            if ((facing == RIGHT and instruction == 'L') or (facing == LEFT  and instruction == 'R')):
                facing = UP     
                continue

        # moving
        if type(instruction) == int:
        
            steps = 0
            while (steps < instruction):

                if facing == UP:
                    print('u')
                    map_copy[coord[1]][coord[0]] = '^'
                    next_coord = [coord[0],coord[1]-1]
                    # check for map boundaries and wrap arounds
                    if (next_coord[1] < 0) or (map[next_coord[1]][next_coord[0]] == VOID):
                        # then wrap around
                        for ny in range(len(map)-1,-1,-1):
                            print(ny,next_coord)
                            if map[ny][next_coord[0]] == SPACE:
                                next_coord = [next_coord[0],ny]
                                break
                            if map[ny][next_coord[0]] == WALL:
                                hit_wall = True
                                break
                
                if facing == DOWN:
                    print('d')
                    map_copy[coord[1]][coord[0]] = 'v'
                    next_coord = [coord[0],coord[1]+1]
                    hit_wall = False
                    # check for map boundaries and wrap arounds
                    if (next_coord[1] > len(map)-1) or (map[next_coord[1]][next_coord[0]] == VOID):
                        # then wrap around
                        for ny in range(0,len(map)-1):
                            if map[ny][next_coord[0]] == SPACE:
                                next_coord = [next_coord[0],ny]
                                break
                            if map[ny][next_coord[0]] == WALL:
                                hit_wall = True
                                break

                if facing == LEFT:
                    print('l')
                    map_copy[coord[1]][coord[0]] = '<'
                    next_coord = [coord[0]-1,coord[1]]
                    hit_wall = False
                    # check for map boundaries and wrap arounds
                    if (next_coord[0] < 0) or (map[next_coord[1]][next_coord[0]] == VOID):
                        # then wrap around
                        for nx in range(len(map[0])-1,-1,-1):
                            print(nx, instruction,len(map),map[next_coord[1]][nx])
                            if map[next_coord[1]][nx] == SPACE:
                                next_coord = [nx,next_coord[1]]
                                break
                            if map[next_coord[1]][nx] == WALL:
                                hit_wall = True
                                break
                
                if facing == RIGHT:
                    map_copy[coord[1]][coord[0]] = '>'
                    next_coord = [coord[0]+1,coord[1]]
                    hit_wall = False
                    # check for map boundaries and wrap arounds
                    if (next_coord[0] > len(map[0])-1) or (map[next_coord[1]][next_coord[0]] == VOID):
                        # then wrap around
                        for nx in range(0,len(map[0])-1):
                            print(nx, instruction,len(map),map[next_coord[1]][nx])
                            if map[next_coord[1]][nx] == SPACE:
                                next_coord = [nx,next_coord[1]]
                                break
                            if map[next_coord[1]][nx] == WALL:
                                hit_wall = True
                                break

                if hit_wall or map[next_coord[1]][next_coord[0]] == WALL:
                    print('HIT WALL')
                    # go to next instruction
                    break

                if printing:
                    print(instructions)
                    print(f'\nsteps = {steps}\ninstrcution = {instruction}\n')
                    for row in map:
                        rowstr = ''
                        for r in row:
                            rowstr += r
                        print(rowstr)

                coord = next_coord
                steps += 1
    
    map_copy[coord[1]][coord[0]] = 'X'
    if printing:
        print(instructions)
        print(f'\nsteps = {steps}\ninstrcution = {instruction}\n')
        for row in map:
            rowstr = ''
            for r in row:
                rowstr += r
            print(rowstr)
    
    print(coord,facing)
    return (1000 * (coord[1]+1)) + (4 * (coord[0]+1)) + facing

def Part2(filename, printing = False):
    """
    Solutions to Part 2
    """
    return 0

def main():
    """
    main function
    """ 
    testfile  = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day22\n")

    print(f'Tests : Answer to Part 1 = {Part1(testfile, printing = True)}')    
    # print(f'Tests : Answer to Part 2 = {Part2(testfile, printing = False)}\n')

    assert(Part1(testfile) == 6032)
    # assert(Part2(testfile) == 20)

    #print(f'Inputs: Answer to Part 1 = {Part1(inputfile, printing = False)}')
    # print(f'Inputs: Answer to Part 2 = {Part2(inputfile, printing = False)}')

if __name__ == "__main__":
    main()