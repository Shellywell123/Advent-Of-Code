
def import_data(filename):
    """
    imports list of numbers delimited by new lines
    """

    with open(filename, 'r') as f:
        content = f.read().split('\n')

    return content   


def format_data(data):
    """
    """

    for i in range(0,len(data)):
        data[i] = list(data[i])
        #print(data[i])
    return data

def q1():
    """
    """
    
    def check_adjacent(character,seat_index,chair_plan):
        """
        """
        row,col=seat_index

        count = 0

        # try row above
        
        def test(chair_plan,row,col):
            """
            """
            try:
                if chair_plan[row][col] == character:
                    return True
                else:
                    return False
            except:
                pass
            return False

        for i in [col-1,col,col+1]:
            if test(chair_plan,row-1,i) == True:
                count += 1

        # cuurent row
        for i in [col-1,col+1]:
            if test(chair_plan,row,i) == True:
                count += 1

        # try row below     
        for i in [col-1,col,col+1]:
            try:
                if test(chair_plan,row+1,i) == True:
                    count += 1
                except IndexError:
               #     count -= 1
                    continue  

        return count

    def count_chairs(character,chair_plan):
        """
        """

        count = 0

        for i in range(0,len(chair_plan)):
            for j in range(0,len(chair_plan[i])):
                symbol = chair_plan[i][j]
                if symbol == character:
                    count += 1

        return count

    def recurr_(current_chair_plan):
        """
        """
            
        next_chair_plan = []
        changes = 0
       
        # iterate through rows
        for i in range(0,len(current_chair_plan)):
            # iterate through cols
            next_chair_plan.append([])
            checks = []
            

            for j in range(0,len(current_chair_plan[i])):
                symbol = current_chair_plan[i][j]
                check  = check_adjacent('#',[i,j],current_chair_plan)
                checks.append(check)
          #      print(check,)
                    
                if symbol == '#' and check >= 4:
                    # if seat occupied check surrounding seats
                    changes += 1
                    next_chair_plan[i].append('L')

                elif symbol == 'L' and check == 0:
                    # if seat empty check surrounding seats
                    changes += 1
                    next_chair_plan[i].append('#')

                else:
                    next_chair_plan[i].append(symbol)

            print(''.join(current_chair_plan[i])+'    '+''.join(next_chair_plan[i]),'   ',checks)


        if changes == 0:
            print(count_chairs('#',next_chair_plan))
            return next_chair_plan
        else:
            print(changes,'\n')
            recurr_(next_chair_plan)
    
    data = format_data(import_data('inputdata.txt'))
    data = format_data(import_data('testdata.txt'))
    data = format_data(['####','####','####','####'])
    chair_plan = recurr_(data)

def main():
    """
    """
    import sys
    sys.setrecursionlimit(10**6)
    q1()
        
if __name__ == "__main__":
    main()