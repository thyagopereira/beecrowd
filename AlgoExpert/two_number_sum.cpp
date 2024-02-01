#include <vector>
#include <map>
using namespace std;

vector<int> twoNumberSum(vector<int> array, int targetSum) {

    map<int, int> aux;
    for (int i=0 ; i < array.size(); i++){
        aux[array[i]] += 0 ;
    }

    int val;
    vector<int> result;
    for (int j = 0 ; j < array.size(); j++){
        val = targetSum - array[j];
        if (!( aux.find(val) == aux.end()) && (val != array[j])){
            result.push_back(val);
            result.push_back(array[j]);

            break;
        }
    }

    return result;
}