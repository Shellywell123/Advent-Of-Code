def ParseData(file):    
    """
    function to read and parse in data from a txt file
    """
    with open(file) as f:
        data = f.read().split('\n')
    
    encrypted_file = []
    for num in data:
        encrypted_file.append(int(num))

    return encrypted_file

def Solve(result,encrypted_file):
    """
    """
    for num in encrypted_file:
        ind = result.index(num)

        # # rotate so list starts with num
        result = result[result.index(num):] + result[:result.index(num)]

        # # remve num from list
        result = result[1:]

        # dont know I need to rotate here
        #result = result[:result.index(num)] + result[result.index(num)+1:]
        
        # rotate list by num
        indn = (num  % len(result))  
        result = result[:indn] + [num] + result[indn:]
    
    return result

def Part1(filename, printing = False):
    """
    Solutions to Part 1
    """
    encrypted_file = ParseData(filename)
    
    result = encrypted_file[:]
    result = Solve(result,encrypted_file)

    # set 0 to begingin of list
    result = result[result.index(0):] + result[:result.index(0)]

    ans = []
    for nth in [1000,2000,3000]:
        resind =  (nth) % (len(result))
        ans.append(result[resind])
    return sum(ans)

def Part2(filename, printing = False):
    """
    Solutions to Part 2
    """
    encrypted_file = ParseData(filename)

    new  = []
    for num in encrypted_file:
        new.append(num*811589153)
    
    encrypted_file = new
    
    result = encrypted_file[:]
    for r in range(0,10):
        result = Solve(result,encrypted_file)
    
    # set 0 to begingin of list
    result = result[result.index(0):] + result[:result.index(0)]

    ans = []
    for nth in [1000,2000,3000]:
        resind =  (nth) % (len(result))
        ans.append(result[resind])
    return sum(ans)

def main():
    """
    main function
    """ 
    testfile  = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day18")

    print(f'Tests : Answer to Part 1 = {Part1(testfile, printing = False)}')    
    print(f'Tests : Answer to Part 2 = {Part2(testfile, printing = False)}\n')

    assert(Part1(testfile) == 3)
    assert(Part2(testfile) == 1623178306)

    print(f'Inputs: Answer to Part 1 = {Part1(inputfile, printing = False)}')
    print(f'Inputs: Answer to Part 2 = {Part2(inputfile, printing = False)}')

if __name__ == "__main__":
    main()