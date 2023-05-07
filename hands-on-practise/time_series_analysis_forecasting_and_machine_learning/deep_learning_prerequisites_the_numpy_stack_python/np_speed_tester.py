from datetime import datetime

import math
import random

import numpy as np

def make_matrix(values, row, col):
	result = []
	for i in range(row):
		axis1 = []
		for j in range(col):
			axis1.append(values[i*col+j])
		result.append(axis1)
	return result

def python_dot_product(arr1, arr2):
	result = 0
	for i in range(len(arr1)):
		result += arr1[i] * arr2[i]
		
	return result

def python_matrix_mul(arr1, arr2):
	result = [[sum(a*b for a,b in zip(X_row,Y_col)) for Y_col in zip(*arr2)] for X_row in arr1]
	return result


def numpy_matrix_mul(arr1, arr2):
	return np.dot(arr1, arr2)

n = 100000
size = 150
rows = 50
cols = 3
print("Prepare data")
py_a =  make_matrix([random.randint(0, 255) for i in range(size)], rows, cols)
py_b =  make_matrix([random.randint(0, 255) for i in range(size)], cols, rows)

np_matx_a = np.array(py_a)
np_matx_b = np.array(py_b)

start_py = datetime.now()
for  i in range(n):
	res = python_matrix_mul(py_a, py_b)

end_py = datetime.now()

start_np = datetime.now()
for  i in range(n):
	res = numpy_matrix_mul(np_matx_a, np_matx_b)

end_np = datetime.now()


print("++++++++++++++ results +++++++++++++++++++")
diff_py = (end_py - start_py).total_seconds()
print(f"Python list took {diff_py}")

diff_np = (end_np - start_np).total_seconds()
print(f"numpy list took {diff_np}")

print(f"the difference is in magnitude of {diff_py/diff_np}")