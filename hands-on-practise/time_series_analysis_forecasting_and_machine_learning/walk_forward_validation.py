import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

from statsmodels.tsa.holtwinters import ExponentialSmoothing

# error metrics
def rmse(y, y_hat):
	return np.sqrt(np.sum((y - y_hat) ** 2))


def mae(y, y_hat):
	return np.mean(np.abs(y - y_hat))

file = '/Users/Michael_Enudi/Documents/Learnings/data_dump/AirPassengers.csv'
df = pd.read_csv(file, index_col='Month', parse_dates=True)
df.index.freq = 'MS'

print(df.shape) # 144

def get_fitted_model(series, param):
	model = ExponentialSmoothing(
		series, 
		trend=param['trend'],
		damped_trend=param['damped_trend'],
		seasonal=param['seasonality'],
		seasonal_periods=12,
		initialization_method=param['init_option'],
		use_boxcox=param['use_boxcox'])
	return model.fit(optimized=param['optimize_fit'])


def run_model_for_step(df_part, param, end_idx):
	series_name = '#Passengers'
	n_test_count = 8
	train = df_part.iloc[:-n_test_count]
	test = df_part.iloc[-n_test_count:]

	try:
		fits = get_fitted_model(train, param)

		y_hats = fits.fittedvalues
		train_rsme = rmse(train[series_name], y_hats)
		train_mae = mae(train[series_name], y_hats)

		y_hats_test = fits.forecast(n_test_count)
		test_rsme = rmse(test[series_name], y_hats_test) 
		test_mae = mae(test[series_name], y_hats_test)

		return {'rec': end_idx, 'param' : param, 'sse' : round(fits.sse, 4), \
			'train_rsme'  : round(train_rsme, 4), \
			'train_mae' : round(train_mae, 4), \
			'test_rsme': round(test_rsme, 4), \
			'test_mae': round(test_mae, 4), \
			'aic': round(fits.aic, 4), \
			'bic': round(fits.bic, 4)}
	except:
		print("Error occured fitting", param)
		return None


if __name__ == '__main__':
	param = {
		"trend" : "add",
		"damped_trend" : False,
		"seasonality" : "mul",
		"init_option" : "estimated",
		"use_boxcox" : False,
		"optimize_fit" : True
	}

	# we know that the shape of the data is 144 or df.shape[0]
	# we begin from 72 and move in steps of 24
	begin = 72
	step = 24

	results = []
	for i in range(72, df.shape[0]+step, step):
		res = run_model_for_step(df.iloc[:i], param, i)
		if res is not None:
			results.append(res)
			print(res)

	# treat and handle the best model for display













