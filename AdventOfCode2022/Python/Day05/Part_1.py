import numpy as np

def get_inputs(file):
    """
    function to read in list of ints from a txt file
    """

    with open(file) as f:
        data =(f.read()).split("\n")

    stacks = data[:data.index("")]
    moves = data[(data.index("")+1):]

    return stacks, moves

def format_stacks(stacks):
    """
    """
    formatted = []
    for stack in stacks:
        formatted.append(list((stack.replace('    ','[#] ').replace('][','] [').replace(']  [','] [').replace(' ','').replace('[','').replace(']',''))))

    formatted = np.fliplr(np.transpose((formatted)))
    formatted2 = []

    for stack in formatted:
        while '#' in list(stack):
            stack = list(stack)
            stack.remove('#')

        formatted2.append(list(stack))

    return formatted2

def main():
    """
    main function
    """ 
    stacks, moves = get_inputs("inputs.txt")
    stacks = format_stacks(stacks)
    
    for move in moves:
        move_info = move.split(' ')

        amount, from_stack, to_stack = int(move_info[1]), int(move_info[3]), int(move_info[5])

        # collect stack amount
        in_transit = stacks[from_stack-1][-amount:]
        in_transit.reverse()
        del stacks[from_stack-1][-amount:] 
        for box in in_transit:
            stacks[to_stack-1].append(box)
        
    ans = ''
    for stack in stacks:
        ans += str(stack[-1])

    print(ans)

    return 0

if __name__ == "__main__":
    main()