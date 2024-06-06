#include <vector>
#include <iostream>

using namespace std;

vector<int> arrayOfProducts(vector<int> array) {
	vector<int> left;
	vector<int> right;
	for (int i = 0; i < array.size(); i++){
		left.push_back(1);
		right.push_back(1);
	}

	int product = 1;
	for (int i = 0; i < array.size(); i++){
		left[i] = product; 
		product *= array[i]; 
	}

	product = 1; 
	for (int j = array.size() -1; j >= 0; j--){
		right[j] = product;
		product *= array[j];
	}

	vector<int> result; 
	for (int k = 0; k < array.size(); k++){
		result.push_back(left[k] * right[k]);
	}
	return result;
}

int main(){
	vector<int> input = vector<int>{5,1,4,2} ; 
	vector<int> expectedOutput = vector<int>{8, 40, 10, 20};
	vector<int> output = arrayOfProducts(input);

	for(int i=0 ; i < expectedOutput.size(); i++){
		if (expectedOutput[i] != output[i]){
			std::cout << "Wrong Answer" << endl;
			exit(1);
		}
	}
}
