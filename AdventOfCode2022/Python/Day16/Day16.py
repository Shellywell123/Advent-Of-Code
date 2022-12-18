import numpy as np

def ParseData(file):    
    """
    function to read and parse in data from a txt file
    (rather messy)
    """
    with open(file) as f:
        data = f.read().split('\n')

    valves = {}
    for i in range(0,len(data)):
        line_string = data[i].split(':')[0]
        name    = line_string.split(' ')[1]
        rate    = int((line_string.split('=')[1]).split(';')[0])
        tunnels = (line_string[line_string.index('to valve')+9:]).strip().split(', ')
        valves[name] = {"Rate": rate, "Tunnels": tunnels}
    return valves

def Part1(filename, printing = False):
    """
    Solutions to Part 1
    """
    valves = ParseData(filename)

    # for valve in valves:
    print(valves)
    
    open_valves = []
    current_valve = 'AA'
    released = 0
    for minute in range(1,30 +1):

        moving = False
        opening = False

        if valves[current_valve]["Rate"] == 0:
            # dont add to open valves
            moving = True
        else:
            opening = True

        if moving:
            nvr = 0
            for nv in valves[current_valve]["Tunnels"]:
                if valves[nv]["Rate"] > nvr:
                    nvr = valves[nv]["Rate"]
                    next_valve = nv

        if opening:
            pass

        pressure = 0
        if printing:
            print(f'== Minute {minute} == (at valve {current_valve})')
            print(f'open valves = {open_valves}, releasing {pressure} pressure')
            if opening:
                print(f'You open valve {current_valve}\n')
            if moving:
                print(f'You move to valve {next_valve}\n')
        if moving:
            current_valve = next_valve
        released += pressure   

    
    return 0

def Matrix(filename):
    """
    """

    valves = ParseData(filename)

    distances = {}
    for valve in valves:
        print(valve)


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

    print("Advent-Of-Code 2022 - Day16")

    #print(f'"{testfile}": Answer to Part 1 = {Part1(testfile, printing = True)}')    
    #print(f'"{testfile}": Answer to Part 2 = {Part2(testfile, 20, printing = False)}\n')
    Matrix(testfile)
    #assert(Part1(testfile) == 1651)
    #assert(Part2(testfile,20) == 56000011)

    #print(f'"{inputfile}": Answer to Part 1 = {Part1(inputfile, 2000000, printing = False)}')
    #print(f'"{inputfile}": Answer to Part 2 = {Part2(inputfile, 4000000, printing = False)}')

if __name__ == "__main__":
    main()