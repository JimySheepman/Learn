import math

V = 5


def minKey(key, mstSet):
    min = 10000000
    min_index = 0

    for v in range(V):

        if (mstSet[v] == False and
                key[v] < min):
            min = key[v]
            min_index = v

    return min_index


def printMST(parent, n, graph):
    print("Edge Weight")
    minProduct = 1

    for i in range(1, V):
        print("{} - {}   {} ".format(parent[i], i,
                                     graph[i][parent[i]]))
        minProduct *= graph[i][parent[i]]

    print("Minimum Obtainable product is {}".format(
        minProduct))


def primMST(inputGraph, logGraph):
    parent = [0 for i in range(V)]
    key = [10000000 for i in range(V)]
    mstSet = [False for i in range(V)]
    key[0] = 0
    parent[0] = -1
    for count in range(0, V - 1):
        u = minKey(key, mstSet)
        mstSet[u] = True
        for v in range(V):
            if 0 < logGraph[u][v] < key[v] and mstSet[v] == False:
                parent[v] = u
                key[v] = logGraph[u][v]

    printMST(parent, V, inputGraph)


def minimumProductMST(graph):
    logGraph = [[0 for j in range(V)]
                for i in range(V)]

    for i in range(V):
        for j in range(V):
            if (graph[i][j] > 0):
                logGraph[i][j] = math.log(graph[i][j])
            else:
                logGraph[i][j] = 0

    primMST(graph, logGraph)


if __name__ == '__main__':
    graph = [[0, 2, 0, 6, 0],
             [2, 0, 3, 8, 5],
             [0, 3, 0, 0, 7],
             [6, 8, 0, 0, 9],
             [0, 5, 7, 9, 0], ]

    minimumProductMST(graph)
