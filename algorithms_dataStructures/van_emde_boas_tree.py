import math


class VEBtree:
    def __init__(self, u):
        if u < 0:
            raise Exception("Universe size cannot be less than zero. u = " + str(u))
        self.min = None
        self.max = None
        self.u = 2
        while self.u < u:
            self.u = self.u * self.u

        if u > 2:
            # define the cluster vectors
            # number of clusters 0, 1, 2, ..., upper sqrt(u) - 1
            self.clusters = [None for i in range(self.high(self.u))]
            # define the summary vectors
            self.summary = None

    def getMin(self):
        """
        Returns cached min in the universe in O(1)
        """
        # if self.min is not None:
        #	return int(self.min)
        return self.min

    def getMax(self):
        """
        Returns cached max in the universe in O(1)
        """
        # if self.max is not None:
        #	return int(self.max)
        return self.max

    def high(self, x):
        """
        Returns the cluster number of element x
        """
        return int(math.floor(x / math.sqrt(self.u)))

    def low(self, x):
        """
        Return the position inside cluster of element x
        """
        return int(x % math.ceil(math.sqrt(self.u)))

    def index(self, x, y):
        """
        Returns index from cluster number and position within cluster
        """
        return int(x * math.floor(math.sqrt(self.u)) + y)

    def isMember(self, x):
        """
        Find if x is a member of the tree in O(lg lg u) time
        """
        if self.min == x or self.max == x:
            return True
        elif self.u <= 2:
            return False
        else:
            cluster_of_x = self.clusters[self.high(x)]
            if cluster_of_x is not None:
                return cluster_of_x.isMember(self.low(x))
            else:
                return False

    def getPredecessor(self, x):
        if self.u <= 2:
            if x == 1 and self.getMin() == 0:
                return 0
            else:
                return None

        elif (self.getMax() is not None) and (x > self.getMax()):
            return self.getMax()

        else:
            high = self.high(x)
            low = self.low(x)
            high_cluster = self.clusters[high]
            min_low = None
            if high_cluster is not None:
                min_low = high_cluster.getMin()
            if (min_low is not None) and (low > min_low):
                offset = high_cluster.getPredecessor(low)
                if offset is not None:
                    return self.index(self.high(x), offset)
                else:
                    offset = 0
                    return self.index(self.high(x), offset)

            else:
                pred_cluster = None
                if self.summary is not None:
                    pred_cluster = self.summary.getPredecessor(high)

                if pred_cluster is None:
                    if (self.getMin() is not None) and (x > self.getMin()):
                        return self.getMin()
                    else:
                        return None
                else:
                    if self.clusters[pred_cluster] is not None:
                        offset = self.clusters[pred_cluster].getMax()
                    if offset is not None:
                        return self.index(pred_cluster, offset)
                    else:
                        offset = 0
                        return self.index(pred_cluster, offset)

    def getSuccessor(self, x):
        if self.u <= 2:

            if x == 0 and self.getMax() == 1:
                return 1

            else:
                return None
        elif (self.getMin() is not None) and (x < self.getMin()):
            return self.getMin()
        else:
            high = self.high(x)
            low = self.low(x)
            high_cluster = self.clusters[high]
            max_low = None
            if high_cluster is not None:
                max_low = high_cluster.getMax()

            if (max_low is not None) and (low < max_low):
                offset = high_cluster.getSuccessor(low)
                # if offset is not None:
                return self.index(self.high(x), offset)
            # else:
            #	return None
            else:
                succ_cluster = None
                if self.summary is not None:
                    succ_cluster = self.summary.getSuccessor(high)
                if succ_cluster is None:
                    return None
                else:
                    offset = 0
                    if self.clusters[succ_cluster] is not None:
                        offset = self.clusters[succ_cluster].getMin()
                    return self.index(succ_cluster, offset)

    def emptyVEBTreeInsert(self, x):
        """
        First insert, updates min and max only
        """
        self.min = x
        self.max = x

    def vebTreeInsert(self, x):
        """
        Inserts an element x in the universe of size u
        """
        if self.min is None:
            self.emptyVEBTreeInsert(x)
        else:
            if x < self.min:
                temp = x
                x = self.min
                self.min = temp

            if self.u > 2:
                cluster_id_x = self.high(x)

                if self.clusters[cluster_id_x] is None:
                    # create a new cluster
                    self.clusters[cluster_id_x] = VEBtree(self.high(self.u))
                if self.summary is None:
                    self.summary = VEBtree(self.high(self.u))
                if self.clusters[cluster_id_x].min is None:
                    self.summary.vebTreeInsert(cluster_id_x)
                    self.clusters[cluster_id_x].emptyVEBTreeInsert(self.low(x))
                else:
                    self.clusters[cluster_id_x].vebTreeInsert(self.low(x))
            if x > self.max:
                self.max = x

    def vebTreeDelete(self, x):
        """
        Deletes an element x in the universe of size u
        """
        if self.min == self.max:
            self.min = None
            self.max = None
        elif self.u == 2:
            if x == 0:
                self.min = 1
            else:
                self.min = 0

            self.max = self.min
        else:
            if x == self.min:
                first_cluster = self.summary.getMin()
                x = self.index(first_cluster, self.clusters[first_cluster].getMin())
                self.min = x

            self.clusters[self.high(x)].vebTreeDelete(self.low(x))

            if self.clusters[self.high(x)].getMin() is None:
                self.summary.vebTreeDelete(self.high(x))

                if x == self.max:
                    summary_max = self.summary.getMax()
                    if summary_max is None:
                        self.max = self.min
                    else:
                        self.max = self.index(summary_max, self.clusters[summary_max].getMax())
            elif x == self.max:
                self.max = self.index(self.high(x), self.clusters[self.high(x)].getMax())