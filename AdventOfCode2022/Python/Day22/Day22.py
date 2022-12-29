UP    = 0
DOWN  = 1
LEFT  = 2
RIGHT = 3
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
    for y in range(0,len(data)):
        row = data[y]
        if row.strip() == '':
            instructions_raw = data[y+1]
            break        
        map.append(row)

    instructions = instructions_raw.replace('L',',L,').replace('R',',R,').split(',')
    for i in range(0,len(instructions)-1):
        if instructions[i] not in ['L','R']:
            instructions[i] = int(instructions[i])

    return map,instructions

def Part1(filename, printing = False):
    """
    Solutions to Part 1
    """
    map,instructions = ParseData(filename)

    if printing:
        print(instructions)
        for row in map:
            print(row)

    # find starting coord
    for x in range(0,len(map[0])):
        if map[0][x] == SPACE:
            start_coord = [x,0]
            break
    
    # intisialise stuff
    #start facing right
    facing = RIGHT
    coord = start_coord

    for instruction in instructions:

        # turning
        if type(instruction) == str:
            if ((facing == UP    and instruction == 'L') or (facing == DOWN  and instruction == 'R')):
                facing = LEFT
            if ((facing == LEFT  and instruction == 'L') or (facing == RIGHT and instruction == 'R')):
                facing = DOWN
            if ((facing == DOWN  and instruction == 'L') or (facing == UP    and instruction == 'R')):
                facing = RIGHT
            if ((facing == RIGHT and instruction == 'L') or (facing == LEFT  and instruction == 'R')):
                facing = UP            

        # moving
        if type(instruction) == int:
        
            steps = 0

            while (steps < instruction):

                if facing == UP:
                    next_coord = start_coord[0],[start_coord[1]-1]
                    # check for map boundaries and wrap arounds
                    if (next_coord[1] < len(map)-1) or (0 > next_coord[1]) or (map[next_coord[1],next_coord[0]] == VOID):
                        # then wrap around
                        for ny in range(len(map),0,-1):
                            if map[ny][next_coord[0]] == SPACE:
                                next_coord = [next_coord[0],ny]
                                break
                            # need to think about wrap arounds to walls
                
                if facing == DOWN:
                    next_coord = start_coord[0],[start_coord[1]+1]

                if facing == LEFT:
                    next_coord = start_coord[0]-1,[start_coord[1]]
                
                if facing == RIGHT:
                    next_coord = start_coord[0]+1,[start_coord[1]]
                
                # if map[next_coord[1]][next_coord[0]] == WALL:
                #     # go to next instructoin
                #     break
                steps += 1

    return 0

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

    # assert(Part1(testfile) == 110)
    # assert(Part2(testfile) == 20)

    # print(f'Inputs: Answer to Part 1 = {Part1(inputfile, printing = False)}')
    # print(f'Inputs: Answer to Part 2 = {Part2(inputfile, printing = False)}')

if __name__ == "__main__":
    main()