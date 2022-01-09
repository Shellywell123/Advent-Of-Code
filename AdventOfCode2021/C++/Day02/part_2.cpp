#include <iostream>
#include <fstream>
#include <string>

using namespace std;

bool is_string_int(const string s){
  return s.find_first_not_of( "0123456789" ) == string::npos;
}

int main ()
{

  int hor_pos = 0;
  int dep_pos = 0;
  int aim = 0;
  string direction;
  int amount;
  
  ifstream file("inputs.txt");

  cout << "reading in file\n\n";

  while (getline(file, direction)) 
  {
    while (file >> direction >> amount)
    {
       cout << direction << " " <<amount << " " << dep_pos << "\n";
       if (direction == "forward")
       {
         hor_pos += amount;
         dep_pos += amount*aim;
       }
       if (direction == "up")
       {
         aim -= amount;
       }
       if (direction == "down")
       {
         aim += amount;
       }
    }
    //int current_entry = stoi(str);
    //cout << str << " a " << str1 <<"\n";2
  }
  cout << "\nans = " << dep_pos*hor_pos << "\n depth = " << dep_pos << "\n horiz = " << hor_pos << "\n aim = " << aim <<"\ndone\n";
}