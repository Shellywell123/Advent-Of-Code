def ParseData(file):    
    """
    function to read and parse in data from a txt file
    """
    with open(file) as f:
        data = f.read()
    

    return list(data)

def Solve(filename, num_of_rocks, printing = False):
    """
    Play Tetris
    """

    jets = ParseData(filename)

    rock = '#'
    space = '.'
    wall_lims = [-1,7]

    chamber = []
    chamber.append([rock,rock,rock,rock,rock,rock,rock])
    for i in range(0,6):
        chamber.append([space,space,space,space,space,space,space])

    current_highest_point = 0
    num_of_rocks_fallen = 0
    piece_type = 1
    j = 0
    cycle = 0
    cycle_len = 0
    cycle_count = 0
    extra = 0
    while num_of_rocks_fallen < num_of_rocks:
        print(cycle_count,num_of_rocks_fallen)

        if len(chamber) % 2000:
            for i in range (0,len(chamber)-1000):
                chamber[i] = []

        #optimizaion cycle begins after len(jets) ** 2
        if num_of_rocks_fallen % len(jets) ** 2 == 0:
            if cycle != 0 and num_of_rocks % ( len(jets) ** 2) == 0:
                cycle_height_increse = current_highest_point - cycle
                
                #num_of_rocks_fallen = numbernum_of_rocks % 
                cycle_len += cycle_height_increse
                
               # print(cycle_height_increse)
            cycle = current_highest_point
            cycle_count+=1
            print(cycle_count)
            if cycle_count == 7:
                seventh = cycle_len
               # print('-a',cycle_height_increse,cycle_len)
            if cycle_count == 14:
                print('cycle found')
                forteenth = cycle_len
                #print('-b')

                cycle_len = forteenth - seventh
                num_of_cylces = num_of_rocks / (7 * len(jets) ** 2)
                extra = cycle_len * int(num_of_cylces)
                print(cycle_len,extra)
                current_highest_point = 0
                num_of_rocks_fallen = 0
                piece_type = 1
                j = 0
                chamber = []
                chamber.append([rock,rock,rock,rock,rock,rock,rock])
                for i in range(0,6):
                    chamber.append([space,space,space,space,space,space,space])
                num_of_rocks = num_of_rocks - (int(num_of_cylces) * 7 * (len(jets) ** 2))
                print(num_of_rocks)
                continue#print(f,cycle_len,num_of_rocks,num_of_cylces,(int(num_of_cylces) * 7 * len(jets) ** 2))
                #exit()
            
         #   print(num_of_rocks_fallen / len(jets) ** 2, current_highest_point,len(jets))

        rock_falling = True

        # extend chamber height
        while current_highest_point + 8 > len(chamber)-1:
            chamber.append([space,space,space,space,space,space,space])

        #.---------------------------------------.
        #| a = highest point in each piece_type  |
        #| b = lowest point in each piece_type   |
        #.---------------------------------------.

        # -
        if piece_type == 1:
            a = [wall_lims[0]+3, current_highest_point + 4]
            b = [wall_lims[0]+4, current_highest_point + 4]
            c = [wall_lims[0]+5, current_highest_point + 4]
            d = [wall_lims[0]+6, current_highest_point + 4]            
        
        # +
        elif piece_type == 2:
            a = [wall_lims[0]+4, current_highest_point + 6]
            b = [wall_lims[0]+4, current_highest_point + 4]
            c = [wall_lims[0]+4, current_highest_point + 5]
            d = [wall_lims[0]+5, current_highest_point + 5]
            e = [wall_lims[0]+3, current_highest_point + 5]
        
        # _|
        elif piece_type == 3:
            a = [wall_lims[0]+5, current_highest_point + 6]
            b = [wall_lims[0]+4, current_highest_point + 4]
            c = [wall_lims[0]+5, current_highest_point + 4]
            d = [wall_lims[0]+5, current_highest_point + 5]
            e = [wall_lims[0]+3, current_highest_point + 4]
        
        # |
        elif piece_type == 4:
            a = [wall_lims[0]+3, current_highest_point + 7]
            b = [wall_lims[0]+3, current_highest_point + 4]
            c = [wall_lims[0]+3, current_highest_point + 6]
            d = [wall_lims[0]+3, current_highest_point + 5]
        
        # box
        elif piece_type == 5:
            a = [wall_lims[0]+3, current_highest_point + 5]
            b = [wall_lims[0]+3, current_highest_point + 4]
            c = [wall_lims[0]+4, current_highest_point + 4]
            d = [wall_lims[0]+4, current_highest_point + 5]

        while rock_falling:
            jet = jets[j]
            if piece_type == 1:
                
                # left side collison
                if (jet == '<') and not (
                    (a[0]-1 == wall_lims[0]) or (chamber[a[1]][a[0]-1] == rock)):
                    a[0] = a[0] - 1
                    b[0] = b[0] - 1
                    c[0] = c[0] - 1
                    d[0] = d[0] - 1
                
                # right side collision
                elif (jet == '>') and not (
                    (d[0]+1 == wall_lims[1]) or (chamber[d[1]][d[0]+1] == rock)):
                    a[0] = a[0] + 1
                    b[0] = b[0] + 1
                    c[0] = c[0] + 1
                    d[0] = d[0] + 1
                
                # bottom collision
                if ((chamber[a[1]-1][a[0]] == rock) or
                    (chamber[b[1]-1][b[0]] == rock) or
                    (chamber[c[1]-1][c[0]] == rock) or
                    (chamber[d[1]-1][d[0]] == rock)):
                    chamber[a[1]][a[0]] = rock
                    chamber[b[1]][b[0]] = rock
                    chamber[c[1]][c[0]] = rock
                    chamber[d[1]][d[0]] = rock
                    if a[1] > current_highest_point:
                        current_highest_point = a[1]    
                    rock_falling = False

                # fall                
                a[1] = a[1] - 1
                b[1] = b[1] - 1
                c[1] = c[1] - 1
                d[1] = d[1] - 1
            
            elif piece_type == 2:
            
                # left side collison
                if (jet == '<') and not (
                    (a[0]-1 == wall_lims[0]) or (chamber[a[1]][a[0]-1] == rock) or
                    (b[0]-1 == wall_lims[0]) or (chamber[b[1]][b[0]-1] == rock) or 
                    (e[0]-1 == wall_lims[0]) or (chamber[e[1]][e[0]-1] == rock)):
                    a[0] = a[0] - 1
                    b[0] = b[0] - 1
                    c[0] = c[0] - 1
                    d[0] = d[0] - 1
                    e[0] = e[0] - 1
                
                # right side collision
                elif (jet == '>') and not (
                    (a[0]+1 == wall_lims[1]) or (chamber[a[1]][a[0]+1] == rock) or
                    (b[0]+1 == wall_lims[1]) or (chamber[b[1]][b[0]+1] == rock) or 
                    (d[0]+1 == wall_lims[1]) or (chamber[d[1]][d[0]+1] == rock)):
                    a[0] = a[0] + 1
                    b[0] = b[0] + 1
                    c[0] = c[0] + 1
                    d[0] = d[0] + 1
                    e[0] = e[0] + 1
                
                # bottom collision
                if ((chamber[b[1]-1][b[0]] == rock) or
                    (chamber[d[1]-1][d[0]] == rock) or
                    (chamber[e[1]-1][e[0]] == rock)):
                    chamber[a[1]][a[0]] = rock
                    chamber[b[1]][b[0]] = rock
                    chamber[c[1]][c[0]] = rock
                    chamber[d[1]][d[0]] = rock 
                    chamber[e[1]][e[0]] = rock
                    if a[1] > current_highest_point:
                        current_highest_point = a[1]                  
                    rock_falling = False
                
                # fall
                a[1] = a[1] - 1
                b[1] = b[1] - 1
                c[1] = c[1] - 1
                d[1] = d[1] - 1
                e[1] = e[1] - 1
            
            elif piece_type == 3:

                # left side collison
                if (jet == '<') and not (
                    (a[0]-1 == wall_lims[0]) or (chamber[a[1]][a[0]-1] == rock) or
                    (d[0]-1 == wall_lims[0]) or (chamber[d[1]][d[0]-1] == rock) or 
                    (e[0]-1 == wall_lims[0]) or (chamber[e[1]][e[0]-1] == rock)):
                    a[0] = a[0] - 1
                    b[0] = b[0] - 1
                    c[0] = c[0] - 1
                    d[0] = d[0] - 1
                    e[0] = e[0] - 1
                
                # right side collision
                elif (jet == '>') and not (
                    (a[0]+1 == wall_lims[1]) or (chamber[a[1]][a[0]+1] == rock) or
                    (d[0]+1 == wall_lims[1]) or (chamber[d[1]][d[0]+1] == rock) or 
                    (c[0]+1 == wall_lims[1]) or (chamber[c[1]][c[0]+1] == rock)):
                    a[0] = a[0] + 1
                    b[0] = b[0] + 1
                    c[0] = c[0] + 1
                    d[0] = d[0] + 1
                    e[0] = e[0] + 1
                
                # bottom collision
                if ((chamber[b[1]-1][b[0]] == rock) or
                    (chamber[c[1]-1][c[0]] == rock) or
                    (chamber[e[1]-1][e[0]] == rock)):
                    chamber[a[1]][a[0]] = rock
                    chamber[b[1]][b[0]] = rock
                    chamber[c[1]][c[0]] = rock
                    chamber[d[1]][d[0]] = rock
                    chamber[e[1]][e[0]] = rock
                    if a[1] > current_highest_point:
                        current_highest_point = a[1]    
                    rock_falling = False
                
                # fall
                a[1] = a[1] - 1
                b[1] = b[1] - 1
                c[1] = c[1] - 1
                d[1] = d[1] - 1
                e[1] = e[1] - 1
            
            elif piece_type == 4:

                # left side collison
                if (jet == '<') and not (
                    (a[0]-1 == wall_lims[0]) or (chamber[a[1]][a[0]-1] == rock) or
                    (b[0]-1 == wall_lims[0]) or (chamber[b[1]][b[0]-1] == rock) or 
                    (c[0]-1 == wall_lims[0]) or (chamber[c[1]][c[0]-1] == rock) or
                    (d[0]-1 == wall_lims[0]) or (chamber[d[1]][d[0]-1] == rock)):
                    a[0] = a[0] - 1
                    b[0] = b[0] - 1
                    c[0] = c[0] - 1
                    d[0] = d[0] - 1
                
                # right side collision
                elif (jet == '>') and not(
                    (a[0]+1 == wall_lims[1]) or (chamber[a[1]][a[0]+1] == rock) or
                    (b[0]+1 == wall_lims[1]) or (chamber[b[1]][b[0]+1] == rock) or 
                    (c[0]+1 == wall_lims[1]) or (chamber[c[1]][c[0]+1] == rock) or
                    (d[0]+1 == wall_lims[1]) or (chamber[d[1]][d[0]+1] == rock)):
                    a[0] = a[0] + 1
                    b[0] = b[0] + 1
                    c[0] = c[0] + 1
                    d[0] = d[0] + 1
            
                # bottom collision
                if ((chamber[b[1]-1][b[0]] == rock)):
                    chamber[a[1]][a[0]] = rock
                    chamber[b[1]][b[0]] = rock
                    chamber[c[1]][c[0]] = rock
                    chamber[d[1]][d[0]] = rock
                    if a[1] > current_highest_point:
                        current_highest_point = a[1]    
                    rock_falling = False
                
                # fall
                a[1] = a[1] - 1
                b[1] = b[1] - 1
                c[1] = c[1] - 1
                d[1] = d[1] - 1
            
            elif piece_type == 5:
                
                # left side collison
                if (jet == '<') and not (
                    (a[0]-1 == wall_lims[0]) or (chamber[a[1]][a[0]-1] == rock) or
                    (b[0]-1 == wall_lims[0]) or (chamber[b[1]][b[0]-1] == rock)):
                    a[0] = a[0] - 1
                    b[0] = b[0] - 1
                    c[0] = c[0] - 1
                    d[0] = d[0] - 1
                
                # right side collision
                elif (jet == '>') and not ( 
                    (c[0]+1 == wall_lims[1]) or (chamber[c[1]][c[0]+1] == rock) or
                    (d[0]+1 == wall_lims[1]) or (chamber[d[1]][d[0]+1] == rock)):
                    a[0] = a[0] + 1
                    b[0] = b[0] + 1
                    c[0] = c[0] + 1
                    d[0] = d[0] + 1

                # bottom collision
                if ((chamber[b[1]-1][b[0]] == rock) or
                    (chamber[c[1]-1][c[0]] == rock)):
                    chamber[a[1]][a[0]] = rock
                    chamber[b[1]][b[0]] = rock
                    chamber[c[1]][c[0]] = rock
                    chamber[d[1]][d[0]] = rock    
                    if a[1] > current_highest_point:
                        current_highest_point = a[1]                
                    rock_falling = False
                
                # fall
                a[1] = a[1] - 1
                b[1] = b[1] - 1
                c[1] = c[1] - 1
                d[1] = d[1] - 1
                   
            # update jet increments
            if j == len(jets) - 1:
                j = -1
            j += 1
        
        # update of rocks fallen stats   
        num_of_rocks_fallen += 1

        # clear buffer
        a = []
        b = []
        c = []
        d = []
        e = []

        # update type increments
        if piece_type == 5:
            piece_type = 0
        piece_type += 1

    if printing:
            print()
            for p in range(len(chamber)-1, -1 ,-1):
                pstr = '|'
                for ps in chamber[p]:
                    pstr += ps
                pstr += '|'
                pint = str(p)
                if p <1000:
                    pint = '0'+str(p)
                if p <100:
                    pint = '00'+str(p)
                if p <10:
                    pint = '000'+str(p)                
                print(pint,pstr)

    print(current_highest_point,extra)
    return current_highest_point + extra

def Part1(filename, printing = False):
    """
    Solutions to Part 1
    """

    return Solve(filename, 2022, printing)

def Part2(filename, printing = False):
    """
    Solutions to Part 2
    """

    return Solve(filename,1000000000000, printing)

def main():
    """
    main function
    """ 
    testfile  = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day17")

    print(f'Tests : Answer to Part 1 = {Part1(testfile, printing = False)}')    
    print(f'Tests : Answer to Part 2 = {Part2(testfile, printing = False)}\n')

    #assert(Part1(testfile) == 3068)
    #assert(Part2(testfile) == 1514285714288)

   # print(f'Inputs: Answer to Part 1 = {Part1(inputfile, printing = False)}')
    print(f'Inputs: Answer to Part 2 = {Part2(inputfile, printing = False)}')

if __name__ == "__main__":
    main()