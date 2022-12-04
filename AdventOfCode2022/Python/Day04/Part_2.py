def get_inputs(file):
    """
    function to read in list of ints from a txt file
    """

    with open(file) as f:
        data =((f.read()).replace(" ","")).split("\n")

    return data 

def main():
    """
    main function
    """ 
    data = get_inputs("tests.txt")

    conflict_pairs = 0
    
    for pair in data:
        first_elf_assignment, second_elf_assignment = pair.split(",")

        feal = int(first_elf_assignment.split('-')[0])
        feah = int(first_elf_assignment.split('-')[1])

        seal = int(second_elf_assignment.split('-')[0])
        seah = int(second_elf_assignment.split('-')[1])

        # if assignment 1 contains assignment 2
        if ((feal <= seal <= feah) or (feal <= seah <= feah)):
            conflict_pairs += 1
        
        # if assignment 2 contains assignment 1
        elif ((seal <= feal <= seah or seal <= feah <= seah)):
            conflict_pairs += 1
        
    print(conflict_pairs)
        
    return 0

if __name__ == "__main__":
    main()