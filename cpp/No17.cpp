#include <iostream>
#include "string"
#include "vector"
#include<algorithm>
#include "stack"
#include "set"
#include "map"
using  namespace std;

class Solution {
public:
    map<int,string> buttonMap={
            {2,"abc"},{3,"def"},{4,"ghi"},
            {5,"jkl"},{6,"mno"},{7,"pqrs"},
            {8,"tuv"},{9,"wxyz"}
    };

    vector<string> lettercom_Recursion(char num, vector<string> s)
    {
        string whole=buttonMap[num-'0'];
        vector<string> ret;
        for(vector<string>::iterator each=s.begin();each!=s.end();++each){
            for(int i=0;i<whole.length();++i){
                ret.push_back((*each)+whole[i]);
            }
        }
        return move(ret);
    }

    vector<string> letterCombinations(string digits) {
        vector<string> ret;
        if(digits.length()==0) return ret;
        for(int i=0;i<buttonMap[digits[0]-'0'].length();++i){
            string tmp=buttonMap[digits[0]-'0'].substr(i,1);
            ret.push_back(tmp);
        }
        for(int i=1;i<digits.length();++i){
            ret = lettercom_Recursion(digits[i],ret);
        }
        return move(ret);
    }
};

int main(){
    Solution s;
    string t = "23";
    vector<string> ret;
    ret = s.letterCombinations(t);
}