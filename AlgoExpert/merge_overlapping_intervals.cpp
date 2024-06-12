#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;
    
vector<vector<int>> mergeOverlappingIntervals(vector<vector<int>> intervals) {
  sort(intervals.begin(), intervals.end(), 
  [](const vector<int>& a, const vector<int>& b){
    return a[0] < b[0];
  });

  vector<vector<int>> mergedIntervals = {intervals[0]};
  int j = 0;

  for(int i = 1; i< intervals.size(); i++){
    if(intervals[i][0] <= mergedIntervals[j][1]){
        if(intervals[i][1] > mergedIntervals[j][1]){
          mergedIntervals[j][1] = intervals[i][1];
        }       
    }else{
        mergedIntervals.push_back(intervals[i]);
        j++;
    }
    
  }
 
  return mergedIntervals;
}

int main(){

    vector<vector<int>> input = {{1,2}, {3,5}, {4,7}, {6,8}, {9, 10}};
    vector<vector<int>> output = mergeOverlappingIntervals(input);
    vector<vector<int>> expectedOutput = {{1,2}, {3,8}, {9, 10}};
    
    for(int i = 0; i < expectedOutput.size(); i++){
        for(int j = 0; j < expectedOutput[i].size(); j++){
            if(expectedOutput[i][j] != output[i][j]){
                cout << "Wrong answer" << endl;
            }
        }
    };

    for (int i=0; i< output.size(); i++){
        cout << output[i][0] << "----" << output[i][1] << endl;
    }

    return 0; 

}