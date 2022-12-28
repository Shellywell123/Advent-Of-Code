def ParseData(file):    
    """
    function to read and parse in data from a txt file
    """
    with open(file) as f:
        data = f.read().split('\n')
    
    elves = []
    elf_no = 0
    for y in range(0,len(data)):
        row = data[y]
        for x in range(0,len(row)):
            elf = row[x]
            if elf == '#':
                elves.append({'elf_no':elf_no,'pos':[x+9,y+9],'next_pos':0, 'last_dir':-1})
                elf_no += 1

    return elves

def Part1(filename, printing = False):
    """
    Solutions to Part 1
    """
    elves = ParseData(filename)

    for r in range(0,10):
        if printing:            
            print()
        # propose next postions
        for e in range(0,len(elves)):

            elf = dict(elves[e])

            [x,y] = elf['pos']

            N_set  = [[x  ,y-1],[x+1,y-1],[x-1,y-1]]
            S_set  = [[x  ,y+1],[x+1,y+1],[x-1,y+1]]
            E_set  = [[x+1,y  ],[x+1,y-1],[x+1,y+1]]
            W_set  = [[x-1,y  ],[x-1,y-1],[x-1,y+1]]

            next_pos_sets = [N_set,S_set,W_set,E_set]
            
            proposed = False
            while not proposed:
                current_elf_positions =  [elf['pos'] for elf in elves]  

                # if no other elf nearbye do nothing
                if not any((epos in current_elf_positions) for epos in N_set+S_set+E_set+W_set):
                    elf['next_pos'] = elf['pos']
                    elves[e] = elf
                    proposed = True
                    break

                # check if surrounded at all sides:
                elif all(any((epos in current_elf_positions) for epos in set) for set in next_pos_sets):
                    elf['next_pos'] = elf['pos']
                    elves[e] = elf
                    proposed = True
                    break

                # propse
                else:  
                    for d in range(0,4):
                        next_pos_set = next_pos_sets[(d+r)%4]  
                        if ((next_pos_set[0] not in current_elf_positions) and
                        (next_pos_set[1] not in current_elf_positions) and
                        (next_pos_set[2] not in current_elf_positions)):  
                            elf['next_pos'] = next_pos_set[0]
                            elf['last_dir'] = d
                            elves[e] = elf
                            proposed = True
                            break
            if printing:
                print(r+1,elf,list('nswe')[elf['last_dir']%4])
            
        # evaluate proposed positions
        proposed_elf_positions =  [elf['next_pos'] for elf in elves]

        for e in range(0,len(elves)):
            elf = dict(elves[e])
            # dups
            if proposed_elf_positions.count(elf['next_pos']) > 1:
                elf['next_pos'] = elf['pos']
            else:
                elf['pos'] = elf['next_pos']            
            elves[e] = elf

        xs = []
        ys = []
        
        for e in range(0,len(elves)):
            elf = dict(elves[e])
            xs.append(elf['pos'][0])
            ys.append(elf['pos'][1])

        if printing:   
            print()                   
            grid = []
            for y in range(-max(ys),max(ys)+1):
                row = []
                for x in range(-max(xs),max(xs)+1):
                    row.append('.')
                grid.append(row)
            
            for elf in elves:
                [x,y] = dict(elf)['pos']
                grid[y][x] = '#'#str(elf['elf_no'])
            
            grid[0][0] = '0'
            for row in grid:
                rowstr = ''
                for r in row:
                    rowstr += r
                print(rowstr)
    
    return (max(xs) - min(xs)+1) * ((max(ys) - min(ys)+1)) - len(elves)

def Part2(filename, printing = False):
    """
    Solutions to Part 2
    """
    elves = ParseData(filename)

    stable = 0
    r = 0
    while stable < len(elves):
        stable = 0
        if printing:            
            print()
        # propose next postions
        for e in range(0,len(elves)):

            elf = dict(elves[e])

            [x,y] = elf['pos']

            N_set  = [[x  ,y-1],[x+1,y-1],[x-1,y-1]]
            S_set  = [[x  ,y+1],[x+1,y+1],[x-1,y+1]]
            E_set  = [[x+1,y  ],[x+1,y-1],[x+1,y+1]]
            W_set  = [[x-1,y  ],[x-1,y-1],[x-1,y+1]]

            next_pos_sets = [N_set,S_set,W_set,E_set]
            
            proposed = False
            while not proposed:
                current_elf_positions =  [elf['pos'] for elf in elves]  

                # if no other elf nearbye do nothing
                if not any((epos in current_elf_positions) for epos in N_set+S_set+E_set+W_set):
                    elf['next_pos'] = elf['pos']
                    elves[e] = elf
                    proposed = True
                    stable += 1 
                    break

                # check if surrounded at all sides:
                elif all(any((epos in current_elf_positions) for epos in set) for set in next_pos_sets):
                    elf['next_pos'] = elf['pos']
                    elves[e] = elf
                    proposed = True
                    stable += 1 
                    break

                # propse
                else:  
                    for d in range(0,4):
                        next_pos_set = next_pos_sets[(d+r)%4]  
                        if ((next_pos_set[0] not in current_elf_positions) and
                        (next_pos_set[1] not in current_elf_positions) and
                        (next_pos_set[2] not in current_elf_positions)):  
                            elf['next_pos'] = next_pos_set[0]
                            elf['last_dir'] = d
                            elves[e] = elf
                            proposed = True
                            break
            if printing:
                print(r+1,elf,list('nswe')[elf['last_dir']%4])
            
        # evaluate proposed positions
        proposed_elf_positions =  [elf['next_pos'] for elf in elves]

        for e in range(0,len(elves)):
            elf = dict(elves[e])
            # dups
            if proposed_elf_positions.count(elf['next_pos']) > 1:
                elf['next_pos'] = elf['pos']
            else:
                elf['pos'] = elf['next_pos']            
            elves[e] = elf

        xs = []
        ys = []
        
        for e in range(0,len(elves)):
            elf = dict(elves[e])
            xs.append(elf['pos'][0])
            ys.append(elf['pos'][1])

        if printing:   
            print()                   
            grid = []
            for y in range(-max(ys),max(ys)+1):
                row = []
                for x in range(-max(xs),max(xs)+1):
                    row.append('.')
                grid.append(row)
            
            for elf in elves:
                [x,y] = dict(elf)['pos']
                grid[y][x] = '#'#str(elf['elf_no'])
            
            grid[0][0] = '0'
            for row in grid:
                rowstr = ''
                for r in row:
                    rowstr += r
                print(rowstr)
        r += 1
    
    return r

def main():
    """
    main function
    """ 
    testfile  = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day23\n")

    print(f'Tests : Answer to Part 1 = {Part1(testfile, printing = False)}')    
    print(f'Tests : Answer to Part 2 = {Part2(testfile, printing = False)}\n')

    assert(Part1(testfile) == 110)
    assert(Part2(testfile) == 20)

    print(f'Inputs: Answer to Part 1 = {Part1(inputfile, printing = False)}')
    print(f'Inputs: Answer to Part 2 = {Part2(inputfile, printing = False)}')

if __name__ == "__main__":
    main()