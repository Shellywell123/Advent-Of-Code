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

    badges = []

    for i in range(0,int(len(data)),3):

        rucksack_1 = [*set(list(data[i]))]
        rucksack_2 = [*set(list(data[i+1]))]
        rucksack_3 = [*set(list(data[i+2]))]

        for x in list(rucksack_1):
            for y in list(rucksack_2):
                for z in list(rucksack_3):
                    if x == y == z and x in list(priority_values):
                        badges.append(priority_values.index(x)+1)
        
    print(sum(badges))
        
    return 0

if __name__ == "__main__":
    main()