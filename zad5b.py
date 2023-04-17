import matplotlib.pyplot as plt
import numpy as np
import csv

resT = np.loadtxt('data/exp5bHLL.txt', unpack='False')
res = resT.transpose()

print(len(res))

Ns = list(range(1, 10001))
print(len(Ns))
Ks = [2, 5, 7, 10, 16]

for i in range(0, len(res)):
    plt.scatter(Ns, res[i], label="K = " + str(Ks[i]), s=2)
    plt.legend(loc='upper right')
    plt.savefig("data/plots/exp5bHLL" + str(Ks[i]), bbox_inches="tight")
    plt.close()


    