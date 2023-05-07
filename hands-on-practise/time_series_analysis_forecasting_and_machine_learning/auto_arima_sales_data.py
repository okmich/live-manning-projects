import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

from pmdarima import auto_arima

file = '/Users/Michael_Enudi/Documents/Learnings/data_dump/perrin-freres-monthly-champagne.csv'
data = pd.read_csv(file, skipfooter=2, parse_dates=True, index_col=0, infer_datetime_format=True)
data.index.freq = 'MS'
data.columns = ['sales']

def split_train_test_sets(df, fc_n):
	train = df[:-fc_n]
	test = df[-fc_n:]
	train_idx = df.index <= train.index[-1]
	test_idx = df.index > train.index[-1]
	return train, test, train_idx, test_idx

forecast_n = 12
(train, test, train_idx, test_idx) = split_train_test_sets(data, forecast_n)

# create the auto arima class and train
arima = auto_arima(train['sales'], error_action='ignore', trace=True)