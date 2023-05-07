import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

import itertools

from statsmodels.tsa.holtwinters import ExponentialSmoothing

def get_dataframe():
	file = '/Users/Michael_Enudi/Documents/Learnings/data_dump/perrin-freres-monthly-champagne.csv'
	df = pd.read_csv(file, skipfooter=2, parse_dates=True, index_col=0, infer_datetime_format=True)
	df.index.freq = 'MS'
	df.columns = ['sales']
	return df

# error metrics
def rmse(y, y_hat):
	return np.sqrt(np.sum((y - y_hat) ** 2))

def mae(y, y_hat):
	return np.mean(np.abs(y - y_hat))

def get_param_combinations():
	TRENDS_OPTIONS = ['add', 'mul']
	DAMPED_TREND_OPTION = [True, False]
	SEASONALITY_OPTION = ['add', 'mul']
	USE_BOXCOX = [True, False, 0]
	INIT_OPTIONS = ['estimated', 'heuristic', 'legacy-heuristic']
	OPTIMIZE_FIT_OPTION = [True, False]

	PARAM_COMBINATIONS = []
	for x in itertools.product(TRENDS_OPTIONS, DAMPED_TREND_OPTION, SEASONALITY_OPTION, INIT_OPTIONS, USE_BOXCOX, OPTIMIZE_FIT_OPTION):
		PARAM_COMBINATIONS.append(x)

	return PARAM_COMBINATIONS


def fit_model(df, param):
	model = ExponentialSmoothing(
		df['sales'], 
		trend=param[0],
		damped_trend=param[1],
		seasonal=param[2],
		seasonal_periods=12,
		initialization_method=param[3],
		use_boxcox=param[4])
	return model.fit(optimized=param[5])


def get_train_test_set(df, forecast_n_cases):
	train = df.iloc[:-forecast_n_cases]
	test = df.iloc[-forecast_n_cases:]

	boundary_dt = df.iloc[-forecast_n_cases].name

	train_idx = df.index < boundary_dt
	test_idx = df.index >= boundary_dt

	return train, test, train_idx, test_idx


def run_model_with_param(train, test, idx, param, forecast_n):
	try:
		fits = fit_model(train, param)

		y_hats = fits.fittedvalues
		train_rsme = rmse(train['sales'], y_hats)
		train_mae = mae(train['sales'], y_hats)

		y_hats_test = fits.forecast(forecast_n)
		test_rsme = rmse(test['sales'], y_hats_test)
		test_mae = mae(test['sales'], y_hats_test)

		return {'id': idx, 'param' : param, 'sse' : round(fits.sse, 4), 'train_rsme'  : round(train_rsme, 4), \
			'train_mae' : round(train_mae, 4), \
			'test_rsme': round(test_rsme, 4), \
			'test_mae': round(test_mae, 4), \
			'aic': round(fits.aic, 4), \
			'bic': round(fits.bic, 4)}
	except:
		print("Error occured fitting", param)
		return None


def sort_limit(results, idx, lim):
	res = None
	sorting_func = lambda x: x[sorting_columns[idx]]
	if idx in [5, 6]:
		res = sorted(results, key=sorting_func, reverse=True)
	else:
		res = sorted(results, key=sorting_func)

	return res[:lim]


def add_fitted_and_test_values(df, res, train_idx, test_idx, forecast_n):
	best_fit = fit_model(df[train_idx], res['param'])

	df.loc[train_idx, 'Holt-Winter'] = best_fit.fittedvalues
	df.loc[test_idx, 'Holt-Winter'] = best_fit.forecast(forecast_n)
	return df


if __name__ == '__main__':
	# get data
	data = get_dataframe()
	# train the model
	params = get_param_combinations()
	forecast_n = 8

	(train, test, train_idx, test_idx) = get_train_test_set(data, forecast_n)

	all_results = []
	for idx, _param in enumerate(params):
		res = run_model_with_param(train, test, idx, _param, forecast_n)
		if res is not None:
			all_results.append(res)


	sorting_columns = {0 : 'sse', 1 : 'train_rsme', 2 : 'train_mae', 3 : 'test_rsme', 4 : 'test_mae', 5 : 'aic', 6 : 'bic'}
	msg = "select one of the following options or enter '-1,' to exit"
	for k in sorting_columns:
		msg += f"\n\t {k} : {sorting_columns[k]}"

	while(True):
		print(msg)
		inp = input()
		cmd = inp.split(',')
		if len(cmd) < 2:
			pass
		elif cmd[0] == '-1':
			break
		else:
			result = sort_limit(all_results, int(cmd[0]), int(cmd[1]))
			for i in result:
				print(i)

			df = add_fitted_and_test_values(data, result[0], train_idx, test_idx, forecast_n)
			df[['sales', 'Holt-Winter']].plot(figsize=(20, 7))
			plt.show()



