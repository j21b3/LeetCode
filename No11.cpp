#include <iostream>
#include "string"
#include "vector"
#include<algorithm>
#include "stack"
#include "set"

using namespace std;
int maxArea(vector<int>& height) {
    vector<int>::iterator before,after;
    before=height.begin();
    after = height.end()-1;

    int maxsize=0;

    while(before!=after){
        int size = after-before;
        if((*before)>(*after)){
            size=size*(*after);
            after-=1;
        } else{
            size=size*(*before);
            before+=1;
        }
        maxsize= (maxsize<size) ? size : maxsize;
    }
    return maxsize;
}

int main(){
    vector<int> ve;
    ve.push_back(4);
    ve.push_back(3);
    ve.push_back(2);
    ve.push_back(1);
    ve.push_back(4);
    cout<<maxArea(ve)<<endl;
}
