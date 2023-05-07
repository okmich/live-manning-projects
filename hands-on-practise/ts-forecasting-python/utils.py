import numpy as np
import pandas as pd

from itertools import product
from typing import Callable, Tuple

from statsmodels.tsa.arima.model import ARIMA



def np_rolling_forecast(arr: np.ndarray, total_len: int, train_len: int, window: int, method: str='mean', 
                     arma_order: Tuple[int, int, int] = None, predict_func:  Callable = None) -> list :
    predicts = []
    for i in range(train_len, total_len, window):
        if method == 'mean':
            mean = np.mean(arr[:i])
            predicts.extend(mean for _ in range(window))
        elif method == 'last':
            last_value = arr[:i][-1]
            predicts.extend([last_value for _ in range(window)])
        elif method == 'arma':
            predicts.extend(predict_func(arr[:i], arma_order, window))
        else:
            raise ValueError('Wrong value of "method" entered.')
    
    return predicts


def pd_rolling_forecast(ser: pd.Series, total_len: int, train_len: int, window: int, method: str='mean', 
                        arma_order: Tuple[int, int, int] = None, predict_func: Callable = None) -> list:
    predicts = []
    for i in range(train_len, total_len, window):
        if method == 'mean':
            mean = np.mean(ser[:i])
            predicts.extend(mean for _ in range(window))
        elif method == 'last':
            last_value = ser[:i].iloc[-1]
            predicts.extend(last_value for _ in range(window))
        elif method == 'ARMA':
            predicts.extend(predict_func(ser[:i], arma_order, window))
        else:
            raise ValueError('Wrong value of "method" entered.')
    
    return predicts


def optimize_arima(ser : pd.Series, p_max : int = 4, q_max : int = 4, d : int = 1):
    ps = range(0, p_max)
    qs = range(0, q_max)
    combinations = product(ps, qs)
    
    arr = []
    for o in combinations:
        try:
            model = ARIMA(ser, order=(o[0], d, o[1])).fit()        
            arr.append((o, model.aic, model.sse))
        except:
            continue
    
    res = pd.DataFrame(arr, columns=['order', 'aic', 'sse'])
    res.sort_values(by=['aic'], inplace=True)
    
    return res
    
    