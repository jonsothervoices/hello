#include <iostream>
// Definition for a binary tree node.
struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};
using namespace std;
void printBST(TreeNode* root) {
        if (!root) {
            cout << "Root is NULL" << endl;
        }
        cout << "Root " << root->val;
        if (root->left) {
            cout << " has left " << root->left->val;
        }
        else {
            cout << " has no left";
        }
        cout << " and";

        if (root->right) {
            cout << " has right " << root->right->val;
        }
        else {
            cout << " has no right";
        }
        cout << endl;

        if (root->left) {
            cout << "ON THE LEFT:" << endl;
            TreeNode* curr = root->left;
            printBST(curr);
        }
        if (root->right) {
            cout << "ON THE RIGHT:" << endl;
            TreeNode* curr = root->right;
            printBST(curr);
        }
    }
