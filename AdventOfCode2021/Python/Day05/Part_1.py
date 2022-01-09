def get_inputs():
    """
    function to read in list of ints from a txt file
    """

    # read in file
    with open("inputs.txt") as f:
        contents = f.readlines()

    # format inputs
    inputs = []

    for line in contents:
        points = line.replace('\n','').split(' -> ')
        coords1 = points[0].split(',')
        coords2 = points[1].split(',')

        inputs.append([coords1,coords2])

    return inputs


def main():
    """
    main function
    """ 
    inputs = get_inputs()
    
    # calculate canvas limits
    canvas_width = 0
    canvas_height = 0
    for line in inputs:
        for coord in line:
            if int(coord[0]) > canvas_width:
                canvas_width = int(coord[0])
            if int(coord[1]) > canvas_height:
                canvas_height = int(coord[1])

    # create canvas
    canvas = []
    for y in range(0,canvas_height+1):
        row = []
        for x in range(0,canvas_width+1):
            row.append(0)
        canvas.append(row)

    # draw lines
    for line in inputs:
        x1 = int(line[0][0])
        y1 = int(line[0][1])
        x2 = int(line[1][0])
        y2 = int(line[1][1])

        # skip diagonal lines
        if (x1!=x2) and (y1!=y2):
            continue

        # y = mx +c
        try:
            m = (y2-y1)/(x2-x1)
        except ZeroDivisionError:
            m = 0

        c = y1-m*x1

        if x1==x2:
            m = 'undefined'
            c = 'undefined'
        
        if m == 'undefined':
            if y2>y1:
                for y in range(y1,y2+1):
                    x = x1
                    canvas[y][x] +=1
            else:
                for y in range(y2,y1+1):
                    x = x1
                    canvas[y][x] +=1
        else:
            if x2>x1:
                for x in range(x1,x2+1):
                    y = int(m*x +c)
                    canvas[y][x] +=1
            else:
                for x in range(x2,x1+1):
                    y = int(m*x +c)
                    canvas[y][x] +=1

    # print canvas
    for row in canvas:
        print(row)

    # count intersections
    count = 0
    for row in canvas:
        for num in row:
            if num >= 2:
                count +=1

    ans = count
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()