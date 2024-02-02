//
// Created by thyago on 2/2/24.
//

#include <vector>
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

vector<int> sortedSquaredArray(vector<int> array) {

    vector<int> result;
    int i = 0 ;
    int j = array.size() - 1 ;

    while(i <= j){
        int iSq = array[i] * array[i];
        int jSq = array[j] * array[j];

        if(iSq > jSq){
            result.push_back(iSq);
            i++;
        }else{
            result.push_back(jSq);
            j--;
        }

    }
    reverse(result.begin(), result.end());
    return result;

}

int main() {
    vector<int> input{-4, 2, 3, 5, 6, 8, 9};
    vector<int> result  = sortedSquaredArray(input);

    for(int e : result){
        cout << e << " " ;
    }
}