#include <iostream>
#include "string"
#include "vector"
#include<algorithm>
#include "stack"
#include "set"
#include "map"
using  namespace std;


class Solution {


private:
    static map<char,int> romanTable;
public:
    static int romanToInt(string s) {
        int n=0;
        for(int i=0;i<s.length();++i){
            int tmp=romanTable[s[i]];
            n = (i+1<s.length() && tmp<romanTable[s[i+1]]) ? n - romanTable[s[i]] : n + romanTable[s[i]];
        }
        return n;
    }
};

map<char,int> Solution::romanTable= {
        {'I',1},{'V',5},{'X',10},{'L',50},
        {'C',100},{'D',500},{'M',1000},
};


int main(){
    string s="III";
    cout<<Solution::romanToInt(s);
}
