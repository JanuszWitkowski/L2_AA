import matplotlib.pyplot as plt
import numpy as np
import csv

resT = np.loadtxt('data/exp5b.txt', unpack='False')
res = resT.transpose()

print(len(res))

Ns = list(range(1, 10001))
print(len(Ns))
Ks = [2, 3, 10, 100, 400]

for i in range(0, len(res)):
    plt.plot(Ns, res[i], label="K = " + str(Ks[i]))
    plt.legend(loc='upper right')
    plt.savefig("data/plots/exp5b" + str(Ks[i]), bbox_inches="tight")
    plt.close()


    