def ParseData(file):    
    """
    function to read and parse in data from a txt file
    """
    with open(file) as f:
        data = f.read().split('\n')

    pairs = []
    pair = []
    for i in range(0,len(data)):
        line = data[i]
        
        if line.strip() == '':
            continue
        else:
            pair.append(eval(line))
            if len(pair) == 2:
                pairs.append(pair)
                pair = []

    return pairs

def Compare(left, right, printing):
    """
    """
    if printing == True:
        print(f'compare {left} vs {right}')
    
    if left == right:
        return None
    
    if type(left) == int:
        if type(right) == int:
            
            if left < right:
                if printing == True:
                    print('>--- Left side is smaller, so inputs are in the right order')
                return True
            
            if left > right:
                if printing == True:
                    print('>--- Right side is smaller, so inputs are not in the right order')
                return False

        else:
            return Compare([left],right,printing)
    
    else:
        if type(right) == int:
            return Compare(left,[right],printing)
        else:
            # recursive compare elements
            for j in range(0, min([len(left),len(right)])):

                result = Compare(left[j],right[j],printing)
                if result != None:
                    return result
            
            if len(left) < len(right):
                if printing == True:
                    print('>- Left side ran out of items, so inputs are in the right order')
                return True
            
            if len(left) > len(right):
                if printing == True:
                    print('>- Right side ran out of items, so inputs are not in the right order')
                return False

def Part1(filename, printing = False):
    """
    Solutions to Part 1
    """

    pairs = ParseData(filename)

    correct_order_indicies = []

    for i in range(0,len(pairs)):

        if printing == True:
            print(f'\nPair {i+1}')
        
        left,right = pairs[i][0],pairs[i][1]

        if Compare(left,right,printing):
            correct_order_indicies.append(i+1)

    if printing == True:
        print(correct_order_indicies)
    return sum(correct_order_indicies)

def Part2(filename, printing = False):
    """
    Solutions to Part 2
    """
    
    pairs = ParseData(filename)

    def sort(pairs):
        """
        """
        sorts = [False]*(len(pairs) -1)
        for z in range(0,1000): # brute forced instead of quick sort as I am tired
            for i in range(0,len(pairs)-1):
                temp1 = pairs[i]
                temp2 = pairs[i+1]
                if Compare(pairs[i],pairs[i+1], printing) == False:
                    pairs[i]   = temp2
                    pairs[i+1] = temp1
                    sorts[i] = True
        return pairs

    new_pairs = []
    for pair in pairs:
        
        new_pairs.append(pair[0])
        new_pairs.append(pair[1])
    
    new_pairs.append([[2]])
    new_pairs.append([[6]])

    pairs = sort(new_pairs)

    if printing == True:
        for pair in pairs:
            print(pair)
        print(pairs.index([[2]]) , pairs.index([[6]]))

    return (pairs.index([[2]])+1) * (pairs.index([[6]])+1)

def main():
    """
    main function
    """ 
    testfile  = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day13")

    print(f'"{testfile}": Answer to Part 1 = {Part1(testfile, printing = False)}')    
    print(f'"{testfile}": Answer to Part 2 = {Part2(testfile, printing = False)}\n')

    assert(Part1(testfile) == 13)
    assert(Part2(testfile) == 140)

    print(f'"{inputfile}": Answer to Part 1 = {Part1(inputfile, printing = False)}')
    print(f'"{inputfile}": Answer to Part 2 = {Part2(inputfile, printing = False)}')

if __name__ == "__main__":
    main()