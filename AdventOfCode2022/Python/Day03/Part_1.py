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
    data = get_inputs("inputs.txt")

    priority_values = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

    priority_scores = []

    for rucksack in data:
        split = int(len(rucksack)/2)
        first_compartment  = [*set(list(rucksack[:split]))]
        second_compartment = [*set(list(rucksack[split:]))]

        for x in list(first_compartment):
            for y in list(second_compartment):
                if x == y and x in list(priority_values):
                    priority_scores.append(priority_values.index(x)+1)
        
    print(sum(priority_scores))
        
    return 0

if __name__ == "__main__":
    main()