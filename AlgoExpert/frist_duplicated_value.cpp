#include <vector>
#include <iostream>
#include <unordered_map>
using namespace std;

//This solution does not modify the array
int firstDuplicateValueB(vector<int> array) { 
    unordered_map<int, int> aux;

    for(int i = 0; i < array.size(); i++){
        if (aux[array[i]] == 1){
            return array[i];
        }else{
            aux[array[i]] += 1;
        }
    }

    return -1;
}

int firstDuplicateValue(vector<int> array) { 
    for (int i = 0; i < array.size(); i++){
        int index = abs(array[i]) -1;
        if(array[index] < 0){
            return abs(array[i]);
        }else{
            array[index] = array[index] * -1;
        } 
    } 

    return -1;

}


int main(){
    vector<int> input = {2,1,5,2,3,3,4};
    int expectedOutput = 2;
    int output = firstDuplicateValue(input);

    if (expectedOutput != output){
        cout << "Wrong answer" << endl; 
        cout << output << endl;
    }

    cout <<  output ; 
    return 0;

}