import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

file = '/Users/Michael_Enudi/Documents/Learnings/data_dump/AirPassengers.csv'

df = pd.read_csv(file, index_col='Month', parse_dates=True)

df.head()

df.plot()
plt.show()

alpha = 0.2

# calc exponential ma
ewm_obj = df['#Passengers'].ewm(alpha=alpha, adjust=False)
type(ewm_obj)

df['passenger_ewma'] = ewm_obj.mean()

# manually calculate ewma

def my_ewma(s, alpha):
	manual_result = []
	numpy_arr = s.to_numpy()
	for i in numpy_arr:
		if len(manual_result) == 0:
			manual_result.append(i)
		elif np.isnan(manual_result[-1]):
			manual_result.append(i)
		else:
			new_value = (alpha * i) + (1 - alpha) * manual_result[-1]
			manual_result.append(new_value)
	return manual_result

df['passenger_my_ewma'] = my_ewma(df['#Passengers'], alpha)

#plot
df.plot()
plt.show()



