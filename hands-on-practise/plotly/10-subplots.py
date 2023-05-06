# to create subplots, we must use the plotly.subplots package and not the plotly.tools package.

from plotly.subplots import make_subplots
import plotly.graph_objects as go

#for data
import numpy as np
from plotly.express import data 


stocks = data.stocks()
fig = make_subplots(rows=3, cols=3,
	subplot_titles=["AMZN/NFLX", "MSFT/NFLX", "MSFT/AMZN", "AAPL/GOOG", "GOOG/NFLX", "FB/AAPL", "NFLX/FB", "NFLX/GOOG", "FB/GOOG"],
	vertical_spacing=0.1)

trace1 = go.Scatter(x=stocks['AMZN'], y=stocks['NFLX'], mode='markers')
trace2 = go.Scatter(x=stocks['MSFT'], y=stocks['NFLX'], mode='markers')
trace3 = go.Scatter(x=stocks['MSFT'], y=stocks['AMZN'], mode='markers')
trace4 = go.Scatter(x=stocks['AAPL'], y=stocks['GOOG'], mode='markers')
trace5 = go.Scatter(x=stocks['GOOG'], y=stocks['NFLX'], mode='markers')
trace6 = go.Scatter(x=stocks['FB'], y=stocks['AAPL'], mode='markers')
trace7 = go.Scatter(x=stocks['NFLX'], y=stocks['FB'], mode='markers')
trace8 = go.Scatter(x=stocks['FB'], y=stocks['GOOG'], mode='markers')
trace9 = go.Scatter(x=stocks['NFLX'], y=stocks['GOOG'], mode='markers')


fig.add_trace(trace1, row=1, col=1)
fig.add_trace(trace2, row=1, col=2)
fig.add_trace(trace3, row=1, col=3)
fig.add_trace(trace4, row=2, col=1)
fig.add_trace(trace5, row=2, col=2)
fig.add_trace(trace6, row=2, col=3)
fig.add_trace(trace7, row=3, col=1)
fig.add_trace(trace8, row=3, col=2)
fig.add_trace(trace9, row=3, col=3)

fig.update_layout(title_text="Stock/Stock Comparison")

fig.write_html("10-subplots.html", auto_open=True)