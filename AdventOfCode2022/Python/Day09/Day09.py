def get_inputs(file):
    """
    function to read in data from a txt file
    """

    with open(file) as f:
        data = f.read().split('\n')

    return data

def Part1(filename):
    """
    Solutions to Part 1
    """
    data = get_inputs(filename)

    H_position = [0,0]
    T_position = [0,0]

    T_positions = [str(T_position[0])+'-'+str(T_position[1])]

    for move in data:

        direction = move.split(' ')[0]
        distance  = int(move.split(' ')[1])

        assert(abs(H_position[0] - T_position[0]) < 2)
        assert(abs(H_position[1] - T_position[1]) < 2)

        for i in range(0,distance):

            if direction == "U":
                H_position[1] += 1

            if direction == "D":
                H_position[1] -= 1

            if direction == "L":
                H_position[0] -= 1

            if direction == "R":
                H_position[0] += 1

            if not ((H_position[0] == T_position[0]) & (H_position[1] == T_position[1])):
     
                # move T up
                if (H_position[0] - T_position[0] == 0) & (H_position[1] - T_position[1] > 1):
                    T_position[1] += 1
                
                # move T down
                if (H_position[0] - T_position[0] == 0) & (H_position[1] - T_position[1] < -1):
                    T_position[1] -= 1

                # move T right
                if (H_position[1] - T_position[1] == 0) & (H_position[0] - T_position[0] >  1):
                    T_position[0] += 1
                
                # move T left
                if (H_position[1] - T_position[1] == 0) & (H_position[0] - T_position[0] < -1):
                    T_position[0] -= 1
                
                # move T up right
                if ((H_position[1] - T_position[1] >= 1) & (H_position[0] - T_position[0] > 1) or
                    (H_position[1] - T_position[1] > 1)  & (H_position[0] - T_position[0] >= 1)):
                    T_position[0] += 1
                    T_position[1] += 1
                
                # move T down right
                if ((H_position[1] - T_position[1] <= -1) & (H_position[0] - T_position[0] > 1) or
                    (H_position[1] - T_position[1] < -1)  & (H_position[0] - T_position[0] >= 1)):
                    T_position[0] += 1
                    T_position[1] -= 1
                
                # move T up left
                if ((H_position[1] - T_position[1] >= 1) & (H_position[0] - T_position[0] < -1) or
                    (H_position[1] - T_position[1] > 1)  & (H_position[0] - T_position[0] <= -1)):
                    T_position[0] -= 1
                    T_position[1] += 1
                
                # move T down left
                if ((H_position[1] - T_position[1] <= -1) & (H_position[0] - T_position[0] < -1) or
                    (H_position[1] - T_position[1] < -1)  & (H_position[0] - T_position[0] <= -1)):
                    T_position[0] -= 1
                    T_position[1] -= 1    

                if str(T_position[0])+'-'+str(T_position[1]) not in T_positions:
                    T_positions.append(str(T_position[0])+'-'+str(T_position[1]))
    
    return len(T_positions)

def Part2(filename):
    """
    Solutions to Part 2
    """
    data = get_inputs(filename)
    
    knot_positions = [[0,0],[0,0],[0,0],[0,0],[0,0],[0,0],[0,0],[0,0],[0,0],[0,0]]

    T_positions = ['0-0']

    for move in data:

        direction = move.split(' ')[0]
        distance  = int(move.split(' ')[1])

        for i in range(0,distance):

            if direction == "U":
                knot_positions[0][1] += 1

            if direction == "D":
                knot_positions[0][1] -= 1

            if direction == "L":
                knot_positions[0][0] -= 1

            if direction == "R":
                knot_positions[0][0] += 1

            for i in range(1,len(knot_positions)):

                if not ((knot_positions[i-1][0] == knot_positions[i][0]) & (knot_positions[i-1][1] == knot_positions[i][1])):
       
                    # move T up
                    if (knot_positions[i-1][0] - knot_positions[i][0] == 0) & (knot_positions[i-1][1] - knot_positions[i][1] > 1):
                        knot_positions[i][1] += 1
                    
                    # move T down
                    if (knot_positions[i-1][0] - knot_positions[i][0] == 0) & (knot_positions[i-1][1] - knot_positions[i][1] < -1):
                        knot_positions[i][1] -= 1

                    # move T right
                    if (knot_positions[i-1][1] - knot_positions[i][1] == 0) & (knot_positions[i-1][0] - knot_positions[i][0] >  1):
                        knot_positions[i][0] += 1
                    
                    # move T left
                    if (knot_positions[i-1][1] - knot_positions[i][1] == 0) & (knot_positions[i-1][0] - knot_positions[i][0] < -1):
                        knot_positions[i][0] -= 1
                    
                    # move T up right
                    if ((knot_positions[i-1][1] - knot_positions[i][1] >= 1) & (knot_positions[i-1][0] - knot_positions[i][0] > 1) or
                        (knot_positions[i-1][1] - knot_positions[i][1] > 1)  & (knot_positions[i-1][0] - knot_positions[i][0] >= 1)):
                        knot_positions[i][0] += 1
                        knot_positions[i][1] += 1
                    
                    # move T down right
                    if ((knot_positions[i-1][1] - knot_positions[i][1] <= -1) & (knot_positions[i-1][0] - knot_positions[i][0] > 1) or
                        (knot_positions[i-1][1] - knot_positions[i][1] < -1)  & (knot_positions[i-1][0] - knot_positions[i][0] >= 1)):
                        knot_positions[i][0] += 1
                        knot_positions[i][1] -= 1
                    
                    # move T up left
                    if ((knot_positions[i-1][1] - knot_positions[i][1] >= 1) & (knot_positions[i-1][0] - knot_positions[i][0] < -1) or
                        (knot_positions[i-1][1] - knot_positions[i][1] > 1)  & (knot_positions[i-1][0] - knot_positions[i][0] <= -1)):
                        knot_positions[i][0] -= 1
                        knot_positions[i][1] += 1
                    
                    # move T down left
                    if ((knot_positions[i-1][1] - knot_positions[i][1] <= -1) & (knot_positions[i-1][0] - knot_positions[i][0] < -1) or
                        (knot_positions[i-1][1] - knot_positions[i][1] < -1)  & (knot_positions[i-1][0] - knot_positions[i][0] <= -1)):
                        knot_positions[i][0] -= 1
                        knot_positions[i][1] -= 1  

                    if i == len(knot_positions)-1:
                        if str(knot_positions[i][0])+'-'+str(knot_positions[i][1]) not in T_positions:
                            T_positions.append(str(knot_positions[i][0])+'-'+str(knot_positions[i][1]))
                                
            for a in range(1,len(knot_positions)):
                assert(abs(knot_positions[a][0] - knot_positions[a-1][0]) < 2)
                assert(abs(knot_positions[a][1] - knot_positions[a-1][1]) < 2)

    return len(T_positions)

def main():
    """
    main function
    """ 
    testfile = "tests.txt"
    testfile2 = "tests2.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day08")

    print(f'"{testfile}": Answer to Part 1 = {Part1(testfile)}')
    print(f'"{testfile}": Answer to Part 2 = {Part2(testfile)}')

    print(f'"{testfile2}": Answer to Part 1 = {Part1(testfile2)}')
    print(f'"{testfile2}": Answer to Part 2 = {Part2(testfile2)}')

    print(f'"{inputfile}": Answer to Part 1 = {Part1(inputfile)}')
    print(f'"{inputfile}": Answer to Part 2 = {Part2(inputfile)}')

if __name__ == "__main__":
    main()