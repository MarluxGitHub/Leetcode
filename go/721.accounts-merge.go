/*
 * @lc app=leetcode id=721 lang=golang
 *
 * [721] Accounts Merge
 */

// @lc code=start
import (
    "sort"
)

func accountsMerge(accounts [][]string) [][]string {
    /* partitioned DFS but dominated by the sorting, so O(sum(nlogn)) */
    nameEmailGroupMap := make(map[string][][]string)
    res := make([][]string, 0)

    /* partition by name */
    for _, account := range accounts {
        nameEmailGroupMap[account[0]] = append(nameEmailGroupMap[account[0]], account[1:])
    }

    for name, emailGroups := range nameEmailGroupMap {
        /* build the graph for this partition */
        graph := make(map[string][]string)
        for _, emails := range emailGroups {
            first := emails[0]
            for i := 1; i < len(emails); i++ {
                graph[first] = append(graph[first], emails[i])
                graph[emails[i]] = append(graph[emails[i]], first)
            }
        }

        /* dfs to get all connected */
        visited := make(map[string]bool)

        for _, emails := range emailGroups {
            for _, email := range emails {
                connected := ([]string)(nil)
                dfs(email, graph, &connected, visited)

                if connected != nil {
                    sort.Strings(connected)
                    account := []string{name}
                    account = append(account, connected...)
                    res = append(res, account)
                }
            }
        }
    }

    return res
}

func dfs(email string, graph map[string][]string, connected *[]string, visited map[string]bool) {
    if visited[email] == true { return }

    visited[email] = true
    *connected = append(*connected, email)

    neighbors := graph[email]
    if neighbors == nil { return }

    for _, nei := range neighbors {
        if visited[nei] == true { continue }
        dfs(nei, graph, connected, visited)
    }
}

// @lc code=end

