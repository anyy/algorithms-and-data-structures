## n個の頂点がリング状に接続されたグラフをCnと記す。グラフC6を図示し、その隣接行列と隣接リストをそれぞれ示しなさい。ただし各辺のコストは1と仮定してよい。
- 隣接行列
  ```
  [0, 1, 0, 0, 0, 1],
  [1, 0, 1, 0, 0, 0],
  [0, 1, 0, 1, 0, 0],
  [0, 0, 1, 0, 1, 0],
  [0, 0, 0, 1, 0, 1],
  [1, 0, 0, 0, 1, 0]
  ```
- 隣接リスト
  ```
  [a] -> [f] -> [b],
  [b] -> [a] -> [c],
  [c] -> [b] -> [d],
  [d] -> [c] -> [e],
  [e] -> [d] -> [f],
  [f] -> [a] -> [e]
  ```
