def ParseData(file):    
    """
    function to read and parse in data from a txt file
    (rather messy)
    """
    with open(file) as f:
        data = f.read().split('\n')

    scan_traces = data

    trace_x_positions = []
    trace_y_positions = []
    for scan_trace in scan_traces:
        for pos in scan_trace.split(' -> '):
            trace_x_positions.append(int((pos.split(','))[0]))
            trace_y_positions.append(int((pos.split(','))[1]))
    
    assert(len(trace_x_positions)==len(trace_y_positions))

    x_start = min(trace_x_positions)
    y_start = min(trace_y_positions)  

    x_end = max(trace_x_positions)
    y_end = max(trace_y_positions)

    rock_formations = []
    for y in range(0,y_end+4):
        row = []
        for x in range (0,1000):
            row.append('.')
        rock_formations.append(row)
    
    rock_formations[0][500-x_start+1] = '+'
    
    for scan_trace in scan_traces:
        positions = scan_trace.split(' -> ')

        for i in range(0,len(positions)-1):
            x_position = int(positions[i].split(',')[0])
            y_position = int(positions[i].split(',')[1])

            x_next_position = int(positions[i+1].split(',')[0])
            y_next_position = int(positions[i+1].split(',')[1])
            
            # right
            if x_next_position - x_position > 0:
                for x in range(x_position, x_next_position+1):
                    rock_formations[y_position][x] = '#'

            # left
            if x_next_position - x_position < 0:
                for x in range(x_position, x_next_position-1,-1):
                    rock_formations[y_position][x] = '#'

            # down
            if y_next_position - y_position > 0:
                for y in range(y_position, y_next_position+1):
                    rock_formations[y][x_position] = '#'
            
            # up
            if y_next_position - y_position < 0:
                for y in range(y_position, y_next_position-1,-1):
                    rock_formations[y][x_position] = '#'

    return x_start,x_end,y_start,y_end,rock_formations

def Sand(rock_formations,x,y):
        """
        Recursive sand physics
        """
        at_rest = False

        while not at_rest:

            # sand on air 
            if rock_formations[y+1][x] =='.':
                y=y+1
                continue

            # sand on sand/rock
            if rock_formations[y+1][x] in ['#','o']:
                # check down left
                if rock_formations[y+1][x-1] == '.':
                    x=x-1
                    y=y+1
                    continue
                # check down  right
                elif rock_formations[y+1][x+1] == '.':
                    x = x+1
                    y = y+1
                    continue
                # sand at rest
                else:
                    rock_formations[y][x] = 'o'
                    at_rest = True

        return rock_formations      

def Part1(filename, printing = False):
    """
    Solutions to Part 1
    """

    x_start,x_end,y_start,y_end,rock_formations = ParseData(filename)

    ans = 0
    for i in range(0,10000000): # brute force lol
        ans = i
        try:
            rock_formation = Sand(rock_formations,500,0)
        except:            
            break

    if printing == True:
        
        for i in range(0,len(rock_formations)):
            form = rock_formations[i]
            row = ''
            for j in range(0,len(form)):
                if j > 450 and j <515:
                    row += form[j]
            istr = str(i)
            if i < 100:
                istr = '0'+str(i)
            if i < 10:
                istr = '00'+str(i)
            
            print(istr,row)

    return ans

def Part2(filename, printing = False):
    """
    Solutions to Part 2
    """
    x_start,x_end,y_start,y_end,rock_formations = ParseData(filename)

    # add infinte floor
    for r in range(0,len(rock_formations[0])):
        rock_formations[y_end+2][r] = '#'

    ans = 0
    while rock_formations[0][500] != 'o':
        ans += 1
        try:
            rock_formation = Sand(rock_formations,500,0)
        except:
            break

    if printing == True:
        
        for i in range(0,len(rock_formations)):
            form = rock_formations[i]
            row = ''
            for j in range(0,len(form)):
                if j > 450 and j <515:
                    row += form[j]
            istr = str(i)
            if i < 100:
                istr = '0'+str(i)
            if i < 10:
                istr = '00'+str(i)
            
            print(istr,row)

    return ans

def main():
    """
    main function
    """ 
    testfile  = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day14")

    print(f'"{testfile}": Answer to Part 1 = {Part1(testfile, printing = False)}')    
    print(f'"{testfile}": Answer to Part 2 = {Part2(testfile, printing = False)}\n')

    assert(Part1(testfile) == 24)
    assert(Part2(testfile) == 93)

    print(f'"{inputfile}": Answer to Part 1 = {Part1(inputfile, printing = False)}')
    print(f'"{inputfile}": Answer to Part 2 = {Part2(inputfile, printing = False)}')

if __name__ == "__main__":
    main()