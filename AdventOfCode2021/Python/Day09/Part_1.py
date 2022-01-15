def get_inputs():
    """
    function to read in list of ints from a txt file
    """

    # read in file
    with open("inputs.txt") as f:
        contents = f.read()

    # format data
    data = []
    for content in contents.split('\n'):
        data.append(content)

    return data


def main():
    """
    main function
    """ 
    data = get_inputs()

    # format edges of canvas
    data = ['9'*len(data[0])] + data + ['9'*len(data[0])]

    for i in range(0,len(data)):
        data[i] = '9' + data[i] + '9'

    # find low points in selected window starting at data[1][1]
    risk = []
    for row_ind in range(1,len(data)-1):
        for col_ind in  range(1,len(data[0])-1):
            selected_num = int(data[row_ind][col_ind])

            l = int(data[row_ind][col_ind-1])
            r = int(data[row_ind][col_ind+1])
            u = int(data[row_ind-1][col_ind])
            d = int(data[row_ind+1][col_ind])

            window = [selected_num,l,r,u,d]
            if min(window)==int(selected_num) and window.count(selected_num)==1:
                risk.append(int(selected_num)+1)

    ans = sum(risk)
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()