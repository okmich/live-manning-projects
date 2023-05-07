import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

from statsmodels.tsa.holtwinters import SimpleExpSmoothing

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
ses = SimpleExpSmoothing(df['#Passengers'], initialization_method='legacy-heuristic')

fitted_model = ses.fit(smoothing_level=alpha, optimized=True) # check the huge difference when optimized=False

fitted_values = fitted_model.fittedvalues

df['ses'] = fitted_model.predict(start=df.index[0], end=df.index[-1])

# use the np.allclose to see any equality between the fitted and predicted values
np.allclose(df['ses'], fitted_values) # yes, equally

# lets add the ewma and compare
df['ewma'] = df['#Passengers'].ewm(alpha=alpha, adjust=True).mean()  # also check the huge difference when optimized=False

df.head()

# plot
df.plot()
plt.show()


########## Now let test the train-test and forecast scenario
N_test = 12

train = df.iloc[:-N_test]
test = df.iloc[N_test:]

#train with training set
ses2 = SimpleExpSmoothing(train['#Passengers'], initialization_method='legacy-heuristic')
fit2 = ses2.fit(smoothing_level=alpha, optimized=True) # check the huge difference when smoothing_level=alpha, optimized=True

print(fit2.params)

train_idx = df.index <= train.index[-1]
test_idx = df.index > train.index[-1]

df.loc[train_idx, 'ses_fitted'] = fit2.fittedvalues
df.loc[test_idx, 'ses_fitted'] = fit2.forecast(N_test)

df[['#Passengers', 'ses_fitted']].plot()
plt.show()


