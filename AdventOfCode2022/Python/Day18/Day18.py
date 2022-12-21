def ParseData(file):    
    """
    function to read and parse in data from a txt file
    """
    with open(file) as f:
        data = f.read().split('\n')
    
    coords = []
    for d in data:
        coord = []
        for num in d.split(','):
            coord.append(int(num))
        coords.append(coord)

    return coords

def Solve(coords, printing = False):
    """
    Count number of faces
    """
    touches = 0
    for coord in coords:
        for coord_check in coords:

            if coord == coord_check:
                continue

            if (((coord[0] == coord_check[0]) and
                (coord[1] == coord_check[1]) and
                (abs(coord[2] - coord_check[2]) == 1))
                or
                ((coord[0] == coord_check[0]) and
                (coord[2] == coord_check[2]) and
                (abs(coord[1] - coord_check[1]) == 1))
                or
                ((coord[2] == coord_check[2]) and
                (coord[1] == coord_check[1]) and
                (abs(coord[0] - coord_check[0]) == 1))):
                touches += 1

    return 6 * len(coords) - (touches)

def Part1(filename, printing = False):
    """
    Solutions to Part 1
    """
    return Solve(ParseData(filename))

def Part2(filename, printing = False):
    """
    Solutions to Part 2
    """
    coords = ParseData(filename)

    xs = []
    ys = []
    zs = []

    for coord in coords:
        xs.append(coord[0])
        ys.append(coord[1])
        zs.append(coord[2])

    def canEscape(x,y,z,path):
        """
        Recurse
        """
        path.append([x,y,z])
        if [x,y,z] in coords:
            return False # hit rock

        elif (
            (x == max(xs)+1 or x == min(xs)-1) or
            (y == max(ys)+1 or y == min(ys)-1) or
            (z == max(zs)+1 or z == min(zs)-1)):
            return True # escaped
        
        # try every grid direction 
        for [xn, yn, zn] in [[x-1,y,z],[x+1,y,z],[x,y-1,z],[x,y+1,z],[x,y,z-1],[x,y,z+1]]:
            if [xn,yn,zn] not in path:
                if canEscape(xn,yn,zn,path):
                    return True 
        
        return False
    
    import sys
    # increase pythons recurssion lim 
    sys.setrecursionlimit(2000) 

    inside = []
    for x in range(min(xs),max(xs)):
        for y in range(min(ys),max(ys)):
            for z in range(min(zs),max(zs)):

                if (([x,y,z] not in coords) and 
                    ([x,y,z] not in inside)):
                    
                    if not canEscape(x,y,z,[]):
                        inside.append([x,y,z])

    return Solve(coords + inside)

def main():
    """
    main function
    """ 
    testfile  = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day18")

    print(f'Tests : Answer to Part 1 = {Part1(testfile, printing = False)}')    
    print(f'Tests : Answer to Part 2 = {Part2(testfile, printing = False)}\n')

    assert(Part1(testfile) == 64)
    assert(Part2(testfile) == 58)

    print(f'Inputs: Answer to Part 1 = {Part1(inputfile, printing = False)}')
    print(f'Inputs: Answer to Part 2 = {Part2(inputfile, printing = False)}')

if __name__ == "__main__":
    main()