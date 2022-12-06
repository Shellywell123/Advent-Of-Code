void Part1(string filename)
{
    string[] lines = System.IO.File.ReadAllLines(filename);

    List<int> elfs = new List<int>();

    int elf = 0;
    foreach (string line in lines)
    {
        try
        {
            elf += int.Parse(line);
        }
        catch
        {
            elfs.Add(elf);
            elf = 0;
        }
    }

    Console.WriteLine("Answer to Part 1 = "+elfs.Max());
}

void Part2(string filename)
{
    string[] lines = System.IO.File.ReadAllLines(filename);

    List<int> elfs = new List<int>();

    int elf = 0;
    foreach (string line in lines)
    {
        try
        {
            elf += int.Parse(line);
        }
        catch
        {
            elfs.Add(elf);
            elf = 0;
        }
    }

    elfs.Sort();
    elfs.Reverse();

    Console.WriteLine("Answer to Part 2 = " + (elfs[0] + elfs[1] + elfs[2]));
}

Console.WriteLine("Advent-Of-Code 2022 - Day01");
Part1(@"..\..\..\inputs.txt");
Part2(@"..\..\..\inputs.txt");