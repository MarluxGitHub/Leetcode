/*
 * @lc app=leetcode id=799 lang=javascript
 *
 * [799] Champagne Tower
 */

// @lc code=start
/**
 * @param {number} poured
 * @param {number} query_row
 * @param {number} query_glass
 * @return {number}
 */
 var champagneTower = function(poured, query_row, query_glass) {
    //build a two dimension array based on row glass
    let flow_through = [];
    flow_through[0] = [poured];
    for( let row=1; row <=query_row; row++) {
        flow_through[row] = new Array(row+1).fill(0);
        for( let glass=0; glass <= row; glass++ ) {
            if(glass <= row -1 && flow_through[row-1][glass] > 1)
                flow_through[row][glass] +=  (flow_through[row-1][glass] -1)/2;
            if(glass-1 >= 0 && flow_through[row-1][glass-1] > 1 ) {
                flow_through[row][glass] +=  (flow_through[row-1][glass-1] -1)/2;
            }

        };
    }

    return flow_through[query_row][query_glass] > 1? 1: flow_through[query_row][query_glass];
};
// @lc code=end

