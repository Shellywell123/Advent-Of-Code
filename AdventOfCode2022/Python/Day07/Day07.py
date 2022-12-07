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

    # format commands
    history = []
    current_command = []
    for line in data:
        # commands
        if line[0] == "$":
            # if current_command is empty
            if not current_command:
                current_command = [line, '']
            else:
                history.append(current_command)
                current_command = [line, '']
        # outputs 
        else:
            current_command[1] += line + ','
    history.append(current_command)

    # execute commands
    tree = {}
    cwd = ''

    for entry in history:
        command = entry [0]
        output = entry[1]

        # cd commands
        if command[:4] == "$ cd":

            cd_location = command[5:]

            # root
            if cd_location == "/":
                cwd = "/"
                tree["/"] = {}

            # move up a dir
            elif cd_location == "..":
                
                new_cwd = ''
                for dir in cwd.split('/')[:-2]:
                    new_cwd += dir +"/"
                cwd = new_cwd           

            # change to supplied location
            else:
                cwd += cd_location+"/"
            
        # list command
        elif command[:4] == "$ ls":

            for result in output.split(','):
                if result:
                    # dirs
                    if result[:3] == 'dir':
                        dir = result.split(' ')[1]
                        tree_str = '["/"]'
                        for dirp in (cwd.split('/')[1:-1]):
                            tree_str += f'["{dirp}"]'
                        eval(f'tree{tree_str}.update('+'{"'+dir+'":{}})')

                    # files
                    else:
                        filesize,file = result.split(' ')[0], result.split(' ')[1]
                        tree_str = '["/"]'
                        for dirp in (cwd.split('/')[1:-1]):
                            tree_str += f'["{dirp}"]'
                        eval(f'tree{tree_str}.update('+'{"'+file+'":'+filesize+'})')        
        
    print("cwd: ",cwd)
    print(json.dumps(tree, indent=4))
    print('#'*10)

    # count bytes
    results = []
    def function(dictt):
        keys = list(dictt.keys())
        vals = list(dictt.values())
        for key in keys:

            if type(dictt[key]) is dict:
                results.append({key: dictt[key]})
                function(dictt[key])

    function(tree)

    dirsizes = []
    for result in results:
        str_ = str(result)
        junk = list(str_.replace('{','').replace('}','').replace("'",'').replace(",",':').split(':'))
        total = 0
        for j in junk:
            try:
                total += int(j)
            except:
                continue
        
        print(list(result.keys())[0],total)
        dirsizes.append([list(result.keys())[0],total])
    
    # retrieve ans
    ans = 0 
    for dirsize in dirsizes:
        dir = dirsize[0]
        size = dirsize[1]

        if size <= 100000:
            ans += size
    
    print(f"Answer to Part 1 = {ans}")

    return dirsizes

def Part2(filename):
    """
    Solutions to Part 2
    """
    dirsizes = Part1(filename)

    used_size = dirsizes[0][1]
    total_size = 70000000
    required_size = 30000000

    options = []
    for dirsize in dirsizes:
        dir = dirsize[0]
        size = dirsize[1]

        if required_size <= (total_size - used_size + size):
            options.append(size)
    
    ans = min(options)
    print(f"Answer to Part 2 = {ans}")

def main():
    """
    main function
    """ 
    print("Advent-Of-Code 2022 - Day07")
    Part1("inputs.txt")
    Part2("inputs.txt")

if __name__ == "__main__":
    main()