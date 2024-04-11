#include <iostream>
#include <vector>
#include <cassert>
#include <algorithm> 

using namespace std;

bool classPhotos(vector<int> redShirtHeights, vector<int> blueShirtHeights) {
    auto maxRed  =  max_element(redShirtHeights.begin(), redShirtHeights.end());
    auto maxBlue =  max_element(blueShirtHeights.begin(), blueShirtHeights.end());
    
    if(*maxRed == *maxBlue){
        return false;
    }

    sort(redShirtHeights.begin(), redShirtHeights.end());
    sort(blueShirtHeights.begin(), blueShirtHeights.end());

    if(*maxRed > *maxBlue){
        //red must be behind
        for(int i = 0; i < redShirtHeights.size(); i++){
            if(redShirtHeights[i] <= blueShirtHeights[i]){
                return false ;
            }
        }
    }
    if(*maxBlue > *maxRed){
        // blue must be behind
        for(int i = 0; i < redShirtHeights.size(); i++){
            if(blueShirtHeights[i] <= redShirtHeights[i] ){
                return false ;
            }
        }  
    }

    return true;
}

int main(){
    vector<int> red{5, 6};
    vector<int> blue{5, 4};

    bool answer = classPhotos(red, blue);
    cout << answer << endl; 

    assert(answer == true); 
}