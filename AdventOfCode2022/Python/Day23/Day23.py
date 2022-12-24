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

    for r in range(0,3):
        print()
        for e in range(0,len(elves)):

            elf = dict(elves[e])

            [x,y] = elf['pos']

            N  = [x  ,y-1]
            NE = [x+1,y-1]
            NW = [x-1,y-1]
            S  = [x  ,y+1]
            SE = [x+1,y+1]
            SW = [x-1,y+1]
            E  = [x+1,y  ]
            W  = [x-1,y  ]

            next_posss = [[N,NE,NW],[S,SE,SW],[W,SW,NW],[E,SE,NE]]
            #next_posss = next_posss[(elf['last_dir']+1)%len(next_posss):] + next_posss[:(elf['last_dir']+1)%len(next_posss)]
            #print(next_posss)
            
            # make all decisions 
            set = False
            for d in range(elf['last_dir']+1,elf['last_dir']+5):
                next_poss = next_posss[d%4]
                for next_pos in next_poss:
                    if not set:                        
                        if next_pos not in [elf['pos'] for elf in elves]:
                            elf['next_pos'] = next_poss[0]
                            elf['next_pos'] = elf['pos']
                            elf['last_dir'] = d
                            elves[e] = elf
                            set = True
                    
            print(r,elf,list('nswe')[elf['last_dir']%4])


        
        for e in range(0,len(elves)):
            elf = dict(elves[e])
            if next_pos not in [elf2['pos'] for elf2 in elves if elf2['elf_no'] != elf['elf_no']]:
                elf['pos'] = elf['next_pos']
                elves[e] = elf

        xs = []
        ys = []
        print()
        for e in range(0,len(elves)):
            elf = dict(elves[e])
            xs.append(elf['pos'][0])
            ys.append(elf['pos'][1])
        
        grid = []
        for y in range(-max(ys),max(ys)+1):
            row = []
            for x in range(-max(xs),max(xs)+1):
                row.append('.')
            grid.append(row)
        
        for row in grid:
            print(row)
        
        for elf in elves:
            [x,y] = dict(elf)['pos']
            grid[y][x] = '#'
        
        for row in grid:
            print(row)
        

        
    ans = (max(xs) - min(xs)) * ((max(ys) - min(ys))) - len(elves)
            
    return ans

def Part2(filename, printing = False):
    """
    Solutions to Part 2
    """
    encrypted_file = ParseData(filename)

    return 0

def main():
    """
    main function
    """ 
    testfile  = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day18")

    print(f'Tests : Answer to Part 1 = {Part1(testfile, printing = True)}')    
    # print(f'Tests : Answer to Part 2 = {Part2(testfile, printing = False)}\n')

    # assert(Part1(testfile) == 3)
    # assert(Part2(testfile) == 1623178306)

    # print(f'Inputs: Answer to Part 1 = {Part1(inputfile, printing = False)}')
    # print(f'Inputs: Answer to Part 2 = {Part2(inputfile, printing = False)}')

if __name__ == "__main__":
    main()