#include <iostream>
#include "string"
#include "vector"
#include<algorithm>
#include "stack"
#include "set"
#include "map"
#include "queue"
#include "unordered_set"
using  namespace std;



class Solution {
public:
    string convertToTitle(int columnNumber) {
        string ret;
        while(columnNumber>0){
            int a = (columnNumber-1)%26;
            char c='A'+a;
            columnNumber=(columnNumber-a-1)/26;
            ret+=c;
        }
        reverse(ret.begin(),ret.end());
        return ret;
    }
};