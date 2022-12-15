import numpy as np

def ParseData(file):    
    """
    function to read and parse in data from a txt file
    (rather messy)
    """
    with open(file) as f:
        data = f.read().split('\n')

    sensors = []
    beacons = []
    for i in range(0,len(data)):
        sensors_string = data[i].split(':')[0]
        beacons_string = data[i].split(':')[1]
        sensor_coord_string = (sensors_string.split(' '))[-2:]
        beacon_coord_string = (beacons_string.split(' '))[-2:]
        sensor_coord = [int(sensor_coord_string[0].replace(',','').replace('x=','')),
                        int(sensor_coord_string[1].replace('y=',''))]
        beacon_coord = [int(beacon_coord_string[0].replace(',','').replace('x=','')),
                        int(beacon_coord_string[1].replace('y=',''))]
        sensors.append(sensor_coord)
        beacons.append(beacon_coord)

    return sensors,beacons

def Part1(filename, row_to_check, printing = False):
    """
    Solutions to Part 1
    """
    sensors, beacons = ParseData(filename)

    ranges = []

    for i in range(0,len(sensors)):        

        sensor = sensors[i]
        beacon = beacons[i]

        manhatten_distance = abs(sensor[0] - beacon[0]) + abs(sensor[1] - beacon[1])

        if printing:
            stri = str(i+1)
            if i < 9:
                stri = '0'+str(i+1)
            print(f'scanning sensor {stri} of {len(sensors)} - dist = {manhatten_distance}')

        # if row to check is relevant
        if (row_to_check >= sensor[1]-manhatten_distance and row_to_check <= sensor[1]+manhatten_distance):

            y = abs(sensor[1] - row_to_check)
            # find first x ind below man hat dist
            x_start = 'null'
            for x in range(sensor[0]-manhatten_distance, sensor[0]+manhatten_distance):
                if ((abs(sensor[0] - x) + y) <= manhatten_distance):
                    x_start = x
                    break
            
            assert(x_start != 'null')

            # find last x ind below man han dist
            x_end = 'null'
            for x in range(sensor[0]+manhatten_distance, sensor[0]-manhatten_distance,-1):
                if ((abs(sensor[0] - x) + y) <= manhatten_distance):
                    x_end = x
                    break

            assert(x_end != 'null')
            ranges.append([x_start,x_end])    

    xs = []
    for r in ranges:
        xs.append(r[0])
        xs.append(r[1])
    return max(xs)-min(xs)

def Part2(filename, max_coord, printing = False):
    """
    Solutions to Part 2
    """
    sensors, beacons = ParseData(filename)

    for row_to_check in range(0,max_coord+1):

        ranges = []

        for i in range(0,len(sensors)):        

            sensor = sensors[i]
            beacon = beacons[i]

            manhatten_distance = abs(sensor[0] - beacon[0]) + abs(sensor[1] - beacon[1])

            if printing:
                stri = str(i+1)
                if i < 9:
                    stri = '0'+str(i+1)
                print(f'scanning sensor {stri} of {len(sensors)} - dist = {manhatten_distance}')

            # if row to check is relevant
            if (row_to_check >= sensor[1]-manhatten_distance and row_to_check <= sensor[1]+manhatten_distance):

                y = abs(sensor[1] - row_to_check)
                # find first x ind below man hat dist
                for x in range(sensor[0]-manhatten_distance, sensor[0]+manhatten_distance):
                    if ((abs(sensor[0] - x) + y) <= manhatten_distance):
                        x_start = x
                        break
                
                # find last x ind below man han dist
                for x in range(sensor[0]+manhatten_distance, sensor[0]-manhatten_distance,-1):
                    if ((abs(sensor[0] - x) + y) <= manhatten_distance):
                        x_end = x
                        break

                ranges.append([x_start,x_end])    

        ranges.sort()
            
        for r in range(0,len(ranges)-1):
            if ranges[r][1] + 2 == ranges[r+1][0]:

                find = ranges[r][1] + 1
                nested = False
                for range_check in ranges:
                    if (range_check[0] <= find) and (find <= range_check[1]): 
                        nested = True
                        break
                    
                if nested == False:
                    return (find * 4000000) + row_to_check  
    return 0

def main():
    """
    main function
    """ 
    testfile  = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day15")

    print(f'"{testfile}": Answer to Part 1 = {Part1(testfile, 10, printing = False)}')    
    print(f'"{testfile}": Answer to Part 2 = {Part2(testfile, 20, printing = False)}\n')

    assert(Part1(testfile,10) == 26)
    assert(Part2(testfile,20) == 56000011)

    print(f'"{inputfile}": Answer to Part 1 = {Part1(inputfile, 2000000, printing = False)}')
    print(f'"{inputfile}": Answer to Part 2 = {Part2(inputfile, 4000000, printing = False)}')

if __name__ == "__main__":
    main()