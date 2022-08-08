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
private:
    vector<vector<int>> neighbors = {{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}};

public:
    vector<pair<string,int>> step(pair<string,int> status){
        int x = status.first.find('0');
        vector<pair<string,int>> ret;
        for(int y:neighbors[x]){
            swap(status.first[x],status.first[y]);
            pair<string,int> tmp(status.first,status.second+1);
            ret.push_back(tmp);
            swap(status.first[x],status.first[y]);
        }
        return ret;
    }

    int slidingPuzzle(vector<vector<int>>& board) {
        string initial;
        for (int i = 0; i < 2; ++i) {
            for (int j = 0; j < 3; ++j) {
                initial += char(board[i][j] + '0');
            }
        }
        if (initial == "123450") {
            return 0;
        }

        queue<pair<string,int>> q;
        q.emplace(initial,0);
        unordered_set<string> seen = {initial};
        while (!q.empty()){
            pair<string,int> cur=q.front();
            q.pop();
            for(pair<string,int> each:step(cur)){
                if(!seen.count(each.first)){
                    if(each.first=="123450") return each.second;
                    q.emplace(each);
                    seen.insert(move(each.first));
                }
            }
        }
        return -1;
    }
};