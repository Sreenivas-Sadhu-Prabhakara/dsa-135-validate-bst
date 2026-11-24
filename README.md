# Validate a BST

> Striver SDE Sheet 135 · Day 20 BST — pass each node the (lo, hi) window its ancestors imply; comparing only parent-child is the famous trap. Go, stdlib only, tested.

## Problem

Is a binary tree a valid BST — every node strictly greater than *all* of its left subtree and strictly smaller than *all* of its right?

## Approach

The trap this problem was built around: checking each node against only its children accepts trees like `5 → right 4 → left 3`, where the violation is against a *grandparent*. The fix is to carry the constraint down: every node must lie inside an open interval (lo, hi) — descending left shrinks hi to the node's value, descending right raises lo. The interval accumulates every ancestor's requirement, so one comparison per node suffices. Equivalent alternative: an inorder walk must produce strictly increasing values; the bounds version is shown because it generalizes to "largest BST in a tree" (#144).

**Complexity:** O(n) time, O(h) recursion.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
