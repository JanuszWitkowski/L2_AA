import matplotlib.pyplot as plt
import numpy as np
import csv

resT = np.loadtxt('data/exp5aHLL.txt', unpack='False')
res = resT.transpose()

print(len(res))

Ns = list(range(1, 10001))
print(len(Ns))
Ks = [1, 2, 3, 4,]

for i in range(0, len(res)):
    plt.scatter(Ns, res[i], label="M = " + str(Ks[i]), s=2)
    plt.legend(loc='upper right')
    plt.savefig("data/plots/exp5aHLL" + str(Ks[i]), bbox_inches="tight")
    plt.close()


    