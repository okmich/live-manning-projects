import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

spy_500_file = '/Users/Michael_Enudi/Documents/Learnings/data_dump/sp500_close.csv'

close = pd.read_csv(spy_500_file, index_col=0, parse_dates=True)

goog = close[['GOOG']].copy().dropna()

goog.head()

goog.plot()
plt.show()


goog_ret = np.log(goog.pct_change() + 1)

goog_ret.plot()
plt.show()

# add sma series
goog_rolling_window = goog['GOOG'].rolling(10)

# check the type of a rolling window
type(goog_rolling_window)

#add the rolling mean series
goog['sma10'] = goog_rolling_window.mean()

goog.plot(figsize=(10,5))
plt.show()

goog['sma50'] = goog['GOOG'].rolling(50).mean()

goog.head(55)

goog.plot(figsize=(10,5))
plt.show()


# covariance and correlation rolling matrices
goog_appl = close[['GOOG', 'AAPL']].dropna().rolling(50)
goog_appl_cov = goog_appl.cov()
goog_appl_corr = goog_appl.corr()
# both operations gives us a multi-index matrix

goog_appl_ret = np.log(1 + close[['GOOG', 'AAPL']].dropna().pct_change())

goog_appl_ret['goog_ret_50'] = goog_appl_ret['GOOG'].rolling(50).mean()
goog_appl_ret['appl_ret_50'] = goog_appl_ret['AAPL'].rolling(50).mean()

goog_appl_ret.plot(figsize=(20,8))
plt.show()

goog_appl_ret_cov = goog_appl_ret[['goog_ret_50', 'appl_ret_50']].rolling(50).cov()
goog_appl_ret_cov.tail()

goog_appl_ret_corr = goog_appl_ret[['goog_ret_50', 'appl_ret_50']].rolling(50).corr()
goog_appl_ret_corr.tail()
