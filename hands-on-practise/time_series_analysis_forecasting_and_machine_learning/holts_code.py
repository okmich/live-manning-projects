import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

from statsmodels.tsa.holtwinters import Holt

file = '/Users/Michael_Enudi/Documents/Learnings/data_dump/AirPassengers.csv'

df = pd.read_csv(file, index_col='Month', parse_dates=True)

df.head()

# assign frequence to the df index
print(df.index)
print(df.index.freq)
df.index.freq = 'MS'
print(df.index)
print(df.index.freq)

alpha = 0.2

holt = Holt(df['#Passengers'], initialization_method='legacy-heuristic')
fits = holt.fit(smoothing_level=alpha, optimized=True)

df['holt'] = fits.fittedvalues

df.plot()
plt.show()


########## Now let test the train-test and forecast scenario
N_test = 12

train = df.iloc[:-N_test]
test = df.iloc[N_test:]

#train with training set
holt2 = Holt(train['#Passengers'], initialization_method='legacy-heuristic')
fit2 = holt2.fit(optimized=True) # check the huge difference when smoothing_level=alpha, optimized=True

print(fit2.params)

train_idx = df.index <= train.index[-1]
test_idx = df.index > train.index[-1]

df.loc[train_idx, 'holt_fitted'] = fit2.fittedvalues
df.loc[test_idx, 'holt_fitted'] = fit2.forecast(N_test)

df[['#Passengers', 'holt_fitted']].plot()
plt.show()
