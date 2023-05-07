import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

from statsmodels.tsa.stattools import adfuller
from statsmodels.tsa.arima.model import ARIMA
from statsmodels.graphics.tsaplots import plot_acf, plot_pacf

file = '/Users/Michael_Enudi/Documents/Learnings/data_dump/perrin-freres-monthly-champagne.csv'
data = pd.read_csv(file, skipfooter=2, parse_dates=True, index_col=0, infer_datetime_format=True)
data.index.freq = 'MS'
data.columns = ['sales']

# first thing should be using exploratory analysis to see the data
# then using acf and pacf plots to kinda figure out the best values 
# for MA(q) and AR(p) respectively.
# time series plot also helps decide if there is seasonality in the
# data or if we should apply some transformation prior to fitting 

# because the original dataset shows seasonality and it is not stationary,
# we will see what transformation can bring out stationarity
data[['sales']].plot(title='Month sales') 
plt.show() # plot show a certain seasonality and not stationary

data['sales_diff'] = data['sales'].diff()
data[['sales_diff']].plot(title='Month sales differenced') 
plt.show() 

data['sqrt_sales'] = np.sqrt(data['sales'])
data[['sqrt_sales']].plot(title='Root of monthly sales') 
plt.show() 

data['sqrt_sales_diff'] = data['sqrt_sales'].diff()
data[['sqrt_sales_diff']].plot(title='Root of monthly sales differenced') 
plt.show()


data['log_sales'] = np.log(data['sales'])
data[['log_sales']].plot(title='Log of monthly sales') 
plt.show() 

data['log_sales_diff'] = data['log_sales'].diff()
data[['log_sales_diff']].plot(title='Log of monthly sales differenced') 
plt.show() 

# drop na
data.dropna(inplace=True)

# we will take 2 years of data for testing
forecast_n = 24

def split_train_test_sets(df, fc_n):
	train = df[:-fc_n]
	test = df[-fc_n:]
	train_idx = df.index <= train.index[-1]
	test_idx = df.index > train.index[-1]

	return train, test, train_idx, test_idx

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



(train, test, train_idx, test_idx) = split_train_test_sets(data, forecast_n)


# now using the training set, we will test for stationarity
# it is pretty obvious that the some of the columns are not stationary
# so I focus on the non-obvious ones.
test_stationarity(train['sales'])			# Non stationary!
test_stationarity(train['sales_diff'])		# Stationary!
test_stationarity(train['sqrt_sales'])		# Non stationary!
test_stationarity(train['sqrt_sales_diff'])	# Stationary!
test_stationarity(train['log_sales'])		# Non stationary!
test_stationarity(train['log_sales_diff'])  # stationary


# I will focus on the stationary series and try to create acf and pacf plots for them
def plot_acf_and_pacf(arr, part_title=''):
	plot_acf(arr, title=f'ACF for {part_title}')
	plt.show()
	plot_pacf(arr, title=f'PACF for {part_title}')
	plt.show()


plot_acf_and_pacf(train['sales_diff'], part_title='differenced monthly sales')
plot_acf_and_pacf(train['sqrt_sales_diff'], part_title='differenced root of monthly sales')
plot_acf_and_pacf(train['log_sales_diff'], part_title='differenced log of monthly sales')

# arima_model = ARIMA(train['sales_diff'], order=(12,0,0))
# arima_result = arima_model.fit()
# arima_result.summary()

# arima_prediction = arima_result.get_prediction(start=train.index[0], end=train.index[-1])
# arima_forecast = arima_result.get_forecast(forecast_n)
# forecast_conf_int = arima_forecast.conf_int()

# data.loc[train_idx, 'arima'] = arima_prediction.predicted_mean
# data.loc[test_idx, 'arima'] = arima_forecast.predicted_mean

# data[['sales_diff', 'arima']].plot()
# plt.show()


def plot_arima_fit_and_results(df, train_idx, test_idx, in_sample, out_sample, out_conf_int):
	fig, ax = plt.subplots(figsize=(20,6))

	ax.plot(df['sales'], label='data')
	ax.plot(data[train_idx].index, in_smaple, label='in-sample')
	ax.plot(data[test_idx].index, out_sample, label='forecast')

	out_conf_int.columns = ['lower', 'upper']
	ax.fill(test.index, out_conf_int['lower'], out_conf_int['upper'], color='red', alpha=0.3)

	ax.legend()
	plt.show()


def rsme(y, y_hat):
	return np.sqrt(np.sum((y - y_hat) ** 2))


def fit_arima_and_return_results(df, train_idx, test_idx, o, forecast_n):
	arima = ARIMA(df[train_idx]['sales'], order=o)

	arima_model = arima.fit()

	in_sample_data = arima_model.fittedvalues # or arima_model.get_prediction(start=train.index[0], end=train.index[-1]).predicted_mean

	arima_forecast = arima_result.get_forecast(forecast_n)
	out_sample_data = arima_forecast.predicted_mean
	conf_ints = arima_forecast.conf_int()
	aic = arima_model.aic
	sse = arima_model.sse
	rsme_metric = rsme(df[test_idx]['sales'], out_sample_data)

	return {'model' : arima_model, 'aic' : aic, 'sse' : sse, 'rsme' : rsme_metric}



