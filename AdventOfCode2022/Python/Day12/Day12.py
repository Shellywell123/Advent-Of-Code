altitude = 'abcdefghijklmnopqrstuvwxyz'

def ParseData(file):    
    """
    function to read and parse in data from a txt file
    """
    with open(file) as f:
        data = f.read().split('\n')

    # find starting pos and ending pos
    rows = []
    for y in range(0,len(data)):
        row = []
        for x in range(0,len(data[0])):
            
            if data[y][x] == "S":
                init_pos = [x,y]

            if data[y][x] == "E":
                final_pos = [x,y]
            
            row.append({"altitude":data[y][x], "dist_from_E":''})
        rows.append(row)
    
    rows[init_pos[1]][init_pos[0]] = {"altitude":'a', "dist_from_E":''}
    rows[final_pos[1]][final_pos[0]] = {"altitude":'z', "dist_from_E":0}

    return rows, init_pos, final_pos

def Part1(filename):
    """
    Solutions to Part 1
    """

    rows, init_pos, final_pos = ParseData(filename)
    paths = rows
    i = 0
    
    while type(paths[init_pos[1]][init_pos[0]]["dist_from_E"]) != int:

        for y in range(0,len(rows)):
            for x in range(0,len(rows[0])):

                if paths[y][x]["dist_from_E"] == i:
    
                    current_square = paths[y][x]

                    coord_left  = [x-1,y  ]
                    coord_right = [x+1,y  ]
                    coord_up    = [x  ,y-1]
                    coord_down  = [x  ,y+1]

                    for [x_check,y_check] in [coord_down,coord_up,coord_left,coord_right]:
                    
                        # ignore if out of bounds
                        if len(paths[0]) -1 < x_check or x_check <0 or len(paths) -1 < y_check or y_check < 0 :
                            continue
                        
                        #ignore if already int
                        if type(paths[y_check][x_check]["dist_from_E"]) == int and paths[y_check][x_check]["dist_from_E"] < current_square["dist_from_E"]+1:
                            continue

                        next_square = paths[y_check][x_check]
                        current_altitude = altitude.index(current_square["altitude"])
                        next_altitude    = altitude.index(next_square["altitude"])

                        if next_altitude + 2 > current_altitude :
                            paths[y_check][x_check]["dist_from_E"] = current_square["dist_from_E"]+1                        
        
        i += 1
    
    elements = []
    for row in paths:
        for square in row:
            if type(square["dist_from_E"])== int:
                elements.append(square["dist_from_E"])

    return max(elements)

def Part2(filename):
    """
    Solutions to Part 2
    """
    
    rows, init_pos, final_pos = ParseData(filename)
    paths = rows
    i = 0

    while type(paths[init_pos[1]][init_pos[0]]["dist_from_E"]) != int:

        for y in range(0,len(rows)):
            for x in range(0,len(rows[0])):
                
                if paths[y][x]["dist_from_E"] == i:
    
                    current_square = paths[y][x]

                    coord_left  = [x-1,y  ]
                    coord_right = [x+1,y  ]
                    coord_up    = [x  ,y-1]
                    coord_down  = [x  ,y+1]

                    for [x_check,y_check] in [coord_down,coord_up,coord_left,coord_right]:
                    
                        # ignore if out of bounds
                        if len(paths[0]) -1 < x_check or x_check <0 or len(paths) -1 < y_check or y_check < 0 :
                            continue
                        
                        #ignore if already int
                        if type(paths[y_check][x_check]["dist_from_E"]) == int and paths[y_check][x_check]["dist_from_E"] < current_square["dist_from_E"]+1:
                            continue

                        next_square = paths[y_check][x_check]
                        current_altitude = altitude.index(current_square["altitude"])
                        next_altitude    = altitude.index(next_square["altitude"])

                        if next_altitude + 2 > current_altitude :
                            paths[y_check][x_check]["dist_from_E"] = current_square["dist_from_E"]+1
        
        i += 1
    
    elements = []
    for row in paths:
        for square in row:
            if type(square["dist_from_E"]) == int and square["altitude"] == 'a':
                elements.append(square["dist_from_E"])

    return min(elements)

def main():
    """
    main function
    """ 
    testfile  = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day12")

    print(f'"{testfile}": Answer to Part 1 = {Part1(testfile)}')
    print(f'"{testfile}": Answer to Part 2 = {Part2(testfile)}\n')

    print(f'"{inputfile}": Answer to Part 1 = {Part1(inputfile)}')
    print(f'"{inputfile}": Answer to Part 2 = {Part2(inputfile)}')

if __name__ == "__main__":
    main()