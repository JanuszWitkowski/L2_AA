import matplotlib.pyplot as plt
import numpy as np
import csv

res = np.loadtxt('data/exp6.txt', unpack='False')
names = ["Sha1", "Sha3", "Sha256", "Blake2b", "Blake2s", "Md5"]

Bs = [8, 16, 24, 32, 40, 48]

for i in range(0, len(res)):
    plt.plot(Bs, res[i]/10000, label=names[i])
  
plt.xlabel("B")
plt.ylabel("Average Difference")
plt.legend(loc='upper right')
plt.savefig("data/plots/exp6", bbox_inches="tight")
