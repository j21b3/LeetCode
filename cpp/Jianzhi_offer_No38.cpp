#include <iostream>
#include "string"
#include "vector"
#include<algorithm>
#include "stack"
#include "set"

using namespace std;

class Solution {
public:
    vector<string> permutation(string s) {
        vector<int> flag;
        set<char> dep;
        string select_s="";



        trace_back(s,0,flag,select_s);

        return ret;
    }

private:
    vector<string> ret;

    void trace_back(string s,int position,vector<int> flag,string select_s){
        if(position==s.size()){
            ret.push_back(select_s);
            return;
        }

        set<char> dep;
        for(int i=0;i<s.size();i++){
            if(find(flag.begin(),flag.end(),i)!=flag.end() || dep.count(s[i])){    //对已经确定位置的字符进行跳过 or 再集合中已经遍历过的字符跳过
                continue;
            }
            vector<int> flag_new;
            flag_new=flag;

            dep.insert(s[i]);
            flag_new.push_back(i);
            trace_back(s,position+1,flag_new,select_s+s[i]);
        }
    }
};

int main(){
    Solution sol;
    string s="abc";
    vector<string> ret;
    ret = sol.permutation(s);
    vector<string>::iterator each;
    for(each=ret.begin();each!=ret.end();++each){
        cout<<(*each)<<" ";
    }
}