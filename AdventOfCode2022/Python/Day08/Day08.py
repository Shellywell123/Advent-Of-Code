import json

def get_inputs(file):
    """
    function to read in list of ints from a txt file
    """

    with open(file) as f:
        data = f.read().split('\n')

    return data

def Part1(filename):
    """
    Solutions to Part 1
    """
    data = get_inputs(filename)

    # count trees
    interior_trees = 0
    for i in range(0,len(data)):
        row_of_trees = data[i]
        for j in range(0, len(row_of_trees)):
            tree_being_checked = int(data[i][j])

            # visibility from top
            top_check = True
            for y in range(0,i):
                if int(data[y][j]) >= tree_being_checked:
                    top_check = False
                    
            # visibility from bottom
            bot_check = True
            for y in range(i+1, len(data)):
                if int(data[y][j]) >= tree_being_checked:
                    bot_check = False

            # visibility from left
            lef_check = True
            for x in range(0, j):
                if int(data[i][x]) >= tree_being_checked:
                    lef_check = False

            # visibility from right
            rig_check = True
            for x in range(j+1, len(data[0])):
                if int(data[i][x]) >= tree_being_checked:
                    rig_check = False
            
            if any ([top_check, bot_check, lef_check, rig_check]):
                interior_trees += 1

    ans = interior_trees #- edge_trees
    return ans
    

def Part2(filename):
    """
    Solutions to Part 2
    """
    data = get_inputs(filename)

    # count trees
    scenic_scores = []
    for i in range(1,len(data)-1):
        row_of_trees = data[i]
        for j in range(1, len(row_of_trees)-1):
            tree_being_checked = int(data[i][j])

            # visibility from top
            top_check = True
            top_view = 0
            top_list = []
            # not sure why python wouldnt let me use a descending range
            for y in range(0,i):
                top_list.append(data[y][j])
            top_list.reverse()
            for d in top_list:
                top_view += 1
                if int(d) >= tree_being_checked:                    
                    top_check = False
                    break
                    
            # visibility from bottom
            bot_check = True
            bot_view = 0
            for y in range(i+1, len(data)):
                bot_view += 1
                if int(data[y][j]) >= tree_being_checked:                    
                    bot_check = False
                    break

            # visibility from left
            lef_check = True
            lef_view = 0
            lef_list = []
            for x in range(0,j):
                lef_list.append(data[i][x])
            lef_list.reverse()
            for d in lef_list:
                lef_view += 1
                if int(d) >= tree_being_checked:
                    lef_check = False
                    break

            # visibility from right
            rig_check = True
            rig_view = 0
            for x in range(j+1, len(data[0])):
                rig_view += 1
                
                if int(data[i][x]) >= tree_being_checked:
                    rig_check = False
                    break
            
            if any ([top_check, bot_check, lef_check, rig_check]):
                scenic_scores.append(top_view * bot_view * lef_view * rig_view)

    ans = max(scenic_scores)
    return ans

def main():
    """
    main function
    """ 
    testfile = "tests.txt"
    inputfile = "inputs.txt"

    print("Advent-Of-Code 2022 - Day08")

    print(f'\n"{testfile}"')
    print(f' Answer to Part 1 = {Part1(testfile)}')
    print(f' Answer to Part 2 = {Part2(testfile)}')

    print(f'\n"{inputfile}"')
    print(f' Answer to Part 1 = {Part1(inputfile)}')
    print(f' Answer to Part 2 = {Part2(inputfile)}')

if __name__ == "__main__":
    main()