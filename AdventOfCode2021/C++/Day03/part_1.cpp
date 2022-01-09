#include <iostream>
#include <fstream>
#include <string>
#include <math.h> 

using namespace std;


int main ()
{

  string gamma_rate;
  string epsilon_rate;

  ifstream file("inputs.txt");
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

  cout << "file has " << num_of_cols << " columns\n";

  for (int col_ind = 0; col_ind < num_of_cols; col_ind++) 
  {
    ifstream file("inputs.txt");
    int gamma_rate_add = 0;
    double line_no = 1;
    string sLine;
    while (getline(file, sLine)) 
      {
        string values = sLine;
        string value = values.substr(col_ind,1);
        gamma_rate_add += stoi(value);
        ++line_no;
      
      }
      cout << gamma_rate_add << " " << sLine << "\n";
      int gamma_rate_col = round (gamma_rate_add / line_no);

      gamma_rate   += to_string(gamma_rate_col);
      epsilon_rate += to_string(1-gamma_rate_col);

  }

  int ans = stoi(gamma_rate, 0, 2) * stoi(epsilon_rate, 0, 2);
  cout << "\ngamma   rate = " << gamma_rate << "\neplison rate = " << epsilon_rate  << "\nans          = " << ans << "\n\ndone\n";
}