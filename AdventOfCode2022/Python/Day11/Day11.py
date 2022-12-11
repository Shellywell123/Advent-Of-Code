def get_inputs(file):
    """
    function to read in data from a txt file
    """

    with open(file) as f:
        data = f.read().split('\n')

    return data

def MonkeyParser(data):
    """
    parse the monkey data into a useable 
    list of dictionaries
    """

    monkeys = []
    current_monkey = {}
    
    for i in range(0,len(data)):

        line = data[i]

        if line.strip() == '':
            continue

        if (i == (len(data)-1)):
            monkeys.append(current_monkey)
        
        if line.split(' ')[0] == "Monkey":
            monkeys.append(current_monkey)
            current_monkey = {}
            current_monkey['number'] = int(line.split(' ')[1][:-1])
        
        else:
            if line.strip().split(' ')[0] == "Starting":
                starting_items = line.split(':')[1].split(',')
                ints = []
                for item in starting_items:
                    ints.append(int(item))
                current_monkey['starting_items'] = ints
            
            if line.strip().split(':')[0] == "Operation":
                current_monkey['operation'] = line.split(':')[1].strip().split('=')[1]
            
            if line.strip().split(' ')[0] == "Test:":
                current_monkey['test_divisible_by'] = int(line.split(':')[1].strip().split(' ')[-1])

            if line.strip().split(':')[0] == "If true":
                current_monkey['if_true_throw_to_monkey'] = int(line.split(':')[1].strip().split(' ')[-1])
            
            if line.strip().split(':')[0] == "If false":
                current_monkey['if_false_throw_to_monkey'] = int(line.split(':')[1].strip().split(' ')[-1])
    
    return monkeys[1:]

def Part1(filename):
    """
    Solutions to Part 1
    """
    data = get_inputs(filename)

    monkeys = MonkeyParser(data)

    inspections = [0]*len(monkeys)
    for round in range(1,20 + 1):

        for m in range(0,len(monkeys)):

            starting_items           = monkeys[m]['starting_items']   
            operation                = monkeys[m]['operation']           
            test_divisible_by        = monkeys[m]['test_divisible_by']
            if_true_throw_to_monkey  = monkeys[m]['if_true_throw_to_monkey']
            if_false_throw_to_monkey = monkeys[m]['if_false_throw_to_monkey']

            for worry_level in starting_items:
                inspections[m] += 1

                old = worry_level
                worry_level = eval(operation)
                worry_level = int(worry_level / 3)

                if worry_level % test_divisible_by == 0:
                    monkeys[if_true_throw_to_monkey]['starting_items'].append(worry_level)
                else:
                    monkeys[if_false_throw_to_monkey]['starting_items'].append(worry_level)
            
            monkeys[m]['starting_items'] = []
        
    inspections.sort()
    inspections.reverse()
    return inspections[0] * inspections[1]

def Part2(filename):
    """
    Solutions to Part 2
    """
    data = get_inputs(filename)
    
    monkeys = MonkeyParser(data)
    LCM = 1
    for monkey in monkeys:
        LCM *= monkey['test_divisible_by']

    inspections = [0]*len(monkeys)
    for round in range(1,10000 + 1):
        
        for m in range(0,len(monkeys)):

            for old in monkeys[m]['starting_items']:
                inspections[m] += 1

                worry_level = eval(monkeys[m]['operation']) 

                # monkey anxiety medicine (thanks fergus)
                worry_level = worry_level % LCM

                if worry_level % monkeys[m]['test_divisible_by'] == 0:
                    monkeys[monkeys[m]['if_true_throw_to_monkey']]['starting_items'].append(worry_level)
                else:
                    monkeys[monkeys[m]['if_false_throw_to_monkey']]['starting_items'].append(worry_level)
            
            monkeys[m]['starting_items'] = []

    inspections.sort()
    inspections.reverse()
    return inspections[0] * inspections[1]

def main():
    """
    main function
    """ 
    testfile  = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day11")

    print(f'"{testfile}": Answer to Part 1 = {Part1(testfile)}')
    print(f'"{testfile}": Answer to Part 2 = {Part2(testfile)}\n')

    print(f'"{inputfile}": Answer to Part 1 = {Part1(inputfile)}')
    print(f'"{inputfile}": Answer to Part 2 = {Part2(inputfile)}')

if __name__ == "__main__":
    main()