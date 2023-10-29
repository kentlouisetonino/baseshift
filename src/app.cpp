#include<iostream>
#include "./headers/clearscreen.hpp"
#include "./headers/appdescription.hpp"
#include "./headers/conversionoptions.hpp"
using namespace std;

int main() {
  char option;

  // Description about the app.
  clearScreen();
  cout << "\n";
  appDescription();

  // Description about the options.
  cout << "\n";
  conversionOptions();

  // Ask input based on options.
  cout << "\n";
  cout << "\t Enter the letter of the option: ";
  cin >> option;
  
  // Exit the app.
  cout << "\n";
  return 0;
}

