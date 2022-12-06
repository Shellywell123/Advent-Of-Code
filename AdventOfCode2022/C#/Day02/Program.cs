using System.Collections.Generic;

void Part1(string filename)
{
    string[] lines = System.IO.File.ReadAllLines(filename);

    List<String> outcomes = new List<String>{ "B X", "C Y", "A Z", "A X", "B Y", "C Z", "C X", "A Y", "B Z"};
    List<int> scores = new List<int> { };

    foreach (string line in lines)
    {
        scores.Add(outcomes.IndexOf(line) + 1);
    }

    Console.WriteLine("Answer to Part 1 = " + scores.Sum());
}


void Part2(string filename)
{
    string[] lines = System.IO.File.ReadAllLines(filename);

    List<String> outcomes = new List<String> { "B X", "C X", "A X", "A Y", "B Y", "C Y", "C Z", "A Z", "B Z" };
    List<int> scores = new List<int> { };

    foreach (string line in lines)
    {
        scores.Add(outcomes.IndexOf(line) + 1);
    }

    Console.WriteLine("Answer to Part 2 = " + scores.Sum());
}

Console.WriteLine("Advent-Of-Code 2022 - Day02");
Part1(@"..\..\..\inputs.txt");
Part2(@"..\..\..\inputs.txt");