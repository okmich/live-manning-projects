import numpy as np

# an array of all zeros
zeros = np.zeros((2,3))

all_ones = np.ones((2,3))

all_tens = 10 * np.ones((2,3))

# identity matrix
id_3 = np.eye(3)
id_10 = np.eye(10)

# random number generation
np.random.random((2,3))

# for specific distributions
np.random.randn(3,4)  # suprisingly doesn't take a tuple
R = np.random.randn(1000, 3)

# stats 
R.mean() #or np.mean(R)
R.var() #or np.var(R)
R.std() #or np.std(R)


# covarian matrix
cov_matrix = np.cov(R.T) # need to transpose the features

# random input
ints_matx = np.random.randint(10, size=(10,10))

# random input
ints_matx = np.random.randint(10, size=(10,10))

#random choice
items = ["BLUE", "RED", "PINK", "GREEN", "PURPLE"]
prob = [0.15, 0.3, 0.1, 0.3, 0.15]
choice_dist_1 = np.random.choice(items, 20, p=prob)
choice_dist_2 = np.random.choice(items, (4, 5), p=prob)




