def import_data(filename):
    """
    imports list of numbers delimited by new lines
    """

    array = []

    with open(filename, 'r') as f:
        for line in f.readlines():
            array.append(line[:-1])
    return array

def tree_count(data,right_jump,down_jump):
    """
    """

    depth,width =len(data),len(data[0])
 #   print(depth,width)

    trees = 0

    x_pos = 1
    for y_pos in range(0,depth,down_jump):
        
        # when x_pos becomes larger that the width of the slope loop back
        if x_pos > width:
            diff = x_pos - width
            x_pos = diff

        grid = data[y_pos][x_pos-1]
   #     print(y_pos,x_pos,grid)

        if grid == '#':
            trees = trees + 1 

        x_pos = x_pos + right_jump
    
    print('- ',trees)
    return trees


def main():
    """
    """
    data = import_data('datainput.txt')
   # print(data)
    

    a = tree_count(data,1,1)
    b = tree_count(data,3,1)
    c = tree_count(data,5,1)
    d = tree_count(data,7,1)
    e = tree_count(data,1,2)

    prod = a*b*c*d*e

    print(prod)

if __name__ == "__main__":
    main()
