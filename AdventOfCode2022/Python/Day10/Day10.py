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

    cycle_number = 1
    X = 1
    key_cycles = []

    for entry in data:
        instruction = entry.split(' ')[0]
        
        if instruction == 'noop':

            cycle_number += 1
            for key_cycle in [20, 60, 100, 140, 180, 220]:
                if cycle_number == key_cycle:
                    key_cycles.append(X*cycle_number)
            
        if instruction == 'addx':

            value = int(entry.split(' ')[1])

            cycle_number += 1
            for key_cycle in [20, 60, 100, 140, 180, 220]:
                if cycle_number == key_cycle:
                    key_cycles.append(X*cycle_number)
            
            X += value
            
            cycle_number += 1
            for key_cycle in [20, 60, 100, 140, 180, 220]:
                if cycle_number == key_cycle:
                    key_cycles.append(X*cycle_number)
    
    return sum(key_cycles)

def Part2(filename):
    """
    Solutions to Part 2
    """
    data = get_inputs(filename)
    
    cycle_number = 1
    X = 1
    pixels = list('.'*40*6)

    for entry in data:

        instruction = entry.split(' ')[0]

        if instruction == 'noop':
            
            for pos in [X-1, X, X+1]:
                for row in [0,1,2,3,4,5]:
                   if ((40*row)+1 <= cycle_number <= 40*(row+1)) and (cycle_number-1 - (row*40) == pos):
                        pixels[cycle_number-1] = '#'
                
            cycle_number += 1

        if instruction == 'addx':

            value = int(entry.split(' ')[1])

            for pos in [X-1, X, X+1]:
                for row in [0,1,2,3,4,5]:
                    if ((40*row)+1 <= cycle_number <= 40*(row+1)) and (cycle_number-1 - (row*40) == pos):
                        pixels[cycle_number-1] = '#'

            cycle_number += 1

            for pos in [X-1, X, X+1]:
                for row in [0,1,2,3,4,5]:
                    if ((40*row)+1 <= cycle_number <= 40*(row+1)) and (cycle_number-1 - (row*40) == pos):
                        pixels[cycle_number-1] = '#'

            cycle_number += 1
            X += value
                
    str_ = ''
    for pixel in pixels:
        str_ += pixel

    print('.'+'-'*41+'.')
    for i in range(0,6):
        print('| '+str_[i*40:(i+1)*40-1] +' |')
    print('.'+'-'*41+'.')
    
    return "read above output"

def main():
    """
    main function
    """ 
    testfile = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day10")

    print(f'"{testfile}": Answer to Part 1 = {Part1(testfile)}')
    print(f'"{testfile}": Answer to Part 2 = {Part2(testfile)}')

    print(f'"{inputfile}": Answer to Part 1 = {Part1(inputfile)}')
    print(f'"{inputfile}": Answer to Part 2 = {Part2(inputfile)}')

if __name__ == "__main__":
    main()