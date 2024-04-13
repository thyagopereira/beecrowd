#include <vector>
#include <iostream>
#include <bits/stdc++.h>
using namespace std; 

int tandemBicycle(
  vector<int> redShirtSpeeds, vector<int> blueShirtSpeeds, bool fastest
) {
  if (fastest == true){
    // Precisa juntar os maiores de um + os menores de outro
    sort(redShirtSpeeds.begin(), redShirtSpeeds.end()); // sort in increasing order
    sort(blueShirtSpeeds.begin(), blueShirtSpeeds.end(), greater<int>()); // Sort in decreasing order
  
    int sum = 0; 
    for(int i = 0; i < redShirtSpeeds.size(); i++){
      sum += max(redShirtSpeeds[i], blueShirtSpeeds[i]);
    }

    return sum; 
  }else {
    // Precisa juntar os menores com os menores
    sort(redShirtSpeeds.begin(), redShirtSpeeds.end());
    sort(blueShirtSpeeds.begin(), blueShirtSpeeds.end());


    int sum = 0;
    for(int i = 0; i < redShirtSpeeds.size(); i++){
      sum += max(redShirtSpeeds[i], blueShirtSpeeds[i]);
    }

    return sum;
  } 
}

int main(){
  vector<int> redSpeed{5, 5, 3, 9, 2};
  vector<int> blueSpeed{3, 6, 7, 2, 1};
  bool fastest = true; 

  int output = tandemBicycle(redSpeed, blueSpeed, fastest);
  cout << output << endl;

  assert(output == 32); 
}