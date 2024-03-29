import networkx as nx

def ParseData(file):    
    """
    function to read and parse in data from a txt file
    """
    with open(file) as f:
        data = f.read().split('\n')

    valves = {}
    for i in range(0,len(data)):
        line_string = data[i].split(':')[0]
        name    = line_string.split(' ')[1]
        rate    = int((line_string.split('=')[1]).split(';')[0])
        tunnels = (line_string[line_string.index('to valve')+9:]).strip().split(', ')
        valves[name] = {"Rate": rate, "Tunnels": tunnels}
    
    return valves

def Part1(filename, printing = False):
    """
    Solutions to Part 1
    First time encountering dymnamic programming
    TOP DOWN recursion with caching
    """

    valves = ParseData(filename)

    Temp = nx.Graph()
    edges = []
    for valve1 in valves:
        Temp.add_node(valve1)
        for valve2 in dict(valves[valve1])["Tunnels"]:
            edges.append((valve1,valve2, {"weight": 1}))
    Temp.add_edges_from(edges)

    distances = dict(nx.all_pairs_shortest_path_length(Temp))

    non_zero_valves = []
    for valve in valves:
        if dict(valves[valve])["Rate"] != 0:
            non_zero_valves.append(valve)

    open_valves = []

    def decide(current_valve, open_valves, minutes):
        """
        """
        if minutes < 0:
            return 0
        
        current_best_decision_result = 0
        # try moving to a valve and opening it
        for next_valve in non_zero_valves:

            if current_valve == next_valve: 
                continue

            dist = distances[current_valve][next_valve]
            rate = dict(valves[next_valve])["Rate"]

            if next_valve not in open_valves and dist < minutes:
                open_valves.append(next_valve)
                minutes_if_open_valve = minutes - (dist + 1)
                relief = rate*minutes_if_open_valve
                open_decision_result, open_valves = decide(next_valve,open_valves, minutes_if_open_valve)
                open_decision_result = open_decision_result + relief

                # check if result was highest result yet
                if open_decision_result >= current_best_decision_result:
                    current_best_decision_result = open_decision_result
                # if it wasnt worth opening close it and try the next valve
                open_valves.remove(next_valve)

        return current_best_decision_result, open_valves

    return decide("AA",[],30)[0]

def Part2(filename, printing = False):
    """
    Solutions to Part 2
    """
   
    valves = ParseData(filename)

    Temp = nx.Graph()
    edges = []
    for valve1 in valves:
        Temp.add_node(valve1)
        for valve2 in dict(valves[valve1])["Tunnels"]:
            edges.append((valve1,valve2, {"weight": 1}))
    Temp.add_edges_from(edges)

    distances = dict(nx.all_pairs_shortest_path_length(Temp))

    non_zero_valves = []
    for valve in valves:
        if dict(valves[valve])["Rate"] != 0:
            non_zero_valves.append(valve)

    def decide(current_valve, ele_current_valve, open_valves, minutes, ele_minutes):
        """
        """
        if minutes < 0:
            return 0
        
        current_best_decision_result = 0
        # try moving to a valve and opening it
        for next_valve in non_zero_valves:

            if current_valve == next_valve: 
                continue

            dist = distances[current_valve][next_valve]
            rate = dict(valves[next_valve])["Rate"]

            if next_valve not in open_valves and dist < minutes:
                open_valves.append(next_valve)
                minutes_if_open_valve = minutes - (dist + 1)
                relief = rate*minutes_if_open_valve
                open_decision_result, open_valves = ele_decide(next_valve, ele_current_valve, open_valves, minutes_if_open_valve, ele_minutes)
                open_decision_result = open_decision_result + relief

                # check if result was highest result yet
                if open_decision_result >= current_best_decision_result:
                    current_best_decision_result = open_decision_result
                # if it wasnt worth opening close it and try the next valve
                open_valves.remove(next_valve)

        return current_best_decision_result, open_valves
    
    def ele_decide(current_valve, ele_current_valve, open_valves, minutes, ele_minutes):
        """
        """
        if minutes < 0:
            return 0
        
        current_best_decision_result = 0
        # try moving to a valve and opening it
        for ele_next_valve in non_zero_valves:

            if ele_current_valve == ele_next_valve: 
                continue

            dist = distances[ele_current_valve][ele_next_valve]
            rate = dict(valves[ele_next_valve])["Rate"]

            if ele_next_valve not in open_valves and dist < ele_minutes:
                open_valves.append(ele_next_valve)
                ele_minutes_if_open_valve = ele_minutes - (dist + 1)
                relief = rate*ele_minutes_if_open_valve
                open_decision_result, open_valves = decide(current_valve, ele_next_valve, open_valves, minutes, ele_minutes_if_open_valve)
                open_decision_result = open_decision_result + relief

                # check if result was highest result yet
                if open_decision_result >= current_best_decision_result:
                    current_best_decision_result = open_decision_result
                # if it wasnt worth opening close it and try the next valve
                open_valves.remove(ele_next_valve)

        return current_best_decision_result, open_valves

    return decide("AA","AA",[],26, 26)[0]

def main():
    """
    main function
    """ 
    testfile  = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day16")

    print(f'Tests : Answer to Part 1 = {Part1(testfile, printing = False)}')    
    print(f'Tests : Answer to Part 2 = {Part2(testfile, printing = False)}\n')

    assert(Part1(testfile) == 1651)
    assert(Part2(testfile) == 1707)

    print(f'Inputs: Answer to Part 1 = {Part1(inputfile, printing = False)}')
    print(f'Inputs: Answer to Part 2 = {Part2(inputfile, printing = False)}')

if __name__ == "__main__":
    main()