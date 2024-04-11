#include <iostream>
#include <vector>
#include <bits/stdc++.h>
using namespace std;


int minimumWaitingTime(vector<int> queries) {
    int total = 0 ;

    sort(queries.begin(), queries.end());

    for(int i = 0; i < queries.size(); i++){
        int queriesLeft = queries.size() - (i + 1);
        total += queries[i] * queriesLeft;
    }

    return total;
}

int main(){
    vector<int> input{3, 2, 1, 2, 6};

    int output = minimumWaitingTime(input);
    cout << output << endl;

    assert(output == 17);
}