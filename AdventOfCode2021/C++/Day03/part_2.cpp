#include <iostream>
#include <fstream>
#include <string>
#include <math.h> 
//#include <filesystem>
//namespace fs = std::filesystem;

using namespace std;


int main ()
{

  string oxygen_rate;
  string CO2_rate;

  ifstream file("tests.txt");
  cout << "reading in file ...\n";

  // get num of cols from counting line
  string sLine;
  int num_of_cols;
  int line_no = 1;
  while (line_no == 1 && getline(file, sLine)) 
  {
    num_of_cols = sLine.length();
    ++line_no;
  }

  // get num of rows
  ifstream file2("tests.txt");
  string sLine2;
  int num_of_rows;
  while (getline(file2, sLine2)) 
  {
    ++num_of_rows;
  }

  cout << "file has " << num_of_cols << " columns and " << num_of_rows << "\n";
  for (int col_ind = 0; col_ind < num_of_cols; col_ind++) 
  {
    ifstream file3("tests.txt");
    int oxygen_rate_add = 0;
    double line_no = 1;
    string sLine3;
    
    while (getline(file3, sLine3)) 
      {
        cout >> num_of_cols-1;
        if (line_no != num_of_cols-1)
        {
          string values = sLine3;
          string value = values.substr(col_ind,1);
          oxygen_rate_add += stoi(value);
          ++line_no;
        }
        else
        {
          oxygen_rate_add = 1;
        }
      }
    int num_of_rows = line_no;
    if (oxygen_rate_add==0.5)
      {
        oxygen_rate = "1";
        CO2_rate += "0";
      }
    else
      {
        cout << oxygen_rate_add << " " << line_no << " " << num_of_rows << "\n";
        int oxygen_rate_col = round (oxygen_rate_add / line_no);

        oxygen_rate   += to_string(oxygen_rate_col);
        CO2_rate += to_string(1-oxygen_rate_col);
      }

  }

  int ans = stoi(oxygen_rate, 0, 2) * stoi(CO2_rate, 0, 2);
  cout << "\noxygen   rate = " << oxygen_rate << "\neplison rate = " << CO2_rate  << "\nans          = " << ans << "\n\ndone\n";
}