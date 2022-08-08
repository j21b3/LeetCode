#include <iostream>
#include "string"
#include "vector"
#include<algorithm>
#include "stack"
#include "set"
#include "map"
#include "queue"
#include "unordered_set"

using namespace std;

class Solution {

public:


    int snakesAndLadders(vector<vector<int>> &board) {
        int target = board.size() * board[0].size();
        reverse(board.begin(),board.end());
        for (int i = 0; i < board.size(); i++) {
            if (i % 2 == 1) {
                reverse(board[i].begin(), board[i].end());
            }
        }
        vector<int>v;
        for (int i = 0; i < board.size(); i++) {
            for (int j = 0; j < board[0].size(); j++) {
                v.push_back(board[i][j]);
            }
        }
        unordered_set<int>s;
        queue<int>q;
        s.insert(1);
        q.push(1);
        int step = 0;
        int n = 1;
        while (!q.empty()) {
            if (n == 0) {
                n = q.size();
            }
            int cur = q.front();
            q.pop();
            for (int i = 1; i <= 6; i++) {
                int temp = cur + i;
                if (temp == target) {
                    return step + 1;
                }
                if (temp > target) {
                    break;
                }
                if (v[temp - 1] > 0) {
                    temp = v[temp - 1];
                    if (temp == target) {
                        return step + 1;
                    }
                    if (!s.count(temp)) {
                        s.insert(temp);
                        q.push(temp);
                    }
                }
                if (!s.count(temp)) {
                    s.insert(temp);
                    q.push(temp);
                }
            }
            if (--n == 0) {
                step++;
            }
        }
        return -1;
    }

};
