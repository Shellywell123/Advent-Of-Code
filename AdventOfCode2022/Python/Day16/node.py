# First networkx library is imported 
# along with matplotlib
import networkx as nx
import matplotlib.pyplot as plt
from itertools import combinations
from random import choice
  
visual = []

        
# addEdge function inputs the vertices of an
# edge and appends it to the visual list
def addEdge( a, b):
    temp = [a, b]
    visual.append(temp)
    
  
def ParseData(file):    
    """
    function to read and parse in data from a txt file
    (rather messy)
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
    
    for v in valves:
        print(v)
    return valves



# Driver code
valves = ParseData("inputs.txt")

    
G = nx.Graph()
for valve in valves:
    print(valve,valves[valve])
    
    for v in dict(valves[valve])["Tunnels"]:
        G.add_edge(valve,v,time=1)

length = dict(nx.all_pairs_shortest_path_length(G))

pos = nx.shell_layout(G)

edge_labels = dict([((n1, n2), 1)
                    for n1, n2, d in G.edges(data=True)])

node_labels = dict([((n), str(n)+'\n' +str(dict(valves[n])['Rate']))
                    for n, d in G.nodes(data=True)])

nx.draw_networkx_edge_labels(G, pos, edge_labels=edge_labels, label_pos=0.5,
                             font_color='red', font_size=10)

nx.draw_networkx(G,pos, edge_color="tab:red",with_labels = True,labels = node_labels, font_size=7, font_color='black',node_size=500)


lengths_no_zero = []

for valve in length:
    
    if dict(valves[valve])["Rate"] != 0:
        lengths_no_zero.append(valve)

print('combs')
#plt.show()
from itertools import permutations
combs = list(combinations(lengths_no_zero,8))

print('perms')
perms = []
for comb in combs:
    perms += list(permutations(comb))

#perms = [['DD'],['DD', 'BB', 'JJ', 'HH', 'EE', 'CC'],['BB']]

print('combs = ',len(combs))
print('perms = ',len(perms))

paths = []
ans = []

for k in range(0,len(perms)-1):


    p = perms[k]

    if p[0] != 'NU':
        continue


    minute = 0
    pressure = 0
    current_valve = 'AA'
    open_valves = []
    i = 0
    while minute <= 30 and i < len(p):
        
      #  while minute <= 30:
        open_valves.append(current_valve)
        # random
        next_valve = p[i]
        
        # pick highest
        # next_valve = 'AA'
        # for v in lengths_no_zero:
        #     if dict(valves[v])["Rate"] > dict(valves[next_valve])["Rate"]:
        #         next_valve = v
        dist_to_next_valve = length[current_valve][next_valve]
        rate_of_next_valve = dict(valves[next_valve])["Rate"]
        #print(minute,current_valve,next_valve,dist_to_next_valve,pressure)

        # move
        minute += dist_to_next_valve + 1
        # open
        pressure += ((30-minute ) * rate_of_next_valve)
        
        current_valve = next_valve
        i+=1
        #
    ans.append(pressure)
    paths.append(open_valves)
    print(f' progress = {round(k*100/len(perms)),3}% higest ans yet = {max(ans)}, cur_ans = {pressure}, cur_min = {minute}, cur_path = {open_valves}')

    # if open_valves == ['DD', 'BB', 'JJ', 'HH', 'EE', 'CC']:
    #     exit()# print(len(ans))
# for i in range(0,len(paths)-1):
#     print(i)
#     print(ans[i], open_valves[i])

print(max(ans))