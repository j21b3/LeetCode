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
    static  vector<vector<int>> fourSum(vector<int>& nums, int target) {
        vector<vector<int>> ret;
        if(nums.size()<4) return ret;
        sort(nums.begin(),nums.end());
        int first,second,third,forth;

        for(first=0;first<nums.size();++first){
            if(first>0&&nums[first]==nums[first-1]) continue;
            for(second=first+1;second<nums.size();++second){
                if(second>first+1&&nums[second]==nums[second-1]) continue;
                int n=target-(nums[first]+nums[second]);
                forth=nums.size()-1;
                for(third=second+1;third<nums.size();++third){
                    if(third>second+1&&nums[third]==nums[third-1]) continue;
                    while (third<forth && nums[third]+nums[forth]>n){
                        --forth;
                    }
                    if(forth==third) break;
                    if(nums[third]+nums[forth]==n) ret.push_back({nums[first],nums[second],nums[third],nums[forth]});
                }

            }
        }
        return ret;
    }
};

int main(){
    vector<int> nums={1,0,-1,0,-2,2};
    vector<vector<int>> ret;
    ret = Solution::fourSum(nums,0);
    for(vector<vector<int>>::iterator i=ret.begin();i!=ret.end();++i){
        cout<<(*i)[0]<<(*i)[1]<<(*i)[2]<<endl;
    }
}

