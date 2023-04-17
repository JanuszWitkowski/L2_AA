import matplotlib.pyplot as plt
import numpy as np
import csv

res = np.loadtxt('data/compare.txt', unpack='False')


print(len(res))

Ns = list(range(1, 10001))
print(len(Ns))
# Ks = [2, 5, 7, 10, 16]


# for i in range(0, len(res)):
plt.scatter(Ns, res[0], label="HyperLogLog", s=2)
plt.scatter(Ns, res[1], label="MinCount", s=2)
plt.legend(loc='upper right')
plt.savefig("data/plots/compare", bbox_inches="tight")
plt.close()

plt.scatter(Ns, res[1], label="MinCount", s=2)
plt.legend(loc='upper right')
plt.savefig("data/plots/compareMinCount", bbox_inches="tight")
plt.close()

plt.scatter(Ns, res[0], label="HyperLogLog", s=2)
plt.legend(loc='upper right')
plt.savefig("data/plots/compareHyperLogLog", bbox_inches="tight")
plt.close()


print("\n")
print("HyperLogLog\t\t MinCount")
print(np.mean(res[0]), "\t", np.mean(res[1]))
print(np.var(res[0]), "\t", np.var(res[1]))