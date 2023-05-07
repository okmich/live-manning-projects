import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from statsmodels.tsa.stattools import adfuller

# testing
file = '/Users/Michael_Enudi/Documents/Learnings/data_dump/AirPassengers.csv'
df = pd.read_csv(file, index_col='Month', parse_dates=True)

adf = adfuller(df['#Passengers'])
t_stats = adf[0]
p_value = adf[1]
used_lags = adf[2]
no_obs = adf[3]
critical_values = adf[4]

def test_stationarity(X):
	adf = adfuller(X)
	print(f'T-Statistics	: {adf[0]}')
	print(f'P-Value	: {adf[1]}')
	print(f'Use lags	: {adf[2]}')
	print(f'Critical values	: {adf[4]}')
	if adf[1] < 0.05:
		print('Stationary!')
	else:
		print('Non stationary!')


# test passengers
test_stationarity(df['#Passengers'])

# test the diff of passengers
test_stationarity(df['#Passengers'].diff().dropna())

# test the log of passengers
df['logpassengers'] = np.log(df['#Passengers'])
df['logpassengers'].plot()
plt.show()

test_stationarity(df['logpassengers'])

# test the diff of log of passengers
df['logpassengers'] = np.log(df['#Passengers'])
df['logpassengers'].plot()
plt.show()

test_stationarity(df['logpassengers'].diff().dropna())



# test a list of random number from a gaussian distribution
g_random = np.random.randn(1000)
plt.plot(g_random)
plt.show()

test_stationarity(g_random)


# testing with stock prices
file = '/Users/Michael_Enudi/Documents/Learnings/data_dump/sp500sub.csv'
df = pd.read_csv(file)

tickers = ['INCY', 'AAPL', 'SBUX', 'COO', 'TSCO', 'CBOE']

for ticker in tickers:
	stock_close = df[df['Name'] == ticker]['Close']

	test_stationarity(stock_close)
	stock_close.plot(title=f'Closing prices for {ticker}')
	plt.show()


	# test on log prices
	log_prices = np.log(stock_close)
	test_stationarity(log_prices)
	log_prices.plot(title=f'Log of closing prices for {ticker}')
	plt.show()


	# test on arithmetic returns 
	arth_ret = stock_close.pct_change().dropna()
	test_stationarity(arth_ret)
	arth_ret.plot(title=f'Arithmetic returns of closing prices for {ticker}')
	plt.show()


	# test on log returns
	log_ret = log_prices.diff().dropna()
	test_stationarity(log_ret)
	log_ret.plot(title=f'Log returns of closing prices for {ticker}')
	plt.show()


