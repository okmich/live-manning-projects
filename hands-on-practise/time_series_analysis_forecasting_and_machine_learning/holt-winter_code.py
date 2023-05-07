import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

from statsmodels.tsa.holtwinters import ExponentialSmoothing

import itertools

# error metrics
def rmse(y, y_hat):
	return np.sqrt(np.sum((y - y_hat) ** 2))


def mae(y, y_hat):
	return np.mean(np.abs(y - y_hat))


TRENDS_OPTIONS = ['add', 'mul']
DAMPED_TREND_OPTION = [True, False]
SEASONALITY_OPTION = ['add', 'mul']
USE_BOXCOX = [True, False, 0]
INIT_OPTIONS = ['estimated', 'heuristic', 'legacy-heuristic']
OPTIMIZE_FIT_OPTION = [True, False]

PARAM_COMBINATIONS = []
for x in itertools.product(TRENDS_OPTIONS, DAMPED_TREND_OPTION, SEASONALITY_OPTION, INIT_OPTIONS, USE_BOXCOX, OPTIMIZE_FIT_OPTION):
	PARAM_COMBINATIONS.append(x)

file = '/Users/Michael_Enudi/Documents/Learnings/data_dump/AirPassengers.csv'
df = pd.read_csv(file, index_col='Month', parse_dates=True)

# assign frequence to the df index
df.index.freq = 'MS'

# split into train/test set
N_test = 12
train = df.iloc[:-N_test]
test = df.iloc[-N_test:]

train_idx = df.index <= train.index[-1]
test_idx = df.index > train.index[-1]

results = []

def get_fitted_model(param):
	model = ExponentialSmoothing(
		train['#Passengers'], 
		trend=param[0],
		damped_trend=param[1],
		seasonal=param[2],
		seasonal_periods=12,
		initialization_method=param[3],
		use_boxcox=param[4])
	return model.fit(optimized=param[5])


def run_model_on_param(idx, param):
	try:
		fits = get_fitted_model(param)

		y_hats = fits.fittedvalues
		train_rsme = rmse(train['#Passengers'], y_hats)
		train_mae = mae(train['#Passengers'], y_hats)

		y_hats_test = fits.forecast(N_test)
		test_rsme = rmse(test['#Passengers'], y_hats_test)
		test_mae = mae(test['#Passengers'], y_hats_test)

		results.append({'id': idx, 'param' : param, 'sse' : round(fits.sse, 4), 'train_rsme'  : round(train_rsme, 4), \
			'train_mae' : round(train_mae, 4), \
			'test_rsme': round(test_rsme, 4), \
			'test_mae': round(test_mae, 4), \
			'aic': round(fits.aic, 4), \
			'bic': round(fits.bic, 4)})
	except:
		print("Error occured fitting", param)

sorting_columns = {0 : 'sse', 1 : 'train_rsme', 2 : 'train_mae', 3 : 'test_rsme', 4 : 'test_mae', 5 : 'aic', 6 : 'bic'}

def sort_limit(idx, lim):
	res = None
	sorting_func = lambda x: x[sorting_columns[idx]]
	if idx in [5, 6]:
		res = sorted(results, key=sorting_func, reverse=True)
	else:
		res = sorted(results, key=sorting_func)

	return res[:lim]


def add_fitted_and_test_values(res):
	best_fit = get_fitted_model(res['param'])

	df.loc[train_idx, 'Holt-Winter'] = best_fit.fittedvalues
	df.loc[test_idx, 'Holt-Winter'] = best_fit.forecast(N_test)
	return df


if __name__ == '__main__':
	# train the model
	for idx, _param in enumerate(PARAM_COMBINATIONS):
		run_model_on_param(idx, _param)

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
			result = sort_limit(int(cmd[0]), int(cmd[1]))
			for i in result:
				print(i)

			df = add_fitted_and_test_values(result[0])
			df[['#Passengers', 'Holt-Winter']].plot(figsize=(20, 7))
			plt.show()










