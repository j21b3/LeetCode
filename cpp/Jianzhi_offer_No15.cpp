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
    static int hammingWeight(uint32_t n) {
        uint32_t i=n;
        int ret=0;
        while (i!=0){
            if(i&1){
                ret+=1;
            }
            i=i>>1;
        }
        return ret;
    }
};

int main(){
    uint32_t n=00000000000000000000000000001011;
    cout<<Solution::hammingWeight(n);
}

