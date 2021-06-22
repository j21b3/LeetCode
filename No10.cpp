//
// Created by 84200 on 2021/6/21.
//
//
// Created by 84200 on 2021/6/21.
//
#include <iostream>
#include "string"
#include "vector"
#include<algorithm>
#include "stack"

using namespace std;

struct index{
    int i=0;
    int j=0;
};

stack<index> star_operation(vector<vector<bool>> dp,int index_i,int index_j){
    stack<index> ret;
    index node;
    if(index_i+1<dp.size()){    //判断往下走的时候会不会越界
        node.i=index_i+1;
        node.j=index_j;
        if(dp[node.i][node.j]) ret.push(node); //判断匹配0个字符的时候是否能够继续，能则加入
    }
    for(int j=index_j+1;j<dp[index_i].size();++j){
        if(!dp[index_i][j]) break;
        node.i=index_i;
        node.j=j;
        ret.push(node);

    }
    return ret; //有符合条件的节点则返回他们的编号，没有的话则返回空
}

stack<index> normal_operation(vector<vector<bool>> dp,int index_i,int index_j){
    stack<index> ret;
    index node;
    if(index_i+1<dp.size() && index_j+1<dp[0].size()){
        if(dp[index_i+1][index_j+1]){
            node.i=index_i+1;
            node.j=index_j+1;
        }
        ret.push(node);
    }
    return ret;//有符合条件的节点则返回他们的编号，没有的话则返回空
}

bool isMatch(string s, string p) {
    vector<string> P;   //初始化P容器，用以存储拆分后的正则表达式
    for(char c:p){
        string tmp;
        if(c!='*'){
            tmp=c;
            P.push_back(tmp);
        } else{
            tmp=P.back()+'*';
            P.pop_back();
            P.push_back(tmp);
        }
    }

    vector<string>::iterator each;
    /*
    for(each=P.begin();each!=P.end();++each){
        cout<<*each<<' ';
    }
    */
    //建立DP表
    vector<vector<bool>> dp;

    for(each=P.begin();each!=P.end();++each){
        string a;
        a = *each;
        char p1=a[0];
        vector<bool> tmp;
        for(char c:s){
            if(c==p1 || p1=='.') tmp.push_back(1);
            else tmp.push_back(0);
        }
        dp.push_back(tmp);
    }
/*
    vector<vector<bool>>::iterator dpi;
    vector<bool>::iterator dpj;
    for(dpi=dp.begin();dpi!=dp.end();++dpi){
        for(dpj=(*dpi).begin();dpj!=(*dpi).end();++dpj){
            cout<<(*dpj)<<' ';
        }
        cout<<endl;
    }

    vector<index> ret= star_operation(dp,2,2);
    vector<index>::iterator a;
    for(a=ret.begin();a!=ret.end();++a){
        cout<<'('<<(*a).i<<','<<(*a).j<<')'<<' ';
    }
*/
    stack<stack<index>> step_stack;
    stack<index> ret;
    //由于.只影响dp矩阵，因此.*的情况就是一行全是1的dp矩阵的*操作
    /*
    if(P[0].length()==2){   //此时第0个为*操作,*操作的时候dp[0][0]可能为0
        ret = star_operation(dp,0,0);
        if(!ret.empty()){
            step_stack.push(ret);
        } else return 0;
    } else if(!dp[0][0]) return 0;  //不是*操作的时候，dp【0】【0】为0的时候表示匹配失败
    else{   //不是*操作，不是0，把0，0压入栈
        index tmp;
        tmp.i=0;tmp.j=0;
        ret.push(tmp);
        step_stack.push(ret);
    }*/
    index tmp;
    tmp.i=0;tmp.j=0;
    ret.push(tmp);
    step_stack.push(ret);
    int i=0;
    int j=0;
    int x=dp.size()-1;
    int y=dp[0].size()-1;
    while (!step_stack.empty()){
        index node = step_stack.top().top();
        i=node.i;
        j=node.j;
        stack<index> ret;
        if(P[i].length()==2){//表示*操作:先扩展，再普通操作
            ret = star_operation(dp,i,j);
            if(ret.empty()){    //要是返回值空，说明该节点无法扩展，则弹出栈顶的栈顶
                while(!step_stack.empty() && step_stack.top().size()==1){
                    step_stack.pop();
                }
                if(step_stack.empty()) break;
                step_stack.top().pop();
            }else{  //不为空，把这个加入栈
                //判断有没有到x，y 到了说明匹配成功
                stack<index> tmp;
                while(!ret.empty()){
                    if (ret.top().i == x && ret.top().j == y) return 1;
                    else {
                        tmp.push(move(ret.top()));
                        ret.pop();
                    }
                }
                step_stack.push(tmp);
            }
        }
        node = step_stack.top().top();
        i=node.i;
        j=node.j;
        ret = normal_operation(dp,i,j);

        if(ret.empty()){    //要是返回值空，说明该节点无法扩展，则弹出栈顶的栈顶,直到遇到一个有多个元素的栈顶
            while(!step_stack.empty() && step_stack.top().size()==1 ){
                step_stack.pop();
            }
            if(step_stack.empty()) break;
            step_stack.top().pop();

        }else{  //不为空，把这个加入栈
            //判断有没有到x，y 到了说明匹配成功
            stack<index> tmp;
            while(!ret.empty()){
                if (ret.top().i == x && ret.top().j == y) return 1;
                else {
                    tmp.push(move(ret.top()));
                    ret.pop();
                }
            }
            step_stack.push(tmp);
        }
    }
    return 0;
}


int main(){
    string s,p;
    //s = "aab";
    //p = "c*a*b";
    s = "aa";
    p = "a*";
    cout<<isMatch(s,p);
    return 0;
}

