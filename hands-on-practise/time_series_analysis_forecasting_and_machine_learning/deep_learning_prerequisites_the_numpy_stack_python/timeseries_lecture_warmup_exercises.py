import numpy as np
import matplotlib.pyplot as plt
from scipy.stats import multivariate_normal

###### Exercise #1

#using numpy, generate 1000 samples from the standard normal 
arr = np.random.randn(1000)
#plot the result as a time series
plt.plot(arr)
plt.show()
#plot the result as a histogram
plt.hist(arr, bin=20)
plt.show()


###### Exercise #2

#add a trend line to the noise

#make a scatterplot


#Bonus: find and plot the best-fit line



###### Exercise #3
#call the np.cumsum() function on your noise
arr_cum = np.cumsum(arr)
#plot the result
plt.plot(arr_cum)
plt.show()
#what do it remind you of?
print("market data")


###### Exercise #4
# Generate and plot 1000 samples from the multivariate normal
mean = [0,0]
sigma = [[1, -0.5], [-0.5, 2]]
multi_norm = multivariate_normal(mean=mean, cov=sigma)
samples = multi_norm.rvs(size=1000)
plt.scatter(samples[:,0], samples[:,1])
plt.show()



###### Exercise #5
# Calcualte the sample mean and sample covariance of the data you just generated - are they close to what you expect?
# Don't use np.mean() or np.cov()
mean = samples.sum(axis=0) / 1000
