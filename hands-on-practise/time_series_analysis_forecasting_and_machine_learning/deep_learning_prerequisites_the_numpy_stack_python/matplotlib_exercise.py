import numpy as np
import matplotlib.pyplot as plt

X = np.linspace(-1, 1, 3000)
Y = np.flip(X)

#shuffle the arrays
np.random.shuffle(X)
np.random.shuffle(Y)

flag = np.bitwise_xor(X>0, Y>0)

plt.scatter(X, Y, c=flag)
plt.title("Assignment Done")
plt.show()