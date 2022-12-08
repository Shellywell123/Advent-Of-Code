string testfile = @"..\..\..\tests.txt";
string inputfile = @"..\..\..\inputs.txt";
string priority_values = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

string RemoveDuplicates(string input_string)
{
    string output_string= "";
    foreach (char char_ in input_string)
    {
        if (output_string.IndexOf(char_) == -1)
        {
            output_string += char_;
        }
    }
    return output_string;
}

int Part1(string filename)
{
    string[] rucksacks = System.IO.File.ReadAllLines(filename);
    
    int priority_scores = 0;

    foreach (string rucksack in rucksacks)
    {
        int midpoint = (rucksack.Length / 2);
        
        string left_compartment  = RemoveDuplicates(rucksack.Substring(0, midpoint));
        string right_compartment = RemoveDuplicates(rucksack.Substring(midpoint, midpoint));

        foreach (char left_item in left_compartment)
        {
            foreach (char right_item in right_compartment)
            {
                if (left_item == right_item)
                {
                    priority_scores += (priority_values.IndexOf(left_item)+1);
                }
            }
        }
    }

    return priority_scores;
}

int Part2(string filename)
{
    string[] rucksacks = System.IO.File.ReadAllLines(filename);

    int badges = 0;

    for (int i = 0; i < rucksacks.Length; i+=3)
    {
        string first_rucksack  = RemoveDuplicates(rucksacks[i]);
        string second_rucksack = RemoveDuplicates(rucksacks[i+1]);
        string third_rucksack  = RemoveDuplicates(rucksacks[i+2]);

        foreach (char first_item in first_rucksack)
        {
            foreach (char second_item in second_rucksack)
            {
                foreach (char third_item in third_rucksack)
                {
                    if (first_item == second_item && second_item == third_item)
                    {
                        badges += (priority_values.IndexOf(first_item) + 1);
                    }
                }
            }
        }
    }
    return badges;
}

Console.WriteLine("Advent-Of-Code 2022 - Day03");

Console.WriteLine("Tests:  Answer to Part 1 = " + (Part1(testfile)).ToString());
Console.WriteLine("Tests:  Answer to Part 2 = " + (Part2(testfile)).ToString());

Console.WriteLine("Inputs: Answer to Part 1 = " + (Part1(inputfile)).ToString());
Console.WriteLine("Inputs: Answer to Part 2 = " + (Part2(inputfile)).ToString());