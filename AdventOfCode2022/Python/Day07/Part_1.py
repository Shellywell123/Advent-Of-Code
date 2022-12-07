def get_inputs(file):
    """
    function to read in list of ints from a txt file
    """

    with open(file) as f:
        data =((f.read()).replace(" ","")).split("\n")

    return data[0] 

def main():
    """
    main function
    """ 
    data = get_inputs("inputs.txt")

    print(data)
    for i in range(0,len(data)-4):
        chars =data[i:i+4]

        print(i, chars)
        dup_checks = []
        for char in chars:
            
            if char not in dup_checks:
                if len(dup_checks) == 3:
                    print('ans = ',i+4)
                    return 0
                dup_checks.append(char)

            if char in dup_checks:
                continue
            print(dup_checks)
        
    return 0

if __name__ == "__main__":
    main()