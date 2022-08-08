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
    static int maxPoints(vector<vector<int>>& points) {
        int len=points.size();
        if(len<3) return len;
        int MaxNum=2;
        for(int i=0;i<len;++i){ //二重循环便利每两个点
            map<double,int> count;
            for(int j=i+1;j<len;++j){
                long long dx=points[i][0]-points[j][0];
                long long dy=points[i][1]-points[j][1];
                double g=dy*1.0/dx;
                if(g==1.0/0 || g==-1.0/0) g=1.0/0;  //c++中inf和-inf不相等，因此对这两个要按一类进行处理
                if(count.count(g)) count[g]++;
                else count[g]=2;
                MaxNum= max(MaxNum,count[g]);

            }
        }
        return MaxNum;
    }
};
int main(){
    vector<vector<int>> po={{0,1},{0, 0}, {0, 4}, {0, -2}, {0, -1}, {0, 3}, {0, -4}};
    cout<<Solution::maxPoints(po);
}