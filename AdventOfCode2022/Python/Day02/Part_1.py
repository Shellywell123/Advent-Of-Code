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

    totals = []
    outcomes = ["BX","CY","AZ","AX","BY","CZ","CX","AY","BZ"]

    for round in data:
        score = outcomes.index(round)+1
        print(round, score)
        totals.append(score)
        
    print(len(totals))
    print(sum(totals))

    return 0

if __name__ == "__main__":
    main()