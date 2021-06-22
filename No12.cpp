#include <iostream>
#include "string"
#include "vector"
#include<algorithm>
#include "stack"
#include "set"
using  namespace std;

const pair<int,string> valueSymbles[] = {
        {1000, "M"},{900,  "CM"},{500,  "D"},{400,  "CD"},
        {100,  "C"},{90,   "XC"},{50,   "L"},{40,   "XL"},
        {10,   "X"},{9,    "IX"},{5,    "V"},{4,    "IV"},
        {1,    "I"},
};

class Solution {
public:


    static string intToRoman(int num) {
        int i=num;
        string ret;
        for(const auto &[n,roman]:valueSymbles){
            while(i>=n){
                i-=n;
                ret+=roman;
            }
            if(i==0) break;
        }
        return ret;
    }
};

int main(){
    int n=4;
    cout<<Solution::intToRoman(n);
}