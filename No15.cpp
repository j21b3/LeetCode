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
    static vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> ret;
        if(nums.size()<3) return ret;
        sort(nums.begin(),nums.end());
        int first,second,third;

        for(first=0;first<nums.size();++first){
            if(first>0&&nums[first]==nums[first-1]) continue;
            third=nums.size()-1;
            int n=-nums[first];
            for(second=first+1;second<nums.size();++second){
                if(second>first+1&&nums[second]==nums[second-1]) continue;
                while(second<third && nums[second]+nums[third]>n){
                    --third;
                }
                if(third==second) break;
                if(nums[second]+nums[third]==n) ret.push_back({nums[first],nums[second],nums[third]});
            }
        }
        return ret;
    }
};

int main(){
    vector<int> nums={-1,0,1,2,-1,-4};
    vector<vector<int>> ret;
    ret = Solution::threeSum(nums);
    for(vector<vector<int>>::iterator i=ret.begin();i!=ret.end();++i){
        cout<<(*i)[0]<<(*i)[1]<<(*i)[2]<<endl;
    }
}