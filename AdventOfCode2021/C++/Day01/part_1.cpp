#include <iostream>
#include <fstream>
#include <string>

using namespace std;

int main ()
{

  int running_count = 0;
  ifstream file("inputs.txt");
  int previous_entry =0;

  cout << "reading in file\n\n";

  string str;
  while (getline(file, str)) 
  {
    int current_entry = stoi(str);
    cout << current_entry;
    if(previous_entry < current_entry) 
    {
        cout << " - increased from " << previous_entry << "\n";
        running_count+=1;
    }
    else
    {
        cout << "\n";
    }
    previous_entry = current_entry;
  }
  cout << "\nans = " << running_count -1 << "\ndone\n";
}