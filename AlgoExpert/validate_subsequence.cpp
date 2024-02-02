#include <vector>
#include <iostream>
using namespace std;

bool isValidSubsequence(vector<int> array, vector<int> sequence) {
    int j = 0;
    for (int i =0; i < array.size(); i++){
        if (array[i] == sequence[j]){
            j++;
        }
    }

    if (j < sequence.size()){
        return false;
    }

    return true;
}


int main() {
    vector<int> array{5, 1, 22, 25, 6, -1, 8, 10};
    vector<int> sequence{1, 6, -1 ,10};

    bool result = isValidSubsequence(array, sequence);

    cout << result ;
}