import numpy as np

# in python matrices are list of list like [[1,2,3], [4,5,6]]
L = [[1,2,3], [4,5,6]]

# in numpy
A = np.array([[1,2,3], [4,5,6]])

# transpose matrics
A.T

# math operations like exp or sqrt
np.exp(A)
np.sqrt(A)

# matrics multiplicaton
B = np.array([[2,4], [5,8]])

A.dot(B) # fails - shapes (2,3) and (2,2) not aligned: 3 (dim 1) != 2 (dim 0) 
B.dot(A) #works
B @ A # also works - the @ operation is a matrics multiplication



#some linear algebra options

#inverse of matrics
np.linalg.inv(A) # fails - numpy.linalg.LinAlgError: Last 2 dimensions of the array must be square
np.linalg.inv(B) # works

# inverse of inverse is the original
np.linalg.inv(np.linalg.inv(B))

# inverse multiple original is identity
np.linalg.inv(B).dot(B)


# generate new array 
fives = np.reshape(np.arange(1,26), (5,5))

# diagonal value of matrix
np.diag(fives)
fives.diagonal()

# trace is the sum of the diagonal vector
np.trace(fives)
np.diagonal().sum()


#eigenvectors
np.
