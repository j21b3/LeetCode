#include <iostream>
#include "string"
#include "vector"
#include<algorithm>

using namespace std;

int count_one(int number) {
    int i = number;
    int ret = 0;
    while (i > 0) {
        if (i & 1 == 1) { ret += 1; }
        i=i >> 1;
    }
    return ret;
}

string generate_num(int number, bool hour) {
    string ret = "";
    if (number < 10 && !hour) ret = '0' + to_string(number);
    else ret = ret = to_string(number);
    return ret;
}

vector<string> readBinaryWatch(int turnedOn) {
    vector<string> ret;
    for (int i = 0; i < 16; ++i) {
        int left_turnedOn = turnedOn - count_one(i);
        if (left_turnedOn < 0) continue;
        else if (left_turnedOn == 0) {
            string tmp;
            tmp = generate_num(i, 1) + ":00";
            ret.push_back(tmp);
        } else {
            for (int j = 0; j < 60; ++j) {
                int le_left_turnedOn;
                le_left_turnedOn = left_turnedOn - count_one(j);
                if (le_left_turnedOn != 0) continue;
                else {
                    string tmp;
                    tmp = generate_num(i, 1) + ':' + generate_num(j, 0);
                    ret.push_back(tmp);
                }
            }
        }
    }
    return ret;
}


int main(){
    int n=1;
    vector<string> ret;
    ret = readBinaryWatch(n);
    for(string s:ret){
        cout<<s<<' ';
    }
    return  0;

}