#include <iostream>
#include <fstream>
#include <string>

using namespace std;

int main ()
{

  int running_count = 0;
  ifstream file("inputs.txt");

  int pprevoius_entry =0;
  int previous_entry =0;
  int previous_running_sum=0;

  cout << "reading in file\n\n";

  string str;
  while (getline(file, str)) 
  {
    int current_entry = stoi(str);
    
    int current_running_sum = current_entry + previous_entry + pprevoius_entry;
    cout << current_entry << " - (" << current_running_sum << "=" << current_entry << "+"<< previous_entry << "+" << pprevoius_entry << ")";
    
    if(previous_running_sum < current_running_sum) 
    {
        cout << " - increased\n";
        running_count+=1;
    }
    else
    {
        cout << "\n";
    }

    pprevoius_entry = previous_entry;
    previous_entry = current_entry;
    

    previous_running_sum = current_running_sum;
  }
  cout << "\nans = " << running_count-3 << "\ndone\n";
}