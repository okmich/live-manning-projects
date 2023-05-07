import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

from statsmodels.tsa.stattools import acf, pacf, pacf_ols, pacf_yw
from statsmodels.graphics.tsaplots import plot_acf, plot_pacf

# using random number
print("Using random number")
rand_number = np.random.randn(400)
plt.plot(rand_number)
plt.show()

print('\n=>Calculate the autocorrelation function.\n')
print(acf(rand_number))

print('\n=>Partial autocorrelation estimate.\n')
print(pacf(rand_number))

print('\n=>Partial autocorrelation estimated via OLS.\n')
print(pacf_ols(rand_number))

print('\n=>Partial autocorrelation estimated with non-recursive yule_walker\n')
print(pacf_yw(rand_number))

#plot
plot_acf(rand_number, title='Partial autocorrelation estimate.')
plt.show()

plot_pacf(rand_number, title='Partial auto-correlation function of random number')
plt.show()




print("Using a generated series")
# generated series
import math
ys = [rand_number[0], rand_number[1], rand_number[2]]
for i in range(3, len(rand_number)):
	y = 0.8 * rand_number[i-2] - 0.45 * rand_number[i-1] + 0.1*rand_number[i]
	ys.append(y)


ys = np.array(ys)

print('\n=>Calculate the autocorrelation function.\n')
print(acf(ys))

print('\n=>Partial autocorrelation estimate.\n')
print(pacf(ys))

print('\n=>Partial autocorrelation estimated via OLS.\n')
print(pacf_ols(ys))

print('\n=>Partial autocorrelation estimated with non-recursive yule_walker\n')
print(pacf_yw(ys))

# plot the graphics
plot_acf(ys, title='Auto-correlation function of ys')
plt.show()

plot_pacf(ys, title='Partial auto-correlation function of ys')
plt.show()


#using real data
print("Using real data")
file = '/Users/Michael_Enudi/Documents/Learnings/data_dump/perrin-freres-monthly-champagne.csv'
df = pd.read_csv(file, skipfooter=2, parse_dates=True, index_col=0, infer_datetime_format=True)
df.index.freq = 'MS'
df.columns = ['sales']

print('\n=>Calculate the autocorrelation function on sales data\n')
print(acf(df['sales']))

print('\n=>Partial autocorrelation estimate on sales series.\n')
print(pacf(df['sales']))

print('\n=>Partial autocorrelation estimated on sales data via OLS.\n')
print(pacf_ols(df['sales']))

print('\n=>Partial autocorrelation estimated on sales data with non-recursive yule_walker\n')
print(pacf_yw(df['sales']))

# plot the graphics
plot_acf(df['sales'], title='Auto-correlation function of sales time series')
plt.show()

plot_pacf(df['sales'], title='Partial auto-correlation function of sales time series')
plt.show()


print("On differenced sales")
df['diff'] = df['sales'].diff()

print('\n=>Calculate the autocorrelation function on differenced sales data\n')
print(acf(df['diff']))

print('\n=>Partial autocorrelation estimate on differenced sales series.\n')
print(pacf(df['diff']))

print('\n=>Partial autocorrelation estimated on differenced sales data via OLS.\n')
# print(pacf_ols(df['diff']))# failed

print('\n=>Partial autocorrelation estimated on differenced sales data with non-recursive yule_walker\n')
print(pacf_yw(df['diff']))

# plot the graphics
plot_acf(df['diff'], title='Auto-correlation function of differenced sales time series')
plt.show()

plot_pacf(df['diff'], title='Partial auto-correlation function of differenced sales time series')
plt.show()
