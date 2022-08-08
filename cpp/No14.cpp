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
    static string longestCommonPrefix(vector<string>& strs) {
        vector<string>::iterator each;
        int i=0;
        while(i<strs[0].length()){

            char c = strs[0][i];
            for(each=strs.begin();each!=strs.end();++each){

                if(i>=(*each).length() || (*each)[i]!=c) return (*each).substr(0,i);

            }
            i+=1;
        }
        return strs[0].substr(0,i);
    }
};

int main(){
    vector<string> a;
    string t ;
    t= "a";
    a.push_back(t);

    cout<<Solution::longestCommonPrefix(a);

}