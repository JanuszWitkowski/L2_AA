import matplotlib.pyplot as plt
import numpy as np
import csv


deltaCheb = [0.22360679774997896, 0.5, 0.7071067811865476]
deltaCher = [0.13813682127470314, 0.16976274088460064, 0.18274403080607504]
deltaMy = [0.0908926694692337, 0.12243496324827381, 0.13204973060982775]
alpha = ["5%", "1%", "0.5%"]

resT = np.loadtxt('data/exp5b.txt', unpack='False')
res = resT.transpose()
Ns = list(range(1, 10001))

for i in range(0, len(alpha)):
    plt.scatter(Ns, res[4], s=2)
    plt.axhline(y = 1 - deltaCheb[i], label = "Chebyschev", c='r')
    plt.axhline(y = 1 + deltaCheb[i], label = "Chebyschev", c='r')
    plt.axhline(y = 1 - deltaCher[i], label = "Chernoff", c='b')
    plt.axhline(y = 1 + deltaCher[i], label = "Chernoff", c='b')
    plt.axhline(y = 1 - deltaMy[i], label = "Experimental", c='g')
    plt.axhline(y = 1 + deltaMy[i], label = "Experimental", c='g')
    plt.legend(loc='upper right')
    plt.savefig("data/plots/exp7" + str(i), bbox_inches="tight")
    plt.close()
